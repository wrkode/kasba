package cmd

import (
	"flag"
	"fmt"
	"github.com/wrkode/kasba/internal/output"
	"github.com/wrkode/kasba/internal/util"
	"log"
	"time"

	_ "embed"
)

const (
	SUSE      = "SUSE Software Solutions"
	BOMFormat = "kasba"
)

var (
	createdAt    = time.Now().Format(time.RFC850)
	Version      = ""
	kubeconfig   util.KubeConfig
	errors       util.Errors
	templateData output.TemplateData
	//kasbaID   = uuid.New().String()
)

func GetInfo() {
	var err error

	templateData.CreatedAt = createdAt
	templateData.BOMFormat = BOMFormat
	templateData.Version = Version

	err = kubeconfig.GetKubeConfigPath()
	if errors.Add(err, true) {
		return // return because of fatal error
	}

	// Get Node Info
	templateData.NodeInfo, err = kubeconfig.GetNodeInfo()
	if errors.Add(err, true) {
		return // return because of fatal error
	}

	if len(templateData.NodeInfo.Items) == 0 {
		errors.Add(fmt.Errorf("unable to get nodeinfo"), true)
		return // return because of fatal error
	}

	templateData.NetworkPlugin, err = kubeconfig.GetNetworkPluginPodName()
	errors.Add(err, false)

	templateData.Longhorn, err = kubeconfig.NamespaceExists("longhorn-system")
	errors.Add(err, false)

	if err != nil {
		fmt.Println("Error checking if Longhorn is installed: ", err)
	}

	templateData.Monitoring, err = kubeconfig.NamespaceExists("cattle-monitoring-system")
	errors.Add(err, false)

	if err != nil {
		fmt.Println("Error checking if Rancher Monitoring is installed: ", err)
	}

	templateData.WorkloadInfo, err = kubeconfig.GetWorkloads()
	if err != nil {
		fmt.Printf("Error getting apps %v\n:", err)
	}

	templateData.StorageClass, err = kubeconfig.GetStorageClasses()
	if err != nil {
		fmt.Printf("Error getting Storage Classes %v\n:", err)
	}

	templateData.PersistentVolumes, err = kubeconfig.GetPersistentVolumes()
	if err != nil {
		fmt.Printf("Error getting Persistent Volumes %v\n", err)
	}

	templateData.PersistentVolumeClaims, err = kubeconfig.GetPersistentVolumeClaims()
	if err != nil {
		fmt.Printf("Error getting Persistent Volume Claims %v\n", err)
	}

	templateData.ConfigMaps, err = kubeconfig.GetConfigMaps()
	if err != nil {
		fmt.Printf("Error getting ConfigMaps %v\n", err)
	}

	templateData.Services, err = kubeconfig.GetAllServices()
	if err != nil {
		fmt.Printf("Error getting Serices %v\n", err)
	}

	templateData.Ingresses, err = kubeconfig.GetAllIngresses()
	if err != nil {
		fmt.Printf("Error getting Ingresses %v\n", err)
	}

	templateData.ClusterRoles, err = kubeconfig.GetAllClusterRoles()
	if err != nil {
		fmt.Printf("Error getting ClusterRoles %v\n", err)
	}

	templateData.ClusterRoleBindings, err = kubeconfig.GetAllClusterRoleBindings()
	if err != nil {
		fmt.Printf("Error getting ClusterRoleBindings %v\n", err)
	}

	templateData.ServiceAccounts, err = kubeconfig.GetAllServiceAccounts()
	if err != nil {
		fmt.Printf("Error getting Service Accounts %v\n", err)
	}
}

func Run() {
	flag.Parse()

	GetInfo()

	templateData.Errors = errors

	err := output.AsText(templateData)
	if err != nil {
		log.Fatalf("Unable to Parse Template: %v", err)
	}
}
