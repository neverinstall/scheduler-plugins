package queuescheduler

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

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
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			plugin := &QueueScheduler{}

			got := plugin.Less(testCase.pInfo1, testCase.pInfo2)

			var success string

			if got == testCase.want {
				success = "Passed"
			} else {
				success = "Failed"
			}

			fmt.Printf("%s Less() = %v, want %v\n", success, got, testCase.want)
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
