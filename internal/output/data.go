package output

import (
	"github.com/wrkode/kasba/internal/nodeinfo"
	"github.com/wrkode/kasba/internal/util"
)

type TemplateData struct {
	CreatedAt              string
	BOMFormat              string
	Version                string
	NodeInfo               nodeinfo.NodesInfo
	NetworkPlugin          string
	Longhorn               bool
	Monitoring             bool
	WorkloadInfo           util.WorkloadInfo
	StorageClass           []util.StorageClassItem
	PersistentVolumes      []util.PersistentVolumeItem
	PersistentVolumeClaims []util.PersistentVolumeClaimItem
	ConfigMaps             []util.ConfigMapItem
	Services               []util.ServiceItem
	Ingresses              []util.IngressItem
	ClusterRoles           []util.ClusterRoleItem
	ClusterRoleBindings    []util.ClusterRoleBindingItem
	ServiceAccounts        []util.ServiceAccountItem
	Errors                 util.Errors
}
