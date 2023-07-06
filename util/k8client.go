package util

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"strings"
)

// ClusterNodes represents the structure of data to be marshalled into JSON
type ClusterNodes struct {
	Cluster string    `json:"cluster"` // Name of the cluster
	Nodes   []v1.Node `json:"nodes"`   // List of nodes in the cluster
}

func GetKubeConfigPath() string {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	return *kubeconfig
}

// GetDeployments lists the deployments in all namespaces.
func GetDeployments(kubeconfig string) ([]string, error) {
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

	var deploymentNames []string
	for _, deployment := range deployments.Items {
		deploymentNames = append(deploymentNames, deployment.Name)
	}

	return deploymentNames, nil
}

// GetNetworkPluginPodName determines which CNI is deployed by explicitly searching for Calico|Cilium pods
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

	return "", fmt.Errorf("No matching pod found")
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
