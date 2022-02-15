package queuescheduler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	v1qos "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Implements framework.Plugin
type QueueScheduler struct{}

var _ framework.QueueSortPlugin = &QueueScheduler{}

const Name = "QueueScheduler"

func log(msg string, logType string) {
	switch logType {
	case "info":
		klog.Info(msg, "\n\n")

	case "error":
		klog.Error(msg, "\n\n")
	}
}

// To implement framework.Plugin
func (qs QueueScheduler) Name() string {
	// Name is the name of the plugin used in the Registry and configurations.
	return Name
}

// Check if the pod has a label called priority and if it does, then give this pod preference
func GetPodPriority(podInfo *framework.QueuedPodInfo, currentUnixTime int64) (int32, bool) {
	pod := podInfo.Pod

	if pod.ObjectMeta.Labels == nil {
		return int32(currentUnixTime - podInfo.Timestamp.Unix()), false
	}

	// 1. check if label "priority" exists
	// 2. check if the value of that is of the format - priority-number
	if labelValue, ok := pod.ObjectMeta.Labels["priority"]; ok {
		if len(labelValue) > 0 {
			prioritySplit := strings.Split(labelValue, "-")

			if prioritySplit[0] == "priority" {
				priority, err := strconv.Atoi(prioritySplit[1])

				if err == nil {
					return int32(priority), true
				}
			}
		}
	}

	// if one pod was scheduled earlier than give that pod higer priority
	return int32(currentUnixTime - podInfo.Timestamp.Unix()), false
}

// Less is the function used by the activeQ heap algorithm to sort pods.
// It sorts pods based on their priorities. When the priorities are equal, it uses
// the Pod QoS classes to break the tie.
func (qs *QueueScheduler) Less(pInfo1, pInfo2 *framework.QueuedPodInfo) bool {
	currentUnixTime := time.Now().Unix()

	priority1, labelExists1 := GetPodPriority(pInfo1, currentUnixTime)
	priority2, labelExists2 := GetPodPriority(pInfo2, currentUnixTime)

	log(fmt.Sprintf("Priority1 = %d, Priority2 = %d\n", priority1, priority2), "info")

	// if label exists for either of the two pods, then just use that as priority
	// else use the Timestamp difference
	if labelExists1 && labelExists2 {
		return priority1 > priority2
	} else if labelExists1 && !labelExists2 {
		return true
	} else if !labelExists1 && labelExists2 {
		return false
	}

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
func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	log("Initialize new framework pluging &QueueScheduler{}", "info")

	return &QueueScheduler{}, nil
}
