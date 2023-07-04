package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type NodesInfo struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Annotations struct {
				ClusterXK8SIoClusterName                         string `json:"cluster.x-k8s.io/cluster-name"`
				ClusterXK8SIoClusterNamespace                    string `json:"cluster.x-k8s.io/cluster-namespace"`
				ClusterXK8SIoMachine                             string `json:"cluster.x-k8s.io/machine"`
				ClusterXK8SIoOwnerKind                           string `json:"cluster.x-k8s.io/owner-kind"`
				ClusterXK8SIoOwnerName                           string `json:"cluster.x-k8s.io/owner-name"`
				CsiVolumeKubernetesIoNodeid                      string `json:"csi.volume.kubernetes.io/nodeid"`
				ManagementCattleIoPodLimits                      string `json:"management.cattle.io/pod-limits"`
				ManagementCattleIoPodRequests                    string `json:"management.cattle.io/pod-requests"`
				NodeAlphaKubernetesIoTTL                         string `json:"node.alpha.kubernetes.io/ttl"`
				Rke2IoHostname                                   string `json:"rke2.io/hostname"`
				Rke2IoInternalIP                                 string `json:"rke2.io/internal-ip"`
				Rke2IoNodeArgs                                   string `json:"rke2.io/node-args"`
				Rke2IoNodeConfigHash                             string `json:"rke2.io/node-config-hash"`
				Rke2IoNodeEnv                                    string `json:"rke2.io/node-env"`
				VolumesKubernetesIoControllerManagedAttachDetach string `json:"volumes.kubernetes.io/controller-managed-attach-detach"`
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Finalizers        []string  `json:"finalizers"`
			Labels            struct {
				BetaKubernetesIoArch                   string `json:"beta.kubernetes.io/arch"`
				BetaKubernetesIoInstanceType           string `json:"beta.kubernetes.io/instance-type"`
				BetaKubernetesIoOs                     string `json:"beta.kubernetes.io/os"`
				KubernetesIoArch                       string `json:"kubernetes.io/arch"`
				KubernetesIoHostname                   string `json:"kubernetes.io/hostname"`
				KubernetesIoOs                         string `json:"kubernetes.io/os"`
				NodeRoleKubernetesIoWorker             string `json:"node-role.kubernetes.io/worker"`
				NodeKubernetesIoInstanceType           string `json:"node.kubernetes.io/instance-type"`
				PlanUpgradeCattleIoSystemAgentUpgrader string `json:"plan.upgrade.cattle.io/system-agent-upgrader"`
				RkeCattleIoMachine                     string `json:"rke.cattle.io/machine"`
			} `json:"labels"`
			ManagedFields []struct {
				APIVersion string `json:"apiVersion"`
				FieldsType string `json:"fieldsType"`
				FieldsV1   struct {
					FMetadata struct {
						FAnnotations struct {
							FManagementCattleIoPodLimits struct {
							} `json:"f:management.cattle.io/pod-limits"`
							FManagementCattleIoPodRequests struct {
							} `json:"f:management.cattle.io/pod-requests"`
						} `json:"f:annotations"`
					} `json:"f:metadata"`
				} `json:"fieldsV1,omitempty"`
				Manager   string    `json:"manager"`
				Operation string    `json:"operation"`
				Time      time.Time `json:"time"`
				FieldsV10 struct {
					FMetadata struct {
						FLabels struct {
							FBetaKubernetesIoInstanceType struct {
							} `json:"f:beta.kubernetes.io/instance-type"`
							FNodeKubernetesIoInstanceType struct {
							} `json:"f:node.kubernetes.io/instance-type"`
						} `json:"f:labels"`
					} `json:"f:metadata"`
					FSpec struct {
						FProviderID struct {
						} `json:"f:providerID"`
					} `json:"f:spec"`
				} `json:"fieldsV1,omitempty"`
				FieldsV11 struct {
					FMetadata struct {
						FAnnotations struct {
							FNodeAlphaKubernetesIoTTL struct {
							} `json:"f:node.alpha.kubernetes.io/ttl"`
						} `json:"f:annotations"`
					} `json:"f:metadata"`
					FSpec struct {
						FPodCIDR struct {
						} `json:"f:podCIDR"`
						FPodCIDRs struct {
							NAMING_FAILED struct {
							} `json:"."`
							V10422024 struct {
							} `json:"v:"10.42.2.0/24""`
						} `json:"f:podCIDRs"`
					} `json:"f:spec"`
				} `json:"fieldsV1,omitempty"`
				FieldsV12 struct {
					FStatus struct {
						FConditions struct {
							KTypeNetworkUnavailable struct {
								NAMING_FAILED struct {
								} `json:"."`
								FLastHeartbeatTime struct {
								} `json:"f:lastHeartbeatTime"`
								FLastTransitionTime struct {
								} `json:"f:lastTransitionTime"`
								FMessage struct {
								} `json:"f:message"`
								FReason struct {
								} `json:"f:reason"`
								FStatus struct {
								} `json:"f:status"`
								FType struct {
								} `json:"f:type"`
							} `json:"k:{"type":"NetworkUnavailable"}"`
						} `json:"f:conditions"`
					} `json:"f:status"`
				} `json:"fieldsV1,omitempty"`
				FieldsV13 struct {
					FMetadata struct {
						FAnnotations struct {
							FCsiVolumeKubernetesIoNodeid struct {
							} `json:"f:csi.volume.kubernetes.io/nodeid"`
						} `json:"f:annotations"`
					} `json:"f:metadata"`
					FStatus struct {
						FConditions struct {
							KTypeDiskPressure struct {
								FLastHeartbeatTime struct {
								} `json:"f:lastHeartbeatTime"`
							} `json:"k:{"type":"DiskPressure"}"`
							KTypeMemoryPressure struct {
								FLastHeartbeatTime struct {
								} `json:"f:lastHeartbeatTime"`
							} `json:"k:{"type":"MemoryPressure"}"`
							KTypePIDPressure struct {
								FLastHeartbeatTime struct {
								} `json:"f:lastHeartbeatTime"`
							} `json:"k:{"type":"PIDPressure"}"`
							KTypeReady struct {
								FLastHeartbeatTime struct {
								} `json:"f:lastHeartbeatTime"`
							} `json:"k:{"type":"Ready"}"`
						} `json:"f:conditions"`
						FImages struct {
						} `json:"f:images"`
					} `json:"f:status"`
				} `json:"fieldsV1,omitempty"`
				Subresource string `json:"subresource,omitempty"`
			} `json:"managedFields"`
			Name            string `json:"name"`
			ResourceVersion string `json:"resourceVersion"`
			UID             string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			PodCIDR    string   `json:"podCIDR"`
			PodCIDRs   []string `json:"podCIDRs"`
			ProviderID string   `json:"providerID"`
		} `json:"spec,omitempty"`
		Status struct {
			Addresses []struct {
				Address string `json:"address"`
				Type    string `json:"type"`
			} `json:"addresses"`
			Allocatable struct {
				CPU              string `json:"cpu"`
				EphemeralStorage string `json:"ephemeral-storage"`
				Hugepages1Gi     string `json:"hugepages-1Gi"`
				Hugepages2Mi     string `json:"hugepages-2Mi"`
				Memory           string `json:"memory"`
				Pods             string `json:"pods"`
			} `json:"allocatable"`
			Capacity struct {
				CPU              string `json:"cpu"`
				EphemeralStorage string `json:"ephemeral-storage"`
				Hugepages1Gi     string `json:"hugepages-1Gi"`
				Hugepages2Mi     string `json:"hugepages-2Mi"`
				Memory           string `json:"memory"`
				Pods             string `json:"pods"`
			} `json:"capacity"`
			Conditions []struct {
				LastHeartbeatTime  time.Time `json:"lastHeartbeatTime"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				Message            string    `json:"message"`
				Reason             string    `json:"reason"`
				Status             string    `json:"status"`
				Type               string    `json:"type"`
			} `json:"conditions"`
			DaemonEndpoints struct {
				KubeletEndpoint struct {
					Port int `json:"Port"`
				} `json:"kubeletEndpoint"`
			} `json:"daemonEndpoints"`
			Images []struct {
				Names     []string `json:"names"`
				SizeBytes int      `json:"sizeBytes"`
			} `json:"images"`
			NodeInfo struct {
				Architecture            string `json:"architecture"`
				BootID                  string `json:"bootID"`
				ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
				KernelVersion           string `json:"kernelVersion"`
				KubeProxyVersion        string `json:"kubeProxyVersion"`
				KubeletVersion          string `json:"kubeletVersion"`
				MachineID               string `json:"machineID"`
				OperatingSystem         string `json:"operatingSystem"`
				OsImage                 string `json:"osImage"`
				SystemUUID              string `json:"systemUUID"`
			} `json:"nodeInfo"`
		} `json:"status"`
		Spec0 struct {
			PodCIDR    string   `json:"podCIDR"`
			PodCIDRs   []string `json:"podCIDRs"`
			ProviderID string   `json:"providerID"`
			Taints     []struct {
				Effect string `json:"effect"`
				Key    string `json:"key"`
			} `json:"taints"`
		} `json:"spec,omitempty"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}

func main() {
	// This should be the JSON data you want to parse.
	jsonData := exec.Command("kubectl", "get", "nodes", "-o", "json")
	var out bytes.Buffer
	jsonData.Stdout = &out
	err := jsonData.Run()
	if err != nil {
		fmt.Println("Error executing kubectl command:", err)
		return
	}

	// Unmarshal the JSON data into the struct.
	var data NodesInfo
	//err := json.Unmarshal([]byte(jsonData), &data)
	err = json.Unmarshal(out.Bytes(), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print information in a human-readable format.
	//	fmt.Println("API Version:", data.APIVersion)
	//	fmt.Println("Kind:", data.Kind)
	//	fmt.Println("Items:")
	for _, item := range data.Items {
		fmt.Printf("Cluster Name:         %s\n", item.Metadata.Annotations.ClusterXK8SIoClusterName)
		fmt.Printf("Instance Type:        %s\n", item.Metadata.Labels.NodeKubernetesIoInstanceType)
		fmt.Printf("Cluster Machine Name: %s\n", item.Metadata.Annotations.ClusterXK8SIoMachine)
		fmt.Printf("Cluster Node Name:    %s\n", item.Metadata.Name)
		fmt.Printf("OS Image:             %s\n", item.Status.NodeInfo.OsImage)
		fmt.Printf("Node Arch:            %s\n", item.Status.NodeInfo.Architecture)
		fmt.Printf("Kernel Version:       %s\n", item.Status.NodeInfo.KernelVersion)
		fmt.Printf("Node Args:            %s\n", item.Metadata.Annotations.Rke2IoNodeArgs)
		fmt.Printf("Pod Limits:           %s\n", item.Metadata.Annotations.ManagementCattleIoPodLimits)
		fmt.Printf("Pod Requests:         %s\n", item.Metadata.Annotations.ManagementCattleIoPodRequests)
		fmt.Println("\n\n")
		//		fmt.Println("  - API Version:", item.APIVersion)
		//		fmt.Println("    Kind:", item.Kind)
		//		fmt.Println("    Metadata:")
		//		fmt.Println("      Name:", item.Metadata.Name)
		//		fmt.Println("      Creation Timestamp:", item.Metadata.CreationTimestamp)
		//		fmt.Println("\tAnnotations:")
		//		fmt.Printf("\t\tcluster.x-k8s.io/cluster-name: %s\n", item.Metadata.Annotations.ClusterXK8SIoClusterName)
		//		fmt.Printf("\t\tcluster.x-k8s.io/cluster-namespace: %s\n", item.Metadata.Annotations.ClusterXK8SIoClusterNamespace)
		//		fmt.Printf("\t\tcluster.x-k8s.io/machine: %s\n", item.Metadata.Annotations.ClusterXK8SIoMachine)
		//		fmt.Printf("\t\tcluster.x-k8s.io/owner-kind: %s\n", item.Metadata.Annotations.ClusterXK8SIoOwnerKind)
		//		fmt.Printf("\t\tcluster.x-k8s.io/owner-name: %s\n", item.Metadata.Annotations.ClusterXK8SIoOwnerName)
		//		fmt.Printf("\t\tcsi.volume.kubernetes.io/nodeid: %s\n", item.Metadata.Annotations.CsiVolumeKubernetesIoNodeid)
		//		fmt.Printf("\t\tmanagement.cattle.io/pod-limits: %s\n", item.Metadata.Annotations.ManagementCattleIoPodLimits)
		//		fmt.Printf("\t\tmanagement.cattle.io/pod-requests: %s\n", item.Metadata.Annotations.ManagementCattleIoPodRequests)
		//		fmt.Printf("\t\tnode.alpha.kubernetes.io/ttl: %s\n", item.Metadata.Annotations.NodeAlphaKubernetesIoTTL)
		//		fmt.Printf("\t\trke2.io/hostname: %s\n", item.Metadata.Annotations.Rke2IoHostname)
		//		fmt.Printf("\t\trke2.io/internal-ip: %s\n", item.Metadata.Annotations.Rke2IoInternalIP)
		//		fmt.Printf("\t\trke2.io/node-args: %s\n", item.Metadata.Annotations.Rke2IoNodeArgs)
		//		fmt.Printf("\t\trke2.io/node-config-hash: %s\n", item.Metadata.Annotations.Rke2IoNodeConfigHash)
		//		fmt.Printf("\t\trke2.io/node-env: %s\n", item.Metadata.Annotations.Rke2IoNodeEnv)
		//		fmt.Printf("\t\tvolumes.kubernetes.io/controller-managed-attach-detach: %s\n", item.Metadata.Annotations.VolumesKubernetesIoControllerManagedAttachDetach)

		//		fmt.Println("\tLabels:")
		//		fmt.Printf("\t\tbeta.kubernetes.io/arch: %s\n", item.Metadata.Labels.BetaKubernetesIoArch)
		//		fmt.Printf("\t\tbeta.kubernetes.io/instance-type: %s\n", item.Metadata.Labels.BetaKubernetesIoInstanceType)
		//		fmt.Printf("\t\tbeta.kubernetes.io/os: %s\n", item.Metadata.Labels.BetaKubernetesIoOs)
		//		fmt.Printf("\t\tkubernetes.io/arch: %s\n", item.Metadata.Labels.KubernetesIoArch)
		//		fmt.Printf("\t\tkubernetes.io/hostname: %s\n", item.Metadata.Labels.KubernetesIoHostname)
		//		fmt.Printf("\t\tkubernetes.io/os: %s\n", item.Metadata.Labels.KubernetesIoOs)
		//		fmt.Printf("\t\tnode-role.kubernetes.io/worker: %s\n", item.Metadata.Labels.NodeRoleKubernetesIoWorker)
		//		fmt.Printf("\t\tnode.kubernetes.io/instance-type: %s\n", item.Metadata.Labels.NodeKubernetesIoInstanceType)
		//		fmt.Printf("\t\tplan.upgrade.cattle.io/system-agent-upgrader: %s\n", item.Metadata.Labels.PlanUpgradeCattleIoSystemAgentUpgrader)
		//		fmt.Printf("\t\trke.cattle.io/machine: %s\n", item.Metadata.Labels.RkeCattleIoMachine)

		// Continue printing other fields in a similar way
	}
}
