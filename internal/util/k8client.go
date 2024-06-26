package util

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wrkode/kasba/internal/nodeinfo"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var kubeconfigFlag = flag.String("kubeconfig", "", "(optional) absolute path to the kubeconfig file")
var VersionFlag = flag.Bool("version", false, "print version information and exit")

func (a *WorkloadInfo) Add(namespace string, appType string, name string) {
	if len(a.Namespaces) == 0 {
		a.Namespaces = []WorkloadInfoNamespace{{
			Namespace: namespace,
			WorkloadTypes: []WorkloadInfoAppType{{
				WorkloadType: appType,
				Workloads:    []string{name},
			}},
		}}
	} else {
		nsFound := false
		for nsIndex, ns := range a.Namespaces {
			if namespace == ns.Namespace {
				nsFound = true
				atFound := false
				for atIndex, at := range ns.WorkloadTypes {
					if appType == at.WorkloadType {
						atFound = true
						a.Namespaces[nsIndex].WorkloadTypes[atIndex].Workloads = append(ns.WorkloadTypes[atIndex].Workloads, name)
					}
				}
				// new WorkloadType
				if !atFound {
					a.Namespaces[nsIndex].WorkloadTypes = append(a.Namespaces[nsIndex].WorkloadTypes,
						WorkloadInfoAppType{
							WorkloadType: appType,
							Workloads:    []string{name},
						},
					)
				}
			}
		}
		// new Namespace
		if !nsFound {
			a.Namespaces = append(a.Namespaces, WorkloadInfoNamespace{
				Namespace: namespace,
				WorkloadTypes: []WorkloadInfoAppType{{
					WorkloadType: appType,
					Workloads:    []string{name},
				}},
			},
			)
		}

	}
}

// GetKubeConfigPath gathers/uses --kubeconfig, KUBECONFIG EnvVar or active kubeconfig context from ~/.kube/config
func (k *KubeConfig) GetKubeConfigPath() error {
	var err error

	// If the flag is set, it takes the highest priority.
	if *kubeconfigFlag != "" {
		k.config, err = clientcmd.BuildConfigFromFlags("", *kubeconfigFlag)
	} else if envKubeconfig := os.Getenv("KUBECONFIG"); envKubeconfig != "" {
		// If the KUBECONFIG environment variable is set, use it.
		k.config, err = clientcmd.BuildConfigFromFlags("", envKubeconfig)
	} else {
		// Otherwise, fall back to the default kubeconfig location.
		if home := homedir.HomeDir(); home != "" {
			k.config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
		} else {
			k.config, err = rest.InClusterConfig()
		}
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

// GetDeployments lists the deployments in all namespaces and returns them with NAMES and NAMESPACE.
func (k *KubeConfig) GetDeployments() {
	list, _ := k.clientset.AppsV1().Deployments(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	for _, listItem := range list.Items {
		k.workloadlist = append(k.workloadlist, WorkloadListItem{
			Name:      listItem.Name,
			Namespace: listItem.Namespace,
			Type:      "Deployments",
		})
	}
}

// GetDaemonSets lists the daemonsets in all namespaces and returns them with NAMES and NAMESPACE.
func (k *KubeConfig) GetDaemonSets() {
	list, _ := k.clientset.AppsV1().DaemonSets(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	for _, listItem := range list.Items {
		k.workloadlist = append(k.workloadlist, WorkloadListItem{
			Name:      listItem.Name,
			Namespace: listItem.Namespace,
			Type:      "DaemonSets",
		})
	}
}

// GetStatefulSets lists the statefulsets in all namespaces and returns them with NAMES and NAMESPACE.
func (k *KubeConfig) GetStatefulSets() {
	list, _ := k.clientset.AppsV1().StatefulSets(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	for _, listItem := range list.Items {
		k.workloadlist = append(k.workloadlist, WorkloadListItem{
			Name:      listItem.Name,
			Namespace: listItem.Namespace,
			Type:      "StatefulSets",
		})
	}
}

// GetWorkloads List all the apps running on the cluster, sorted by namespace and type
func (k *KubeConfig) GetWorkloads() (WorkloadInfo, error) {
	var workloadInfo WorkloadInfo
	k.GetDeployments()
	k.GetDaemonSets()
	k.GetStatefulSets()
	for _, a := range k.workloadlist {
		workloadInfo.Add(a.Namespace, a.Type, a.Name)
	}
	return workloadInfo, nil

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

// GetNodeInfo get nodes info from active context
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

// GetStorageClasses lists the Storage Classes in the cluster and returns them.
func (k *KubeConfig) GetStorageClasses() ([]StorageClassItem, error) {
	list, err := k.clientset.StorageV1().StorageClasses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var storageClasses []StorageClassItem
	for _, listItem := range list.Items {
		sc := StorageClassItem{
			Name:        listItem.Name,
			Provisioner: listItem.Provisioner,
			Parameters:  listItem.Parameters,
		}
		storageClasses = append(storageClasses, sc)
	}
	return storageClasses, nil
}

// GetPersistentVolumes lists the Persistent Volumes available in the cluster and returns them.
func (k *KubeConfig) GetPersistentVolumes() ([]PersistentVolumeItem, error) {
	list, err := k.clientset.CoreV1().PersistentVolumes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var persistentVolumes []PersistentVolumeItem
	for _, listItem := range list.Items {
		size, exists := listItem.Spec.Capacity["storage"]
		if !exists {
			continue // skip this PV if it doesn't have a storage size defined
		}

		pv := PersistentVolumeItem{
			Name:              listItem.Name,
			Type:              "PersistentVolumes",
			Size:              size,
			AccessModes:       listItem.Spec.AccessModes,
			ReclamationPolicy: listItem.Spec.PersistentVolumeReclaimPolicy,
		}
		persistentVolumes = append(persistentVolumes, pv)
	}
	return persistentVolumes, nil
}

// GetPersistentVolumeClaims lists all Persistent Volume Claims across all namespaces.
func (k *KubeConfig) GetPersistentVolumeClaims() ([]PersistentVolumeClaimItem, error) {
	list, err := k.clientset.CoreV1().PersistentVolumeClaims(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var persistentVolumeClaims []PersistentVolumeClaimItem
	for _, listItem := range list.Items {
		pvc := PersistentVolumeClaimItem{
			Namespace:    listItem.Namespace,
			Name:         listItem.Name,
			Status:       listItem.Status.Phase,
			Volume:       listItem.Spec.VolumeName,
			Capacity:     listItem.Status.Capacity[v1.ResourceStorage],
			AccessModes:  listItem.Status.AccessModes,
			StorageClass: *listItem.Spec.StorageClassName,
			Age:          listItem.CreationTimestamp,
		}
		persistentVolumeClaims = append(persistentVolumeClaims, pvc)
	}
	return persistentVolumeClaims, nil
}

// GetConfigMaps lists all ConfigMaps across all namespaces.
func (k *KubeConfig) GetConfigMaps() ([]ConfigMapItem, error) {
	list, err := k.clientset.CoreV1().ConfigMaps(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var configMaps []ConfigMapItem
	for _, listItem := range list.Items {
		cm := ConfigMapItem{
			Namespace: listItem.Namespace,
			Name:      listItem.Name,
			Data:      listItem.Data,
			Age:       listItem.CreationTimestamp,
		}
		configMaps = append(configMaps, cm)
	}
	return configMaps, nil
}

// GetAllServices lists all Services across all namespaces.
func (k *KubeConfig) GetAllServices() ([]ServiceItem, error) {
	svcList, err := k.clientset.CoreV1().Services(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var services []ServiceItem
	for _, svc := range svcList.Items {
		var externalIP string
		if len(svc.Status.LoadBalancer.Ingress) > 0 {
			externalIP = svc.Status.LoadBalancer.Ingress[0].IP
		}

		// Calculate Age in seconds
		ageInSeconds := time.Since(svc.ObjectMeta.CreationTimestamp.Time).Seconds()
		// Convert Age to days
		ageInDays := int(ageInSeconds / (60 * 60 * 24)) // convert seconds to days

		serviceItem := ServiceItem{
			Namespace:  svc.Namespace,
			Name:       svc.Name,
			Type:       svc.Spec.Type,
			ClusterIP:  svc.Spec.ClusterIP,
			ExternalIP: externalIP,
			Ports:      svc.Spec.Ports,
			Age:        ageInDays,
		}
		services = append(services, serviceItem)
	}
	return services, nil
}

// GetAllIngresses lists all Ingresses across all namespaces.
func (k *KubeConfig) GetAllIngresses() ([]IngressItem, error) {
	ingList, err := k.clientset.NetworkingV1().Ingresses(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var ingresses []IngressItem
	for _, ing := range ingList.Items {
		var rules []IngressRuleDetail
		for _, rule := range ing.Spec.Rules {
			var paths []string
			for _, path := range rule.HTTP.Paths {
				paths = append(paths, path.Path)
			}
			rules = append(rules, IngressRuleDetail{
				Host:  rule.Host,
				Paths: paths,
			})
		}

		var defaultBackend IngressBackendDetail
		if ing.Spec.DefaultBackend != nil {
			defaultBackend = IngressBackendDetail{
				ServiceName: ing.Spec.DefaultBackend.Service.Name,
				ServicePort: ing.Spec.DefaultBackend.Service.Port.String(),
			}
		}

		// Fetch addresses
		var addresses []string
		for _, addr := range ing.Status.LoadBalancer.Ingress {
			if addr.IP != "" {
				addresses = append(addresses, addr.IP)
			}
			if addr.Hostname != "" {
				addresses = append(addresses, addr.Hostname)
			}
		}

		// Calculate Age
		ageInSeconds := time.Since(ing.ObjectMeta.CreationTimestamp.Time).Seconds()
		ageInDays := int(ageInSeconds / (60 * 60 * 24))

		ingressItem := IngressItem{
			Namespace:      ing.Namespace,
			Name:           ing.Name,
			Hosts:          rules,
			DefaultBackend: defaultBackend,
			Addresses:      addresses, // Added this
			Age:            ageInDays,
		}
		ingresses = append(ingresses, ingressItem)
	}
	return ingresses, nil
}

// GetAllClusterRoles lists all ClusterRoles defined.
func (k *KubeConfig) GetAllClusterRoles() ([]ClusterRoleItem, error) {
	roles, err := k.clientset.RbacV1().ClusterRoles().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var roleItems []ClusterRoleItem
	for _, role := range roles.Items {
		verbs := []string{} // Collecting all verbs
		for _, rule := range role.Rules {
			for _, verb := range rule.Verbs {
				verbs = append(verbs, verb)
			}
		}
		roleItem := ClusterRoleItem{
			Name:  role.Name,
			Verbs: verbs,
		}
		roleItems = append(roleItems, roleItem)
	}
	return roleItems, nil
}

// GetAllClusterRoleBindings lists all ClusterRolesBindings defined.
func (k *KubeConfig) GetAllClusterRoleBindings() ([]ClusterRoleBindingItem, error) {
	crbList, err := k.clientset.RbacV1().ClusterRoleBindings().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var clusterRoleBindings []ClusterRoleBindingItem
	for _, crb := range crbList.Items {
		crbItem := ClusterRoleBindingItem{
			Name:     crb.Name,
			RoleName: crb.RoleRef.Name,
			Subjects: crb.Subjects,
		}
		clusterRoleBindings = append(clusterRoleBindings, crbItem)
	}
	return clusterRoleBindings, nil
}

// GetAllServiceAccounts lists all Service Accounts defined
func (k *KubeConfig) GetAllServiceAccounts() ([]ServiceAccountItem, error) {
	serviceAccountList, err := k.clientset.CoreV1().ServiceAccounts(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var serviceAccounts []ServiceAccountItem
	for _, sa := range serviceAccountList.Items {
		age := time.Since(sa.ObjectMeta.CreationTimestamp.Time)
		ageStr := ""

		// Convert age to hours and days
		totalHours := int(age.Hours())
		if totalHours < 24 {
			ageStr = fmt.Sprintf("%dh", totalHours)
		} else {
			days := totalHours / 24
			ageStr = fmt.Sprintf("%dd", days)
		}

		serviceAccount := ServiceAccountItem{
			Name:      sa.Name,
			Namespace: sa.Namespace,
			Secrets:   len(sa.Secrets),
			Age:       ageStr,
		}
		serviceAccounts = append(serviceAccounts, serviceAccount)
	}
	return serviceAccounts, nil
}

// GetAllNetworkPolicies Implements a method to fetch all network policies
func (k *KubeConfig) GetAllNetworkPolicies() ([]NetworkPolicyItem, error) {
	netPolList, err := k.clientset.NetworkingV1().NetworkPolicies(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var netPolicies []NetworkPolicyItem
	for _, netPol := range netPolList.Items {
		age := time.Since(netPol.ObjectMeta.CreationTimestamp.Time)

		// Convert age to hours or days
		var ageStr string
		if age.Hours() < 24 {
			ageStr = fmt.Sprintf("%.0fh", age.Hours())
		} else {
			ageStr = fmt.Sprintf("%.0fd", age.Hours()/24)
		}

		netPolItem := NetworkPolicyItem{
			Name:        netPol.Name,
			Namespace:   netPol.Namespace,
			PodSelector: netPol.Spec.PodSelector,
			Ingress:     netPol.Spec.Ingress,
			Egress:      netPol.Spec.Egress,
			PolicyTypes: netPol.Spec.PolicyTypes,
			Age:         ageStr,
		}
		netPolicies = append(netPolicies, netPolItem)
	}

	return netPolicies, nil
}
