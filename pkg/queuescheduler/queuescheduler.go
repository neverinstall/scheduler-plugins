package queuescheduler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	v1qos "k8s.io/kubernetes/pkg/apis/core/v1/helper/qos"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// Implements framework.Plugin
type QueueScheduler struct{}

var _ framework.QueueSortPlugin = &QueueScheduler{}

const Name = "QueueScheduler"

func log(msg string) {
	fmt.Printf("\n\n-------------------------------------------------------\n")
	fmt.Printf("%s \n", msg)
	fmt.Printf("-------------------------------------------------------\n\n")
}

// To implement framework.Plugin
func (qs QueueScheduler) Name() string {
	// Name is the name of the plugin used in the Registry and configurations.
	return Name
}

// Check if the pod has a label called priority and if it does, then give this pod preference
func GetPodPriority(podInfo *framework.QueuedPodInfo, currentUnixTime int64) int32 {
	pod := podInfo.Pod

	if pod.ObjectMeta.Labels == nil {
		return int32(currentUnixTime - podInfo.Timestamp.Unix())
	}

	// fmt.Println("pod.ObjectMeta.Labels", pod.ObjectMeta.Labels)
	if labelValue, ok := pod.ObjectMeta.Labels["priority"]; ok {
		priority, err := strconv.Atoi(strings.Split(labelValue, "-")[1])

		if err == nil {
			return int32(priority)
		}
	}

	// if one pod was scheduled earlier than give that pod higer priority
	return int32(currentUnixTime - podInfo.Timestamp.Unix())
}

// Less is the function used by the activeQ heap algorithm to sort pods.
// It sorts pods based on their priorities. When the priorities are equal, it uses
// the Pod QoS classes to break the tie.
func (qs *QueueScheduler) Less(pInfo1, pInfo2 *framework.QueuedPodInfo) bool {
	// log("Made it into the Less function of the custon scheduler")

	currentUnixTime := time.Now().Unix()

	priority1 := GetPodPriority(pInfo1, currentUnixTime)
	priority2 := GetPodPriority(pInfo2, currentUnixTime)

	log(fmt.Sprintf("Priority1 = %d, Priority2 = %d\n", priority1, priority2))

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
	log("Initialize new framework pluging &QueueScheduler{}")

	return &QueueScheduler{}, nil
}
