package queuescheduler

import (
	"strconv"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/component-helpers/scheduling/corev1"
	"k8s.io/klog/v2"
	v1qos "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Implements framework.Plugin
type QueueScheduler struct{}

var _ framework.QueueSortPlugin = &QueueScheduler{}

const PluginName = "QueueScheduler"

// To implement framework.Plugin
func (qs QueueScheduler) Name() string {
	// Name is the name of the plugin used in the Registry and configurations.
	return PluginName
}

// Check if the pod has a label called priority and if it does, then give this pod preference
func GetPodPriority(pod *v1.Pod) int32 {
	if labelValue, ok := pod.ObjectMeta.Labels["priority"]; ok {
		priority, err := strconv.Atoi(strings.Split(labelValue, "-")[1])

		if err != nil {
			return int32(priority)
		}

	}

	return corev1.PodPriority(pod)
}

// Less is the function used by the activeQ heap algorithm to sort pods.
// It sorts pods based on their priorities. When the priorities are equal, it uses
// the Pod QoS classes to break the tie.
func (qs *QueueScheduler) Less(pInfo1, pInfo2 *framework.QueuedPodInfo) bool {
	// k := "label = string-index"

	klog.Info("\n\n-------------------------------------------------------\n")
	klog.Infof("Made it into the Less function of the custon scheduler \n")
	klog.Info("-------------------------------------------------------\n\n")

	priority1 := GetPodPriority(pInfo1.Pod)
	priority2 := GetPodPriority(pInfo2.Pod)

	return (priority1 > priority2) || ((priority1 == priority2) && compareQualityOfService(pInfo1.Pod, pInfo2.Pod))
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
