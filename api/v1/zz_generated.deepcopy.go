// +build !ignore_autogenerated

/*

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFed) DeepCopyInto(out *MeshFed) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFed.
func (in *MeshFed) DeepCopy() *MeshFed {
	if in == nil {
		return nil
	}
	out := new(MeshFed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshFed) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFedList) DeepCopyInto(out *MeshFedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MeshFed, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFedList.
func (in *MeshFedList) DeepCopy() *MeshFedList {
	if in == nil {
		return nil
	}
	out := new(MeshFedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshFedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFedSpec) DeepCopyInto(out *MeshFedSpec) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFedSpec.
func (in *MeshFedSpec) DeepCopy() *MeshFedSpec {
	if in == nil {
		return nil
	}
	out := new(MeshFedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshFedStatus) DeepCopyInto(out *MeshFedStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshFedStatus.
func (in *MeshFedStatus) DeepCopy() *MeshFedStatus {
	if in == nil {
		return nil
	}
	out := new(MeshFedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBinding) DeepCopyInto(out *ServiceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBinding.
func (in *ServiceBinding) DeepCopy() *ServiceBinding {
	if in == nil {
		return nil
	}
	out := new(ServiceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingList) DeepCopyInto(out *ServiceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingList.
func (in *ServiceBindingList) DeepCopy() *ServiceBindingList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingSpec) DeepCopyInto(out *ServiceBindingSpec) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingSpec.
func (in *ServiceBindingSpec) DeepCopy() *ServiceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingStatus) DeepCopyInto(out *ServiceBindingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingStatus.
func (in *ServiceBindingStatus) DeepCopy() *ServiceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExposition) DeepCopyInto(out *ServiceExposition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExposition.
func (in *ServiceExposition) DeepCopy() *ServiceExposition {
	if in == nil {
		return nil
	}
	out := new(ServiceExposition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExposition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExpositionList) DeepCopyInto(out *ServiceExpositionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceExposition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExpositionList.
func (in *ServiceExpositionList) DeepCopy() *ServiceExpositionList {
	if in == nil {
		return nil
	}
	out := new(ServiceExpositionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExpositionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExpositionSpec) DeepCopyInto(out *ServiceExpositionSpec) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExpositionSpec.
func (in *ServiceExpositionSpec) DeepCopy() *ServiceExpositionSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceExpositionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExpositionStatus) DeepCopyInto(out *ServiceExpositionStatus) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExpositionStatus.
func (in *ServiceExpositionStatus) DeepCopy() *ServiceExpositionStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceExpositionStatus)
	in.DeepCopyInto(out)
	return out
}
