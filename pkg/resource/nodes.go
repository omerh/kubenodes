package resource

import (
	"context"
	"fmt"
	"sort"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeInfo struct {
	Name         string
	CapacityType string
	InstanceType string
	InstanceArch string
	InstanceAZ   string
	Pod          []string
}

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

func DescribeNode(apiClient *kubernetes.Clientset, nodeName string) NodeInfo {
	node, _ := apiClient.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	capacityType := "on-demand"
	if node.Labels["karpenter.sh/capacity-type"] == "spot" {
		capacityType = "spot"
	}

	n := NodeInfo{
		Name:         nodeName,
		CapacityType: capacityType,
		InstanceType: node.Labels["node.kubernetes.io/instance-type"],
		InstanceArch: node.Labels["kubernetes.io/arch"],
		InstanceAZ:   node.Labels["topology.kubernetes.io/zone"],
	}
	return n
}

func NodeMapToNodes(nodeMap map[string][]string) []NodeInfo {
	var nodes []NodeInfo

	for node, pods := range nodeMap {
		node := NodeInfo{
			Name: node,
			Pod:  pods,
		}
		nodes = append(nodes, node)
	}

	return nodes
}

func UpdateNodeInfoSlice(apiClient *kubernetes.Clientset, nodes []NodeInfo) []NodeInfo {
	for i, n := range nodes {
		info := DescribeNode(apiClient, n.Name)
		nodes[i].CapacityType = info.CapacityType
		nodes[i].InstanceType = info.InstanceType
		nodes[i].InstanceArch = info.InstanceArch
		nodes[i].InstanceAZ = info.InstanceAZ
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})

	return nodes
}
