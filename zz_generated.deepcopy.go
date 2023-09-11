//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package licenseapi

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analytics) DeepCopyInto(out *Analytics) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make([]Request, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analytics.
func (in *Analytics) DeepCopy() *Analytics {
	if in == nil {
		return nil
	}
	out := new(Analytics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Announcement) DeepCopyInto(out *Announcement) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Announcement.
func (in *Announcement) DeepCopy() *Announcement {
	if in == nil {
		return nil
	}
	out := new(Announcement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Button) DeepCopyInto(out *Button) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Button.
func (in *Button) DeepCopy() *Button {
	if in == nil {
		return nil
	}
	out := new(Button)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimSubdomainInput) DeepCopyInto(out *ClaimSubdomainInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimSubdomainInput.
func (in *ClaimSubdomainInput) DeepCopy() *ClaimSubdomainInput {
	if in == nil {
		return nil
	}
	out := new(ClaimSubdomainInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimSubdomainOutput) DeepCopyInto(out *ClaimSubdomainOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimSubdomainOutput.
func (in *ClaimSubdomainOutput) DeepCopy() *ClaimSubdomainOutput {
	if in == nil {
		return nil
	}
	out := new(ClaimSubdomainOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Feature.
func (in *Feature) DeepCopy() *Feature {
	if in == nil {
		return nil
	}
	out := new(Feature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceCreateInput) DeepCopyInto(out *InstanceCreateInput) {
	*out = *in
	if in.InstanceTokenAuth != nil {
		in, out := &in.InstanceTokenAuth, &out.InstanceTokenAuth
		*out = new(InstanceTokenAuth)
		**out = **in
	}
	if in.AllocatedResources != nil {
		in, out := &in.AllocatedResources, &out.AllocatedResources
		*out = new(map[string]ResourceQuantity)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]ResourceQuantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceCreateInput.
func (in *InstanceCreateInput) DeepCopy() *InstanceCreateInput {
	if in == nil {
		return nil
	}
	out := new(InstanceCreateInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceCreateOutput) DeepCopyInto(out *InstanceCreateOutput) {
	*out = *in
	if in.License != nil {
		in, out := &in.License, &out.License
		*out = new(License)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceCreateOutput.
func (in *InstanceCreateOutput) DeepCopy() *InstanceCreateOutput {
	if in == nil {
		return nil
	}
	out := new(InstanceCreateOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceTokenAuth) DeepCopyInto(out *InstanceTokenAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceTokenAuth.
func (in *InstanceTokenAuth) DeepCopy() *InstanceTokenAuth {
	if in == nil {
		return nil
	}
	out := new(InstanceTokenAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntercomHashCreateInput) DeepCopyInto(out *IntercomHashCreateInput) {
	*out = *in
	if in.InstanceTokenAuth != nil {
		in, out := &in.InstanceTokenAuth, &out.InstanceTokenAuth
		*out = new(InstanceTokenAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntercomHashCreateInput.
func (in *IntercomHashCreateInput) DeepCopy() *IntercomHashCreateInput {
	if in == nil {
		return nil
	}
	out := new(IntercomHashCreateInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntercomHashCreateOutput) DeepCopyInto(out *IntercomHashCreateOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntercomHashCreateOutput.
func (in *IntercomHashCreateOutput) DeepCopy() *IntercomHashCreateOutput {
	if in == nil {
		return nil
	}
	out := new(IntercomHashCreateOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LegacyLicense) DeepCopyInto(out *LegacyLicense) {
	*out = *in
	if in.Announcements != nil {
		in, out := &in.Announcements, &out.Announcements
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Analytics != nil {
		in, out := &in.Analytics, &out.Analytics
		*out = new(Analytics)
		(*in).DeepCopyInto(*out)
	}
	if in.Buttons != nil {
		in, out := &in.Buttons, &out.Buttons
		*out = make([]Button, len(*in))
		copy(*out, *in)
	}
	if in.BlockRequests != nil {
		in, out := &in.BlockRequests, &out.BlockRequests
		*out = make([]Request, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]ResourceQuantity, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LegacyLicense.
func (in *LegacyLicense) DeepCopy() *LegacyLicense {
	if in == nil {
		return nil
	}
	out := new(LegacyLicense)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *License) DeepCopyInto(out *License) {
	*out = *in
	if in.Analytics != nil {
		in, out := &in.Analytics, &out.Analytics
		*out = new(Analytics)
		(*in).DeepCopyInto(*out)
	}
	if in.Buttons != nil {
		in, out := &in.Buttons, &out.Buttons
		*out = make([]Button, len(*in))
		copy(*out, *in)
	}
	if in.Announcements != nil {
		in, out := &in.Announcements, &out.Announcements
		*out = make([]Announcement, len(*in))
		copy(*out, *in)
	}
	if in.BlockRequests != nil {
		in, out := &in.BlockRequests, &out.BlockRequests
		*out = make([]Request, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]ResourceQuantity, len(*in))
		copy(*out, *in)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]Feature, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new License.
func (in *License) DeepCopy() *License {
	if in == nil {
		return nil
	}
	out := new(License)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Request) DeepCopyInto(out *Request) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Request.
func (in *Request) DeepCopy() *Request {
	if in == nil {
		return nil
	}
	out := new(Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuantity) DeepCopyInto(out *ResourceQuantity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuantity.
func (in *ResourceQuantity) DeepCopy() *ResourceQuantity {
	if in == nil {
		return nil
	}
	out := new(ResourceQuantity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardRequestInput) DeepCopyInto(out *StandardRequestInput) {
	*out = *in
	if in.InstanceTokenAuth != nil {
		in, out := &in.InstanceTokenAuth, &out.InstanceTokenAuth
		*out = new(InstanceTokenAuth)
		**out = **in
	}
	out.StandardRequestInputFrontEnd = in.StandardRequestInputFrontEnd
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardRequestInput.
func (in *StandardRequestInput) DeepCopy() *StandardRequestInput {
	if in == nil {
		return nil
	}
	out := new(StandardRequestInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardRequestInputFrontEnd) DeepCopyInto(out *StandardRequestInputFrontEnd) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardRequestInputFrontEnd.
func (in *StandardRequestInputFrontEnd) DeepCopy() *StandardRequestInputFrontEnd {
	if in == nil {
		return nil
	}
	out := new(StandardRequestInputFrontEnd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardRequestOutput) DeepCopyInto(out *StandardRequestOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardRequestOutput.
func (in *StandardRequestOutput) DeepCopy() *StandardRequestOutput {
	if in == nil {
		return nil
	}
	out := new(StandardRequestOutput)
	in.DeepCopyInto(out)
	return out
}
