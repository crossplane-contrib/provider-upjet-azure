//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2Environment) DeepCopyInto(out *Gen2Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2Environment.
func (in *Gen2Environment) DeepCopy() *Gen2Environment {
	if in == nil {
		return nil
	}
	out := new(Gen2Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gen2Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2EnvironmentInitParameters) DeepCopyInto(out *Gen2EnvironmentInitParameters) {
	*out = *in
	if in.IDProperties != nil {
		in, out := &in.IDProperties, &out.IDProperties
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WarmStoreDataRetentionTime != nil {
		in, out := &in.WarmStoreDataRetentionTime, &out.WarmStoreDataRetentionTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2EnvironmentInitParameters.
func (in *Gen2EnvironmentInitParameters) DeepCopy() *Gen2EnvironmentInitParameters {
	if in == nil {
		return nil
	}
	out := new(Gen2EnvironmentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2EnvironmentList) DeepCopyInto(out *Gen2EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gen2Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2EnvironmentList.
func (in *Gen2EnvironmentList) DeepCopy() *Gen2EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(Gen2EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gen2EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2EnvironmentObservation) DeepCopyInto(out *Gen2EnvironmentObservation) {
	*out = *in
	if in.DataAccessFqdn != nil {
		in, out := &in.DataAccessFqdn, &out.DataAccessFqdn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IDProperties != nil {
		in, out := &in.IDProperties, &out.IDProperties
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WarmStoreDataRetentionTime != nil {
		in, out := &in.WarmStoreDataRetentionTime, &out.WarmStoreDataRetentionTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2EnvironmentObservation.
func (in *Gen2EnvironmentObservation) DeepCopy() *Gen2EnvironmentObservation {
	if in == nil {
		return nil
	}
	out := new(Gen2EnvironmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2EnvironmentParameters) DeepCopyInto(out *Gen2EnvironmentParameters) {
	*out = *in
	if in.IDProperties != nil {
		in, out := &in.IDProperties, &out.IDProperties
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.WarmStoreDataRetentionTime != nil {
		in, out := &in.WarmStoreDataRetentionTime, &out.WarmStoreDataRetentionTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2EnvironmentParameters.
func (in *Gen2EnvironmentParameters) DeepCopy() *Gen2EnvironmentParameters {
	if in == nil {
		return nil
	}
	out := new(Gen2EnvironmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2EnvironmentSpec) DeepCopyInto(out *Gen2EnvironmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2EnvironmentSpec.
func (in *Gen2EnvironmentSpec) DeepCopy() *Gen2EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(Gen2EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gen2EnvironmentStatus) DeepCopyInto(out *Gen2EnvironmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gen2EnvironmentStatus.
func (in *Gen2EnvironmentStatus) DeepCopy() *Gen2EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(Gen2EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageInitParameters) DeepCopyInto(out *StorageInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageInitParameters.
func (in *StorageInitParameters) DeepCopy() *StorageInitParameters {
	if in == nil {
		return nil
	}
	out := new(StorageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageObservation) DeepCopyInto(out *StorageObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageObservation.
func (in *StorageObservation) DeepCopy() *StorageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageParameters) DeepCopyInto(out *StorageParameters) {
	*out = *in
	out.KeySecretRef = in.KeySecretRef
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageParameters.
func (in *StorageParameters) DeepCopy() *StorageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageParameters)
	in.DeepCopyInto(out)
	return out
}