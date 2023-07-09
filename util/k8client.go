package util

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
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

// GetKubeConfigPath gathers/uses active kubeconfig TODO: Implement path flag
func GetKubeConfigPath() string {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	return *kubeconfig
}

// NamespaceExists checks if the given namespace exists in the cluster.
func NamespaceExists(kubeconfig, namespaceName string) (bool, error) {
	var config *rest.Config
	var err error

	// build config from flags
	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	// handle config error
	if err != nil {
		return false, fmt.Errorf("failed to build config: %v", err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, fmt.Errorf("failed to create clientset: %v", err)
	}

	// get namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
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
func GetDeployments(kubeconfig string) ([]DeploymentInfo, error) {

	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	deployments, err := clientset.AppsV1().Deployments(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
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
func GetNetworkPluginPodName(kubeconfig string) (string, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return "", err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}

	pods, err := clientset.CoreV1().Pods(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
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

	return "", fmt.Errorf("Unable to detect CNI - is this K3s ??")
}

// FetchClustersJSON fetches node information for the active context in kubeconfig and returns it as JSON.
func FetchClustersJSON(kubeconfig string) ([]byte, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, fmt.Errorf("Error creating client config for active context: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Error creating clientset for active context: %v", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error getting nodes for active context: %v", err)
	}

	activeContext := config.Host // This will get the API Server URL instead of the context name

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
