//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 K8sGPT Contributors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AISpec) DeepCopyInto(out *AISpec) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AISpec.
func (in *AISpec) DeepCopy() *AISpec {
	if in == nil {
		return nil
	}
	out := new(AISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureBackend) DeepCopyInto(out *AzureBackend) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureBackend.
func (in *AzureBackend) DeepCopy() *AzureBackend {
	if in == nil {
		return nil
	}
	out := new(AzureBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backstage) DeepCopyInto(out *Backstage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backstage.
func (in *Backstage) DeepCopy() *Backstage {
	if in == nil {
		return nil
	}
	out := new(Backstage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsRef) DeepCopyInto(out *CredentialsRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsRef.
func (in *CredentialsRef) DeepCopy() *CredentialsRef {
	if in == nil {
		return nil
	}
	out := new(CredentialsRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraOptionsRef) DeepCopyInto(out *ExtraOptionsRef) {
	*out = *in
	if in.Backstage != nil {
		in, out := &in.Backstage, &out.Backstage
		*out = new(Backstage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraOptionsRef.
func (in *ExtraOptionsRef) DeepCopy() *ExtraOptionsRef {
	if in == nil {
		return nil
	}
	out := new(ExtraOptionsRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Failure) DeepCopyInto(out *Failure) {
	*out = *in
	if in.Sensitive != nil {
		in, out := &in.Sensitive, &out.Sensitive
		*out = make([]Sensitive, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Failure.
func (in *Failure) DeepCopy() *Failure {
	if in == nil {
		return nil
	}
	out := new(Failure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSBackend) DeepCopyInto(out *GCSBackend) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSBackend.
func (in *GCSBackend) DeepCopy() *GCSBackend {
	if in == nil {
		return nil
	}
	out := new(GCSBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullSecrets) DeepCopyInto(out *ImagePullSecrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullSecrets.
func (in *ImagePullSecrets) DeepCopy() *ImagePullSecrets {
	if in == nil {
		return nil
	}
	out := new(ImagePullSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Integrations) DeepCopyInto(out *Integrations) {
	*out = *in
	if in.Trivy != nil {
		in, out := &in.Trivy, &out.Trivy
		*out = new(Trivy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Integrations.
func (in *Integrations) DeepCopy() *Integrations {
	if in == nil {
		return nil
	}
	out := new(Integrations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sGPT) DeepCopyInto(out *K8sGPT) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sGPT.
func (in *K8sGPT) DeepCopy() *K8sGPT {
	if in == nil {
		return nil
	}
	out := new(K8sGPT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8sGPT) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sGPTList) DeepCopyInto(out *K8sGPTList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]K8sGPT, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sGPTList.
func (in *K8sGPTList) DeepCopy() *K8sGPTList {
	if in == nil {
		return nil
	}
	out := new(K8sGPTList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *K8sGPTList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sGPTSpec) DeepCopyInto(out *K8sGPTSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]ImagePullSecrets, len(*in))
		copy(*out, *in)
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraOptions != nil {
		in, out := &in.ExtraOptions, &out.ExtraOptions
		*out = new(ExtraOptionsRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(WebhookRef)
		**out = **in
	}
	if in.AI != nil {
		in, out := &in.AI, &out.AI
		*out = new(AISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RemoteCache != nil {
		in, out := &in.RemoteCache, &out.RemoteCache
		*out = new(RemoteCacheRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Integrations != nil {
		in, out := &in.Integrations, &out.Integrations
		*out = new(Integrations)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sGPTSpec.
func (in *K8sGPTSpec) DeepCopy() *K8sGPTSpec {
	if in == nil {
		return nil
	}
	out := new(K8sGPTSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sGPTStatus) DeepCopyInto(out *K8sGPTStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sGPTStatus.
func (in *K8sGPTStatus) DeepCopy() *K8sGPTStatus {
	if in == nil {
		return nil
	}
	out := new(K8sGPTStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteCacheRef) DeepCopyInto(out *RemoteCacheRef) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(CredentialsRef)
		**out = **in
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(GCSBackend)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Backend)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureBackend)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteCacheRef.
func (in *RemoteCacheRef) DeepCopy() *RemoteCacheRef {
	if in == nil {
		return nil
	}
	out := new(RemoteCacheRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Result) DeepCopyInto(out *Result) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Result.
func (in *Result) DeepCopy() *Result {
	if in == nil {
		return nil
	}
	out := new(Result)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Result) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultList) DeepCopyInto(out *ResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Result, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultList.
func (in *ResultList) DeepCopy() *ResultList {
	if in == nil {
		return nil
	}
	out := new(ResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultSpec) DeepCopyInto(out *ResultSpec) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = make([]Failure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultSpec.
func (in *ResultSpec) DeepCopy() *ResultSpec {
	if in == nil {
		return nil
	}
	out := new(ResultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultStatus) DeepCopyInto(out *ResultStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultStatus.
func (in *ResultStatus) DeepCopy() *ResultStatus {
	if in == nil {
		return nil
	}
	out := new(ResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Backend) DeepCopyInto(out *S3Backend) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Backend.
func (in *S3Backend) DeepCopy() *S3Backend {
	if in == nil {
		return nil
	}
	out := new(S3Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sensitive) DeepCopyInto(out *Sensitive) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sensitive.
func (in *Sensitive) DeepCopy() *Sensitive {
	if in == nil {
		return nil
	}
	out := new(Sensitive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trivy) DeepCopyInto(out *Trivy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trivy.
func (in *Trivy) DeepCopy() *Trivy {
	if in == nil {
		return nil
	}
	out := new(Trivy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookRef) DeepCopyInto(out *WebhookRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookRef.
func (in *WebhookRef) DeepCopy() *WebhookRef {
	if in == nil {
		return nil
	}
	out := new(WebhookRef)
	in.DeepCopyInto(out)
	return out
}
