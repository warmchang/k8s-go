// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckPolicy) DeepCopyInto(out *HealthCheckPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckPolicy.
func (in *HealthCheckPolicy) DeepCopy() *HealthCheckPolicy {
	if in == nil {
		return nil
	}
	out := new(HealthCheckPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthCheckPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckPolicyList) DeepCopyInto(out *HealthCheckPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HealthCheckPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckPolicyList.
func (in *HealthCheckPolicyList) DeepCopy() *HealthCheckPolicyList {
	if in == nil {
		return nil
	}
	out := new(HealthCheckPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HealthCheckPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckPolicySpec) DeepCopyInto(out *HealthCheckPolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckPolicySpec.
func (in *HealthCheckPolicySpec) DeepCopy() *HealthCheckPolicySpec {
	if in == nil {
		return nil
	}
	out := new(HealthCheckPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckPolicyStatus) DeepCopyInto(out *HealthCheckPolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckPolicyStatus.
func (in *HealthCheckPolicyStatus) DeepCopy() *HealthCheckPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(HealthCheckPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
