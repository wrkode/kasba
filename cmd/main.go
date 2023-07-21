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

	templateData.AppsInfo, err = kubeconfig.GetWorkloads()
	if err != nil {
		fmt.Printf("Error getting apps %v\n:", err)
	}

	// Write call to GetPersistentVolumes()

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
