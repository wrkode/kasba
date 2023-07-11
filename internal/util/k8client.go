package util

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wrkode/kasba/internal/nodeinfo"
	"path/filepath"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// ClusterNodes represents the structure of data to be marshalled into JSON
type ClusterNodes struct {
	Cluster string    `json:"cluster"` // Name of the cluster
	Nodes   []v1.Node `json:"nodes"`   // List of nodes in the cluster
}

type KubeConfig struct {
	kubeconfig *string
	config     *rest.Config
	clientset  *kubernetes.Clientset
}

// GetKubeConfigPath gathers/uses active kubeconfig TODO: Implement path flag
func (k *KubeConfig) GetKubeConfigPath() error {
	var err error
	if home := homedir.HomeDir(); home != "" {
		k.kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		k.kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// build Config
	if *k.kubeconfig != "" {
		k.config, err = clientcmd.BuildConfigFromFlags("", *k.kubeconfig)
	} else {
		k.config, err = rest.InClusterConfig()
	}

	// handle config error
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}

	// creates the clientset
	k.clientset, err = kubernetes.NewForConfig(k.config)

	// handle clientset error
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	return nil
}

// NamespaceExists checks if the given namespace exists in the cluster.
func (k *KubeConfig) NamespaceExists(namespaceName string) (bool, error) {
	// get namespaces
	namespaces, err := k.clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to get namespaces: %v", err)
	}

	// check if the desired namespace exists
	for _, ns := range namespaces.Items {
		if ns.Name == namespaceName {
			return true, nil
		}
	}

	return false, nil
}

type DeploymentInfo struct {
	Name      string
	Namespace string
}

// GetDeployments lists the deployments in all namespaces and returns them with NAMES and NAMESPACE.
func (k *KubeConfig) GetDeployments() ([]DeploymentInfo, error) {
	deployments, err := k.clientset.AppsV1().Deployments(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var deploymentInfoList []DeploymentInfo
	for _, deployment := range deployments.Items {
		deploymentInfoList = append(deploymentInfoList, DeploymentInfo{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
		})
	}

	return deploymentInfoList, nil
}

// GetNetworkPluginPodName determines which CNI is deployed by explicitly searching for Calico|Cilium pods. K3s Will return an error.
func (k *KubeConfig) GetNetworkPluginPodName() (string, error) {
	pods, err := k.clientset.CoreV1().Pods(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return "", err
	}

	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, "calico") || strings.Contains(pod.Name, "cilium") {
			// Truncate the pod name at the first hyphen
			truncatedName := strings.SplitN(pod.Name, "-", 2)[0]
			return truncatedName, nil
		}
	}

	return "", fmt.Errorf("unable to detect CNI - is this K3s")
}

// FetchClustersJSON fetches node information for the active context in kubeconfig and returns it as JSON.
func (k *KubeConfig) FetchClustersJSON() ([]byte, error) {
	nodes, err := k.clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error getting nodes for active context: %v", err)
	}

	activeContext := k.config.Host // This will get the API Server URL instead of the context name

	clusterNodes := ClusterNodes{
		Cluster: activeContext,
		Nodes:   nodes.Items,
	}

	jsonData, err := json.MarshalIndent(clusterNodes, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Error marshaling cluster nodes to JSON: %v", err)
	}

	return jsonData, nil
}

// GetNodeInfo get nodeinfo from active context
func (k *KubeConfig) GetNodeInfo() (nodeinfo.NodesInfo, error) {
	// Fetch JSON data
	jsonData, err := k.FetchClustersJSON()
	if err != nil {
		return nodeinfo.NodesInfo{}, fmt.Errorf("error fetching nodes: %v", err)
	}
	//fmt.Println(string(jsonData)) # using it for check jsonData
	// Unmarshal the JSON data into the struct.
	var data nodeinfo.NodesInfo
	err = json.Unmarshal(jsonData, &data) // jsonData is already []byte, so no need to cast it again
	if err != nil {
		return nodeinfo.NodesInfo{}, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return data, nil
}
