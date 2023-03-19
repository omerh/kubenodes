package client

import (
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func LoadAPIClient(pathFlag string) *kubernetes.Clientset {
	path := kubeConfigPath(pathFlag)
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func kubeConfigPath(pathFlag string) string {
	var kubeconfig string
	if pathFlag != "" {
		kubeconfig = pathFlag
	} else {
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
	return kubeconfig
}
