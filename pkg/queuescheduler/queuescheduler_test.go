package queuescheduler

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const green = "\u001b[32m"
const red = "\u001b[31m"
const reset = "\u001b[0m"

func TestCustomLessFunction(t *testing.T) {
	tests := []struct {
		name   string
		pInfo1 *framework.QueuedPodInfo
		pInfo2 *framework.QueuedPodInfo
		want   bool
	}{
		{
			name: "P1 > P2",
			pInfo1: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(createPod("P1", "subscribedUser-100", nil, nil)),
			},
			pInfo2: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(createPod("P2", "", nil, nil)),
			},
			want: true,
		},
		{
			name: "P1 < P2",
			pInfo1: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(createPod("P1", "", nil, nil)),
			},
			pInfo2: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(createPod("P2", "subscribedUser-100", nil, nil)),
			},
			want: false,
		},
		{
			name: "P1 > P2",
			pInfo1: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(createPod("P1", "subscribedUser-100", nil, nil)),
			},
			pInfo2: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(createPod("P2", "newUser-90", nil, nil)),
			},
			want: true,
		},
		{
			name: "P1 and P2 are Guaranteed",
			pInfo1: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(
					createPod("p1", "", getResourceLimits("100m", "100Mi"), getResourceLimits("100m", "100Mi")),
				),
			},
			pInfo2: &framework.QueuedPodInfo{
				PodInfo: framework.NewPodInfo(
					createPod("p2", "", getResourceLimits("100m", "100Mi"), getResourceLimits("100m", "100Mi")),
				),
			},
			want: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			plugin := &QueueScheduler{}

			returnedValue := plugin.Less(testCase.pInfo1, testCase.pInfo2)

			var success string

			if returnedValue == testCase.want {
				success = fmt.Sprintf("%sPassed %s", green, reset)
			} else {
				success = fmt.Sprintf("%sFailed %s", red, reset)
			}

			fmt.Printf("%s Less() = %v, want %v\n\n", success, returnedValue, testCase.want)
		})
	}
}

func createPod(name string, priorityLabel string, requests, limits v1.ResourceList) *v1.Pod {
	var labelsDict map[string]string

	if len(priorityLabel) > 0 {
		labelsDict = map[string]string{
			"priority": priorityLabel,
		}
	}

	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labelsDict,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: name,
					Resources: v1.ResourceRequirements{
						Requests: requests,
						Limits:   limits,
					},
				},
			},
		},
	}
}

func getResourceLimits(cpu, memory string) v1.ResourceList {
	res := v1.ResourceList{}
	if cpu != "" {
		res[v1.ResourceCPU] = resource.MustParse(cpu)
	}
	if memory != "" {
		res[v1.ResourceMemory] = resource.MustParse(memory)
	}
	return res
}
