package queuescheduler

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/component-helpers/scheduling/corev1"
	v1qos "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Implements framework.Plugin
type QueueScheduler struct{}

func (qs QueueScheduler) Name() string {
	// Name is the name of the plugin used in the Registry and configurations.
	return "CustomQueueScheduler"
}

func GetPodPriority(pod *v1.Pod) string {
	if value, ok := pod.ObjectMeta.Labels["priority"]; ok {
		return value
	}

	return ""
}

// Less is the function used by the activeQ heap algorithm to sort pods.
// It sorts pods based on their priorities. When the priorities are equal, it uses
// the Pod QoS classes to break the tie.
func (*QueueScheduler) Less(pInfo1, pInfo2 *framework.QueuedPodInfo) bool {
	// k := "label = string-index"

	p1 := corev1.PodPriority(pInfo1.Pod)
	p2 := corev1.PodPriority(pInfo2.Pod)

	return (p1 > p2) || ((p1 == p2) && compareQualityOfService(pInfo1.Pod, pInfo2.Pod))
}

func compareQualityOfService(p1, p2 *v1.Pod) bool {
	p1QOS, p2QOS := v1qos.GetPodQOS(p1), v1qos.GetPodQOS(p2)

	if p1QOS == v1.PodQOSGuaranteed {
		return true
	}

	if p1QOS == v1.PodQOSBurstable {
		return p2QOS != v1.PodQOSGuaranteed
	}

	return p2QOS == v1.PodQOSBestEffort
}

// New initializes a new plugin and returns it.
func New(obj runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &QueueScheduler{}, nil
}
