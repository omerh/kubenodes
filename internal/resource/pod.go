package resource

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPods(apiClient *kubernetes.Clientset, namespace string, deployment string) v1.PodList {
	var allPods v1.PodList
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", deployment),
	}
	pods, _ := apiClient.CoreV1().Pods(namespace).List(context.TODO(), listOptions)
	allPods.Items = append(allPods.Items, pods.Items...)
	return allPods
}
