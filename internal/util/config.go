package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

// GetKubeConfigPath gathers/uses active kubeconfig
func (k *KubeConfig) GetKubeConfigPath(cmd *cobra.Command) error {
	var err error
	var kubeconfigPath string

	// Check flag first
	if k.KubeconfigFlag != "" {
		kubeconfigPath = k.KubeconfigFlag
	} else if envKubeConfig := os.Getenv("KUBECONFIG"); envKubeConfig != "" {
		// Then check environment variable
		kubeconfigPath = envKubeConfig
	} else if home := homedir.HomeDir(); home != "" {
		// Finally, default to ~/.kube/config
		kubeconfigPath = filepath.Join(home, ".kube", "config")
	}

	if kubeconfigPath == "" {
		return fmt.Errorf("no kubeconfig path provided")
	}

	k.config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}

	k.clientset, err = kubernetes.NewForConfig(k.config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %v", err)
	}

	return nil
}
