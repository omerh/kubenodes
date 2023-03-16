package resource

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNodes(apiClient *kubernetes.Clientset) {
	nodes, _ := apiClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	// fmt.Println(nodes)
	for _, n := range nodes.Items {
		nodeType := n.Labels["karpenter.sh/capacity-type"]
		if nodeType != "spot" {
			nodeType = "on-demand"
		}
		fmt.Printf("%s\t\t%s\t\t%s\n", n.Name, nodeType, n.Labels["node.kubernetes.io/instance-type"])
		fmt.Println("======================================================================================")
		// nsList := getNamespaces(clientset)
		// pods := getPodsOnNode(clientset, n.Name, nsList)
		// for _, p := range pods.Items {
		// fmt.Printf("%s\t\t%s\n", p.Name, p.Namespace)
		// }
		// fmt.Println()
	}
}
