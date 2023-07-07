package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/wrkode/kasba/nodeinfo"
	"github.com/wrkode/kasba/util"
	"log"
	"time"
)

const (
	SUSE      = "SUSE Software Solutions"
	BOMFormat = "kasba"
	kasbaVersion	  = "v0.1"
)

var (
	createdAt = time.Now()
	kasbaID   = uuid.New().String()

)
func main() {
	printLabel()

	flag.Parse()
	kubeconfig := util.GetKubeConfigPath()

	// Fetch JSON data
	jsonData, err := util.FetchClustersJSON(kubeconfig)
	if err != nil {
		log.Fatalf("Error fetching nodes: %v", err)
	}
	//fmt.Println(string(jsonData)) # using it for check jsonData
	// Unmarshal the JSON data into the struct.
	var data nodeinfo.NodesInfo
	err = json.Unmarshal(jsonData, &data) // jsonData is already []byte, so no need to cast it again
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}


	// Check if there are any items in data
	if len(data.Items) > 0 {
		// Print Cluster Name and Instance Type once, using the first item as a reference
		fmt.Printf("Cluster Name:  %s\n", data.Items[0].Metadata.Annotations.ClusterXK8SIoClusterName)
		fmt.Printf("Instance Type: %s\n", data.Items[0].Metadata.Labels.NodeKubernetesIoInstanceType)
		fmt.Printf("K8s Version:   %s\n", data.Items[0].Status.NodeInfo.KubeletVersion)
	}
	networkPluginPodName, err := util.GetNetworkPluginPodName(kubeconfig)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("CNI:          ", networkPluginPodName)
	LonghornExists, err := util.NamespaceExists(kubeconfig, "longhorn-system")
	if err != nil {
		log.Fatalf("Error checking if Longhorn is installed: %v", err)
	}
	fmt.Printf("Longhorn installed: %t\n", LonghornExists)

	fmt.Println()

	for _, item := range data.Items {
		fmt.Printf("Cluster Machine Name: %s\n", item.Metadata.Annotations.ClusterXK8SIoMachine)
		fmt.Printf("Cluster Node Name:    %s\n", item.Metadata.Name)
		fmt.Printf("Operating System:     %s\n", item.Status.NodeInfo.OperatingSystem)
		fmt.Printf("OS Image:             %s\n", item.Status.NodeInfo.OsImage)
		fmt.Printf("Node Arch:            %s\n", item.Status.NodeInfo.Architecture)
		fmt.Printf("Kernel Version:       %s\n", item.Status.NodeInfo.KernelVersion)
		fmt.Printf("System UUID:          %s\n", item.Status.NodeInfo.SystemUUID)
		fmt.Printf("Container Runtime:    %s\n", item.Status.NodeInfo.ContainerRuntimeVersion)
		fmt.Printf("Kube Version:         %s\n", item.Status.NodeInfo.KubeletVersion)
		fmt.Printf("KubeProxy Version:    %s\n", item.Status.NodeInfo.KubeProxyVersion)
		fmt.Printf("Node Args:            %s\n", item.Metadata.Annotations.Rke2IoNodeArgs)
		fmt.Println("Pod CIDR:                   ", item.Spec.PodCIDR)
		fmt.Printf("Pod Limits:           %s\n", item.Metadata.Annotations.ManagementCattleIoPodLimits)
		fmt.Printf("Pod Requests:         %s\n", item.Metadata.Annotations.ManagementCattleIoPodRequests)
		fmt.Println(" --- Allocatable ---")
		fmt.Printf("CPU:                 %s\n", item.Status.Allocatable.CPU)
		fmt.Printf("Memory:              %s\n", item.Status.Allocatable.Memory)
		fmt.Printf("Ephemeral Storage:   %s\n", item.Status.Allocatable.EphemeralStorage)
		fmt.Printf("Pods:                %s\n", item.Status.Allocatable.Pods)
		fmt.Println()
		fmt.Printf("---Messages %s---\n", item.Metadata.Name)
		for _, condition := range item.Status.Conditions {
			fmt.Println("  Condition Type:", condition.Type)
			fmt.Println("    Last Heartbeat Time:", condition.LastHeartbeatTime)
			fmt.Println("    Last Transition Time:", condition.LastTransitionTime)
			fmt.Println("    Message:", condition.Message)
			fmt.Println("    Reason:", condition.Reason)
			fmt.Println("    Status:", condition.Status)
		}
		fmt.Println()
		//fmt.Println(out.String())
	}
	deployments, err := util.GetDeployments(kubeconfig)
	if err != nil {
		fmt.Printf("Error getting deployments %v\n:", err)
		//os.Exit(1)
	}
	fmt.Println("---Deployments---")
	// Determine the maximum width of the deployment name for formatting
	maxNameWidth := len("NAME")
	for _, deployment := range deployments {
		if len(deployment.Name) > maxNameWidth {
			maxNameWidth = len(deployment.Name)
		}
	}

	// Print the header
	fmt.Printf("%-*s NAMESPACE\n", maxNameWidth, "NAME")

	// Print the deployments
	for _, deployment := range deployments {
		fmt.Printf("%-*s %s\n", maxNameWidth, deployment.Name, deployment.Namespace)
	}
}

// Write call to GetPersistentVolumes()

//printLabel prints out KASBA label
func printLabel() {
	fmt.Println("#####################################################################")
	fmt.Println("Kubernetes As-Built Assessment - KASBA")
	fmt.Printf( "Date:         %s\n", createdAt)
	fmt.Printf("Format:        %s\n", BOMFormat)
	fmt.Printf("KASBA Version: %s\n", kasbaVersion)
	fmt.Println("#####################################################################")
	fmt.Println()
}