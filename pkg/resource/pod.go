package resource

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func GetPods(apiClient *kubernetes.Clientset, namespace string, labelSlice []string) v1.PodList {
	// create label selector according to app=
	// currently not free form labels for comfort
	var listOptions metav1.ListOptions
	if len(labelSlice) > 0 {
		labels, err := labels.NewRequirement("app", "in", labelSlice)
		if err != nil {
			panic(err)
		}
		listOptions = metav1.ListOptions{
			LabelSelector: labels.String(),
		}
	} else {
		listOptions = metav1.ListOptions{}
	}
	// list pods
	pods, _ := apiClient.CoreV1().Pods(namespace).List(context.TODO(), listOptions)

	var allPods v1.PodList
	allPods.Items = append(allPods.Items, pods.Items...)
	return allPods
}

func MakeUniquePodsOnNode(pods v1.PodList) map[string][]string {
	nodePodMap := make(map[string][]string)
	for _, p := range pods.Items {
		nodePodMap[p.Spec.NodeName] = append(nodePodMap[p.Spec.NodeName], p.Name)
	}
	return nodePodMap
}
