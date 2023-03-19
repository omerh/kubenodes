package resource

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNamespaces(apiClient *kubernetes.Clientset) []v1.Namespace {
	// listOptions := metav1.ListOptions{
	// 	FieldSelector: "metadata.name!=kube-system",
	// }
	ns, _ := apiClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	return ns.Items
}
