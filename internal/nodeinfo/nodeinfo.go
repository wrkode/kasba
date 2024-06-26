package nodeinfo

import "time"

type NodesInfo struct {
	Cluster string `json:"cluster"`
	Items   []struct {
		Metadata struct {
			Name              string    `json:"name"`
			UID               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
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
			Finalizers    []string `json:"finalizers"`
			ManagedFields []struct {
				Manager    string    `json:"manager"`
				Operation  string    `json:"operation"`
				APIVersion string    `json:"apiVersion"`
				Time       time.Time `json:"time"`
				FieldsType string    `json:"fieldsType"`
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
		} `json:"metadata"`
		Spec struct {
			PodCIDR    string   `json:"podCIDR"`
			PodCIDRs   []string `json:"podCIDRs"`
			ProviderID string   `json:"providerID"`
		} `json:"spec,omitempty"`
		Status struct {
			Capacity struct {
				CPU              string `json:"cpu"`
				EphemeralStorage string `json:"ephemeral-storage"`
				Hugepages1Gi     string `json:"hugepages-1Gi"`
				Hugepages2Mi     string `json:"hugepages-2Mi"`
				Memory           string `json:"memory"`
				Pods             string `json:"pods"`
			} `json:"capacity"`
			Allocatable struct {
				CPU              string `json:"cpu"`
				EphemeralStorage string `json:"ephemeral-storage"`
				Hugepages1Gi     string `json:"hugepages-1Gi"`
				Hugepages2Mi     string `json:"hugepages-2Mi"`
				Memory           string `json:"memory"`
				Pods             string `json:"pods"`
			} `json:"allocatable"`
			Conditions []struct {
				Type               string    `json:"type"`
				Status             string    `json:"status"`
				LastHeartbeatTime  time.Time `json:"lastHeartbeatTime"`
				LastTransitionTime time.Time `json:"lastTransitionTime"`
				Reason             string    `json:"reason"`
				Message            string    `json:"message"`
			} `json:"conditions"`
			Addresses []struct {
				Type    string `json:"type"`
				Address string `json:"address"`
			} `json:"addresses"`
			DaemonEndpoints struct {
				KubeletEndpoint struct {
					Port int `json:"Port"`
				} `json:"kubeletEndpoint"`
			} `json:"daemonEndpoints"`
			NodeInfo struct {
				MachineID               string `json:"machineID"`
				SystemUUID              string `json:"systemUUID"`
				BootID                  string `json:"bootID"`
				KernelVersion           string `json:"kernelVersion"`
				OsImage                 string `json:"osImage"`
				ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
				KubeletVersion          string `json:"kubeletVersion"`
				KubeProxyVersion        string `json:"kubeProxyVersion"`
				OperatingSystem         string `json:"operatingSystem"`
				Architecture            string `json:"architecture"`
			} `json:"nodeInfo"`
			Images []struct {
				Names     []string `json:"names"`
				SizeBytes int      `json:"sizeBytes"`
			} `json:"images"`
		} `json:"status"`
		Spec0 struct {
			PodCIDR    string   `json:"podCIDR"`
			PodCIDRs   []string `json:"podCIDRs"`
			ProviderID string   `json:"providerID"`
			Taints     []struct {
				Key    string `json:"key"`
				Effect string `json:"effect"`
			} `json:"taints"`
		} `json:"spec,omitempty"`
	} `json:"nodes"`
}
