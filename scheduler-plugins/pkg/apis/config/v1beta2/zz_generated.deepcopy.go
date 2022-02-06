//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/kube-scheduler/config/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoschedulingArgs) DeepCopyInto(out *CoschedulingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PermitWaitingTimeSeconds != nil {
		in, out := &in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds
		*out = new(int64)
		**out = **in
	}
	if in.DeniedPGExpirationTimeSeconds != nil {
		in, out := &in.DeniedPGExpirationTimeSeconds, &out.DeniedPGExpirationTimeSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoschedulingArgs.
func (in *CoschedulingArgs) DeepCopy() *CoschedulingArgs {
	if in == nil {
		return nil
	}
	out := new(CoschedulingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoschedulingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadVariationRiskBalancingArgs) DeepCopyInto(out *LoadVariationRiskBalancingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.MetricProvider.DeepCopyInto(&out.MetricProvider)
	if in.WatcherAddress != nil {
		in, out := &in.WatcherAddress, &out.WatcherAddress
		*out = new(string)
		**out = **in
	}
	if in.SafeVarianceMargin != nil {
		in, out := &in.SafeVarianceMargin, &out.SafeVarianceMargin
		*out = new(float64)
		**out = **in
	}
	if in.SafeVarianceSensitivity != nil {
		in, out := &in.SafeVarianceSensitivity, &out.SafeVarianceSensitivity
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadVariationRiskBalancingArgs.
func (in *LoadVariationRiskBalancingArgs) DeepCopy() *LoadVariationRiskBalancingArgs {
	if in == nil {
		return nil
	}
	out := new(LoadVariationRiskBalancingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadVariationRiskBalancingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricProviderSpec) DeepCopyInto(out *MetricProviderSpec) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricProviderSpec.
func (in *MetricProviderSpec) DeepCopy() *MetricProviderSpec {
	if in == nil {
		return nil
	}
	out := new(MetricProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceTopologyMatchArgs) DeepCopyInto(out *NodeResourceTopologyMatchArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.ScoringStrategy != nil {
		in, out := &in.ScoringStrategy, &out.ScoringStrategy
		*out = new(ScoringStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceTopologyMatchArgs.
func (in *NodeResourceTopologyMatchArgs) DeepCopy() *NodeResourceTopologyMatchArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourceTopologyMatchArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourceTopologyMatchArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesAllocatableArgs) DeepCopyInto(out *NodeResourcesAllocatableArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesAllocatableArgs.
func (in *NodeResourcesAllocatableArgs) DeepCopy() *NodeResourcesAllocatableArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesAllocatableArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesAllocatableArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreemptionTolerationArgs) DeepCopyInto(out *PreemptionTolerationArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.MinCandidateNodesPercentage != nil {
		in, out := &in.MinCandidateNodesPercentage, &out.MinCandidateNodesPercentage
		*out = new(int32)
		**out = **in
	}
	if in.MinCandidateNodesAbsolute != nil {
		in, out := &in.MinCandidateNodesAbsolute, &out.MinCandidateNodesAbsolute
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreemptionTolerationArgs.
func (in *PreemptionTolerationArgs) DeepCopy() *PreemptionTolerationArgs {
	if in == nil {
		return nil
	}
	out := new(PreemptionTolerationArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreemptionTolerationArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSchedulerArgs) DeepCopyInto(out *QueueSchedulerArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSchedulerArgs.
func (in *QueueSchedulerArgs) DeepCopy() *QueueSchedulerArgs {
	if in == nil {
		return nil
	}
	out := new(QueueSchedulerArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueSchedulerArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoringStrategy) DeepCopyInto(out *ScoringStrategy) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1.ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoringStrategy.
func (in *ScoringStrategy) DeepCopy() *ScoringStrategy {
	if in == nil {
		return nil
	}
	out := new(ScoringStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetLoadPackingArgs) DeepCopyInto(out *TargetLoadPackingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DefaultRequests != nil {
		in, out := &in.DefaultRequests, &out.DefaultRequests
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.DefaultRequestsMultiplier != nil {
		in, out := &in.DefaultRequestsMultiplier, &out.DefaultRequestsMultiplier
		*out = new(string)
		**out = **in
	}
	if in.TargetUtilization != nil {
		in, out := &in.TargetUtilization, &out.TargetUtilization
		*out = new(int64)
		**out = **in
	}
	in.MetricProvider.DeepCopyInto(&out.MetricProvider)
	if in.WatcherAddress != nil {
		in, out := &in.WatcherAddress, &out.WatcherAddress
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetLoadPackingArgs.
func (in *TargetLoadPackingArgs) DeepCopy() *TargetLoadPackingArgs {
	if in == nil {
		return nil
	}
	out := new(TargetLoadPackingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetLoadPackingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
