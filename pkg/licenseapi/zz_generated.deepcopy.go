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
	if in.Buttons != nil {
		in, out := &in.Buttons, &out.Buttons
		*out = make([]*Button, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Button)
				**out = **in
			}
		}
	}
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
func (in *ButtonActionInput) DeepCopyInto(out *ButtonActionInput) {
	*out = *in
	if in.InstanceTokenAuth != nil {
		in, out := &in.InstanceTokenAuth, &out.InstanceTokenAuth
		*out = new(InstanceTokenAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ButtonActionInput.
func (in *ButtonActionInput) DeepCopy() *ButtonActionInput {
	if in == nil {
		return nil
	}
	out := new(ButtonActionInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ButtonActionOutput) DeepCopyInto(out *ButtonActionOutput) {
	*out = *in
	if in.Buttons != nil {
		in, out := &in.Buttons, &out.Buttons
		*out = make([]Button, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ButtonActionOutput.
func (in *ButtonActionOutput) DeepCopy() *ButtonActionOutput {
	if in == nil {
		return nil
	}
	out := new(ButtonActionOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChatAuthCreateInput) DeepCopyInto(out *ChatAuthCreateInput) {
	*out = *in
	if in.InstanceTokenAuth != nil {
		in, out := &in.InstanceTokenAuth, &out.InstanceTokenAuth
		*out = new(InstanceTokenAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChatAuthCreateInput.
func (in *ChatAuthCreateInput) DeepCopy() *ChatAuthCreateInput {
	if in == nil {
		return nil
	}
	out := new(ChatAuthCreateInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChatAuthCreateOutput) DeepCopyInto(out *ChatAuthCreateOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChatAuthCreateOutput.
func (in *ChatAuthCreateOutput) DeepCopy() *ChatAuthCreateOutput {
	if in == nil {
		return nil
	}
	out := new(ChatAuthCreateOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainToken) DeepCopyInto(out *DomainToken) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainToken.
func (in *DomainToken) DeepCopy() *DomainToken {
	if in == nil {
		return nil
	}
	out := new(DomainToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
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
		*out = make(map[string]ResourceQuantity, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
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
func (in *InstanceTokenClaims) DeepCopyInto(out *InstanceTokenClaims) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceTokenClaims.
func (in *InstanceTokenClaims) DeepCopy() *InstanceTokenClaims {
	if in == nil {
		return nil
	}
	out := new(InstanceTokenClaims)
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
		*out = make([]*Button, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Button)
				**out = **in
			}
		}
	}
	if in.Announcements != nil {
		in, out := &in.Announcements, &out.Announcements
		*out = make([]*Announcement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Announcement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]*Module, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Module)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.Routes = in.Routes
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
func (in *LicenseAPIRoute) DeepCopyInto(out *LicenseAPIRoute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAPIRoute.
func (in *LicenseAPIRoute) DeepCopy() *LicenseAPIRoute {
	if in == nil {
		return nil
	}
	out := new(LicenseAPIRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseAPIRoutes) DeepCopyInto(out *LicenseAPIRoutes) {
	*out = *in
	out.ChatAuth = in.ChatAuth
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseAPIRoutes.
func (in *LicenseAPIRoutes) DeepCopy() *LicenseAPIRoutes {
	if in == nil {
		return nil
	}
	out := new(LicenseAPIRoutes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]*ResourceQuantity, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceQuantity)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]*Feature, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Feature)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
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
func (in *ResourceCount) DeepCopyInto(out *ResourceCount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCount.
func (in *ResourceCount) DeepCopy() *ResourceCount {
	if in == nil {
		return nil
	}
	out := new(ResourceCount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuantity) DeepCopyInto(out *ResourceQuantity) {
	*out = *in
	if in.AdjustButton != nil {
		in, out := &in.AdjustButton, &out.AdjustButton
		*out = new(Button)
		**out = **in
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceCount)
		**out = **in
	}
	if in.Usage != nil {
		in, out := &in.Usage, &out.Usage
		*out = new(ResourceCount)
		**out = **in
	}
	if in.BlockRequests != nil {
		in, out := &in.BlockRequests, &out.BlockRequests
		*out = make([]Request, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
