package util

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// ClusterNodes represents the structure of data to be marshalled into JSON
type ClusterNodes struct {
	Cluster string    `json:"cluster"` // Name of the cluster
	Nodes   []v1.Node `json:"nodes"`   // List of nodes in the cluster
}

// FetchClustersJSON fetches node information for the active context in kubeconfig and returns it as JSON.
func FetchClustersJSON() ([]byte, error) {
	// Determine the path to kubeconfig file
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		// If home directory is present, use kubeconfig file from ~/.kube/config
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		// Otherwise, expect user to provide the path to kubeconfig file through command-line flag
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Create a configuration with the provided kubeconfig file
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: *kubeconfig},
		&clientcmd.ConfigOverrides{},
	)

	// Get REST config for active context
	restConfig, err := config.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("Error creating client config for active context: %v", err)
	}

	// Create clientset to interact with the Kubernetes API server
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("Error creating clientset for active context: %v", err)
	}

	// Retrieve the list of nodes in the active context
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error getting nodes for active context: %v", err)
	}

	// NOTE: The following line gets the default filename, not the active context name.
	// You may need to use another approach to get the actual context name.
	activeContext := config.ConfigAccess().GetDefaultFilename()

	// Populate the structure with retrieved node information
	clusterNodes := ClusterNodes{
		Cluster: activeContext,
		Nodes:   nodes.Items,
	}

	// Marshal the structure into JSON format with indents for readability
	jsonData, err := json.MarshalIndent(clusterNodes, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Error marshaling cluster nodes to JSON: %v", err)
	}

	// Return JSON data
	return jsonData, nil
}
