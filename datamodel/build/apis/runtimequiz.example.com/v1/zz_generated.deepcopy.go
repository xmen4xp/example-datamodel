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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Child) DeepCopyInto(out *Child) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Child.
func (in *Child) DeepCopy() *Child {
	if in == nil {
		return nil
	}
	out := new(Child)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NexusStatus) DeepCopyInto(out *NexusStatus) {
	*out = *in
	out.SyncerStatus = in.SyncerStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NexusStatus.
func (in *NexusStatus) DeepCopy() *NexusStatus {
	if in == nil {
		return nil
	}
	out := new(NexusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeQuiz) DeepCopyInto(out *RuntimeQuiz) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeQuiz.
func (in *RuntimeQuiz) DeepCopy() *RuntimeQuiz {
	if in == nil {
		return nil
	}
	out := new(RuntimeQuiz)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeQuiz) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeQuizList) DeepCopyInto(out *RuntimeQuizList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuntimeQuiz, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeQuizList.
func (in *RuntimeQuizList) DeepCopy() *RuntimeQuizList {
	if in == nil {
		return nil
	}
	out := new(RuntimeQuizList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuntimeQuizList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeQuizNexusStatus) DeepCopyInto(out *RuntimeQuizNexusStatus) {
	*out = *in
	out.Status = in.Status
	out.Nexus = in.Nexus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeQuizNexusStatus.
func (in *RuntimeQuizNexusStatus) DeepCopy() *RuntimeQuizNexusStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeQuizNexusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeQuizSpec) DeepCopyInto(out *RuntimeQuizSpec) {
	*out = *in
	if in.AnswersGvk != nil {
		in, out := &in.AnswersGvk, &out.AnswersGvk
		*out = make(map[string]Child, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.QuizGvk != nil {
		in, out := &in.QuizGvk, &out.QuizGvk
		*out = new(Link)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeQuizSpec.
func (in *RuntimeQuizSpec) DeepCopy() *RuntimeQuizSpec {
	if in == nil {
		return nil
	}
	out := new(RuntimeQuizSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeQuizStatus) DeepCopyInto(out *RuntimeQuizStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeQuizStatus.
func (in *RuntimeQuizStatus) DeepCopy() *RuntimeQuizStatus {
	if in == nil {
		return nil
	}
	out := new(RuntimeQuizStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncerStatus) DeepCopyInto(out *SyncerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncerStatus.
func (in *SyncerStatus) DeepCopy() *SyncerStatus {
	if in == nil {
		return nil
	}
	out := new(SyncerStatus)
	in.DeepCopyInto(out)
	return out
}
