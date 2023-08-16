package util

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	rbacv1 "k8s.io/api/rbac/v1"
)

// ClusterNodes represents the structure of data to be marshalled into JSON
type ClusterNodes struct {
	Cluster string    `json:"cluster"` // Name of the cluster
	Nodes   []v1.Node `json:"nodes"`   // List of nodes in the cluster
}

type KubeConfig struct {
	kubeconfig   *string
	config       *rest.Config
	clientset    *kubernetes.Clientset
	workloadlist []WorkloadListItem
}

type WorkloadListItem struct {
	Name      string
	Namespace string
	Type      string
}

type WorkloadInfoAppType struct {
	WorkloadType string
	Workloads    []string
}

type WorkloadInfoNamespace struct {
	Namespace     string
	WorkloadTypes []WorkloadInfoAppType
}

type WorkloadInfo struct {
	Namespaces []WorkloadInfoNamespace
}

type StorageClassItem struct {
	Name        string
	Provisioner string
	Parameters  map[string]string
}

type PersistentVolumeItem struct {
	Name              string
	Namespace         string
	Type              string
	Size              resource.Quantity
	AccessModes       []v1.PersistentVolumeAccessMode
	ReclamationPolicy v1.PersistentVolumeReclaimPolicy
}

type PersistentVolumeClaimItem struct {
	Namespace    string
	Name         string
	Status       v1.PersistentVolumeClaimPhase
	Volume       string
	Capacity     resource.Quantity
	AccessModes  []v1.PersistentVolumeAccessMode
	StorageClass string
	Age          metav1.Time
}

type ConfigMapItem struct {
	Namespace string
	Name      string
	Data      map[string]string
	Age       metav1.Time
}

type ServiceItem struct {
	Namespace   string
	Name       string
	Type       v1.ServiceType
	ClusterIP  string
	ExternalIP string // This can be a list or a single IP. Improvement Required to handle multiple IPs.
	Ports      []v1.ServicePort
	Age        int
}

// IngressBackendDetail is used to capture the backend service and port
type IngressBackendDetail struct {
	ServiceName string
	ServicePort string
}

// IngressRuleDetail captures the hosts and paths for a rule
type IngressRuleDetail struct {
	Host  string
	Paths []string
}

// IngressItem represents an Ingress in the cluster
type IngressItem struct {
	Namespace string
	Name      string
	Hosts     []IngressRuleDetail
	DefaultBackend IngressBackendDetail // This will capture the default backend, if any
	Addresses      []string
	Age       int
}

// ClusterRoleItem simplified to just verbs for now, can be expanded to include resources, API groups, etc.
type ClusterRoleItem struct {
	Name  string
	Verbs []string
}

type ClusterRoleBindingItem struct {
	Name     string
	RoleName string // Name of the ClusterRole that this ClusterRoleBinding refers to
	Subjects []rbacv1.Subject // List of subjects associated with this ClusterRoleBinding
}