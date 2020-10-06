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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerCAInjectorConfig) DeepCopyInto(out *CertManagerCAInjectorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Flags = in.Flags
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerCAInjectorConfig.
func (in *CertManagerCAInjectorConfig) DeepCopy() *CertManagerCAInjectorConfig {
	if in == nil {
		return nil
	}
	out := new(CertManagerCAInjectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerCAInjectorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerCAInjectorFlags) DeepCopyInto(out *CertManagerCAInjectorFlags) {
	*out = *in
	out.loggingFlags = in.loggingFlags
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerCAInjectorFlags.
func (in *CertManagerCAInjectorFlags) DeepCopy() *CertManagerCAInjectorFlags {
	if in == nil {
		return nil
	}
	out := new(CertManagerCAInjectorFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerControllerConfig) DeepCopyInto(out *CertManagerControllerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Flags.DeepCopyInto(&out.Flags)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerControllerConfig.
func (in *CertManagerControllerConfig) DeepCopy() *CertManagerControllerConfig {
	if in == nil {
		return nil
	}
	out := new(CertManagerControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerControllerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerControllerFlags) DeepCopyInto(out *CertManagerControllerFlags) {
	*out = *in
	out.loggingFlags = in.loggingFlags
	if in.AutoCertificateAnnotations != nil {
		in, out := &in.AutoCertificateAnnotations, &out.AutoCertificateAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNS01RecursiveNameservers != nil {
		in, out := &in.DNS01RecursiveNameservers, &out.DNS01RecursiveNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerControllerFlags.
func (in *CertManagerControllerFlags) DeepCopy() *CertManagerControllerFlags {
	if in == nil {
		return nil
	}
	out := new(CertManagerControllerFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerWebhookConfig) DeepCopyInto(out *CertManagerWebhookConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.Flags.DeepCopyInto(&out.Flags)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerWebhookConfig.
func (in *CertManagerWebhookConfig) DeepCopy() *CertManagerWebhookConfig {
	if in == nil {
		return nil
	}
	out := new(CertManagerWebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerWebhookConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerWebhookFlags) DeepCopyInto(out *CertManagerWebhookFlags) {
	*out = *in
	out.loggingFlags = in.loggingFlags
	if in.DynamicServingDNSNames != nil {
		in, out := &in.DynamicServingDNSNames, &out.DynamicServingDNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSCipherSuites != nil {
		in, out := &in.TLSCipherSuites, &out.TLSCipherSuites
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerWebhookFlags.
func (in *CertManagerWebhookFlags) DeepCopy() *CertManagerWebhookFlags {
	if in == nil {
		return nil
	}
	out := new(CertManagerWebhookFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceLocation) DeepCopyInto(out *TraceLocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceLocation.
func (in *TraceLocation) DeepCopy() *TraceLocation {
	if in == nil {
		return nil
	}
	out := new(TraceLocation)
	in.DeepCopyInto(out)
	return out
}
