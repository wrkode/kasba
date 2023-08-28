package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wrkode/kasba/internal/output"
	"github.com/wrkode/kasba/internal/util"
	"os"
	"time"

	_ "embed"
)

const (
	SUSE      = "SUSE Software Solutions"
	BOMFormat = "kasba"
)

var (
	createdAt    = time.Now().Format(time.RFC850)
	Version      = "dev"
	kubeconfig   util.KubeConfig
	errors       util.Errors
	templateData output.TemplateData
	//kasbaID   = uuid.New().String()
	showVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "KASBA",
	Short: "Kubernetes As-Built Assessment",
	RunE:  executeRun,
}

func GetInfo(cmd *cobra.Command) {
	var err error

	templateData.CreatedAt = createdAt
	templateData.BOMFormat = BOMFormat
	templateData.Version = Version

	err = kubeconfig.GetKubeConfigPath(cmd)
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

	templateData.NetworkPolicies, err = kubeconfig.GetAllNetworkPolicies()
	if err != nil {
		fmt.Printf("Error getting Network Policies %v\n", err)
	}
}

func init() {
	kubeconfig.BindFlags(rootCmd)
	rootCmd.PersistentFlags().BoolVar(&showVersion, "version", false, "display the version number")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func executeRun(cmd *cobra.Command, args []string) error {
	if err := kubeconfig.GetKubeConfigPath(cmd); err != nil {
		return err
	}
	if showVersion {
		fmt.Printf("KASBA Version: %s\n", Version)
		return nil
	}

	GetInfo(cmd)

	templateData.Errors = errors

	if err := output.AsText(templateData); err != nil {
		return fmt.Errorf("unable to parse template: %v", err)
	}

	return nil
}
