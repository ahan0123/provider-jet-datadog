//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
func (in *JSON) DeepCopyInto(out *JSON) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSON.
func (in *JSON) DeepCopy() *JSON {
	if in == nil {
		return nil
	}
	out := new(JSON)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JSON) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONList) DeepCopyInto(out *JSONList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JSON, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONList.
func (in *JSONList) DeepCopy() *JSONList {
	if in == nil {
		return nil
	}
	out := new(JSONList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JSONList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONObservation) DeepCopyInto(out *JSONObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONObservation.
func (in *JSONObservation) DeepCopy() *JSONObservation {
	if in == nil {
		return nil
	}
	out := new(JSONObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONParameters) DeepCopyInto(out *JSONParameters) {
	*out = *in
	if in.Monitor != nil {
		in, out := &in.Monitor, &out.Monitor
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONParameters.
func (in *JSONParameters) DeepCopy() *JSONParameters {
	if in == nil {
		return nil
	}
	out := new(JSONParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSpec) DeepCopyInto(out *JSONSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSpec.
func (in *JSONSpec) DeepCopy() *JSONSpec {
	if in == nil {
		return nil
	}
	out := new(JSONSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONStatus) DeepCopyInto(out *JSONStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONStatus.
func (in *JSONStatus) DeepCopy() *JSONStatus {
	if in == nil {
		return nil
	}
	out := new(JSONStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitor) DeepCopyInto(out *Monitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitor.
func (in *Monitor) DeepCopy() *Monitor {
	if in == nil {
		return nil
	}
	out := new(Monitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Monitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorList) DeepCopyInto(out *MonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Monitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorList.
func (in *MonitorList) DeepCopy() *MonitorList {
	if in == nil {
		return nil
	}
	out := new(MonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorObservation) DeepCopyInto(out *MonitorObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorObservation.
func (in *MonitorObservation) DeepCopy() *MonitorObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorParameters) DeepCopyInto(out *MonitorParameters) {
	*out = *in
	if in.EnableLogsSample != nil {
		in, out := &in.EnableLogsSample, &out.EnableLogsSample
		*out = new(bool)
		**out = **in
	}
	if in.EscalationMessage != nil {
		in, out := &in.EscalationMessage, &out.EscalationMessage
		*out = new(string)
		**out = **in
	}
	if in.EvaluationDelay != nil {
		in, out := &in.EvaluationDelay, &out.EvaluationDelay
		*out = new(float64)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.GroupbySimpleMonitor != nil {
		in, out := &in.GroupbySimpleMonitor, &out.GroupbySimpleMonitor
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = new(bool)
		**out = **in
	}
	if in.Locked != nil {
		in, out := &in.Locked, &out.Locked
		*out = new(bool)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.MonitorThresholdWindows != nil {
		in, out := &in.MonitorThresholdWindows, &out.MonitorThresholdWindows
		*out = make([]MonitorThresholdWindowsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MonitorThresholds != nil {
		in, out := &in.MonitorThresholds, &out.MonitorThresholds
		*out = make([]MonitorThresholdsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NewGroupDelay != nil {
		in, out := &in.NewGroupDelay, &out.NewGroupDelay
		*out = new(float64)
		**out = **in
	}
	if in.NewHostDelay != nil {
		in, out := &in.NewHostDelay, &out.NewHostDelay
		*out = new(float64)
		**out = **in
	}
	if in.NoDataTimeframe != nil {
		in, out := &in.NoDataTimeframe, &out.NoDataTimeframe
		*out = new(float64)
		**out = **in
	}
	if in.NotifyAudit != nil {
		in, out := &in.NotifyAudit, &out.NotifyAudit
		*out = new(bool)
		**out = **in
	}
	if in.NotifyNoData != nil {
		in, out := &in.NotifyNoData, &out.NotifyNoData
		*out = new(bool)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(float64)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.RenotifyInterval != nil {
		in, out := &in.RenotifyInterval, &out.RenotifyInterval
		*out = new(float64)
		**out = **in
	}
	if in.RenotifyOccurrences != nil {
		in, out := &in.RenotifyOccurrences, &out.RenotifyOccurrences
		*out = new(float64)
		**out = **in
	}
	if in.RenotifyStatuses != nil {
		in, out := &in.RenotifyStatuses, &out.RenotifyStatuses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RequireFullWindow != nil {
		in, out := &in.RequireFullWindow, &out.RequireFullWindow
		*out = new(bool)
		**out = **in
	}
	if in.RestrictedRoles != nil {
		in, out := &in.RestrictedRoles, &out.RestrictedRoles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TimeoutH != nil {
		in, out := &in.TimeoutH, &out.TimeoutH
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Validate != nil {
		in, out := &in.Validate, &out.Validate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorParameters.
func (in *MonitorParameters) DeepCopy() *MonitorParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorSpec) DeepCopyInto(out *MonitorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorSpec.
func (in *MonitorSpec) DeepCopy() *MonitorSpec {
	if in == nil {
		return nil
	}
	out := new(MonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorStatus) DeepCopyInto(out *MonitorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorStatus.
func (in *MonitorStatus) DeepCopy() *MonitorStatus {
	if in == nil {
		return nil
	}
	out := new(MonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorThresholdWindowsObservation) DeepCopyInto(out *MonitorThresholdWindowsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorThresholdWindowsObservation.
func (in *MonitorThresholdWindowsObservation) DeepCopy() *MonitorThresholdWindowsObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorThresholdWindowsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorThresholdWindowsParameters) DeepCopyInto(out *MonitorThresholdWindowsParameters) {
	*out = *in
	if in.RecoveryWindow != nil {
		in, out := &in.RecoveryWindow, &out.RecoveryWindow
		*out = new(string)
		**out = **in
	}
	if in.TriggerWindow != nil {
		in, out := &in.TriggerWindow, &out.TriggerWindow
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorThresholdWindowsParameters.
func (in *MonitorThresholdWindowsParameters) DeepCopy() *MonitorThresholdWindowsParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorThresholdWindowsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorThresholdsObservation) DeepCopyInto(out *MonitorThresholdsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorThresholdsObservation.
func (in *MonitorThresholdsObservation) DeepCopy() *MonitorThresholdsObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorThresholdsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorThresholdsParameters) DeepCopyInto(out *MonitorThresholdsParameters) {
	*out = *in
	if in.Critical != nil {
		in, out := &in.Critical, &out.Critical
		*out = new(string)
		**out = **in
	}
	if in.CriticalRecovery != nil {
		in, out := &in.CriticalRecovery, &out.CriticalRecovery
		*out = new(string)
		**out = **in
	}
	if in.Ok != nil {
		in, out := &in.Ok, &out.Ok
		*out = new(string)
		**out = **in
	}
	if in.Unknown != nil {
		in, out := &in.Unknown, &out.Unknown
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(string)
		**out = **in
	}
	if in.WarningRecovery != nil {
		in, out := &in.WarningRecovery, &out.WarningRecovery
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorThresholdsParameters.
func (in *MonitorThresholdsParameters) DeepCopy() *MonitorThresholdsParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorThresholdsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryObservation) DeepCopyInto(out *QueryObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryObservation.
func (in *QueryObservation) DeepCopy() *QueryObservation {
	if in == nil {
		return nil
	}
	out := new(QueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryParameters) DeepCopyInto(out *QueryParameters) {
	*out = *in
	if in.Denominator != nil {
		in, out := &in.Denominator, &out.Denominator
		*out = new(string)
		**out = **in
	}
	if in.Numerator != nil {
		in, out := &in.Numerator, &out.Numerator
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryParameters.
func (in *QueryParameters) DeepCopy() *QueryParameters {
	if in == nil {
		return nil
	}
	out := new(QueryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjective) DeepCopyInto(out *ServiceLevelObjective) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjective.
func (in *ServiceLevelObjective) DeepCopy() *ServiceLevelObjective {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLevelObjective) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveList) DeepCopyInto(out *ServiceLevelObjectiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceLevelObjective, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveList.
func (in *ServiceLevelObjectiveList) DeepCopy() *ServiceLevelObjectiveList {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceLevelObjectiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveObservation) DeepCopyInto(out *ServiceLevelObjectiveObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Thresholds != nil {
		in, out := &in.Thresholds, &out.Thresholds
		*out = make([]ThresholdsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveObservation.
func (in *ServiceLevelObjectiveObservation) DeepCopy() *ServiceLevelObjectiveObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveParameters) DeepCopyInto(out *ServiceLevelObjectiveParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceDelete != nil {
		in, out := &in.ForceDelete, &out.ForceDelete
		*out = new(bool)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MonitorIds != nil {
		in, out := &in.MonitorIds, &out.MonitorIds
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = make([]QueryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Thresholds != nil {
		in, out := &in.Thresholds, &out.Thresholds
		*out = make([]ThresholdsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Validate != nil {
		in, out := &in.Validate, &out.Validate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveParameters.
func (in *ServiceLevelObjectiveParameters) DeepCopy() *ServiceLevelObjectiveParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveSpec) DeepCopyInto(out *ServiceLevelObjectiveSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveSpec.
func (in *ServiceLevelObjectiveSpec) DeepCopy() *ServiceLevelObjectiveSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceLevelObjectiveStatus) DeepCopyInto(out *ServiceLevelObjectiveStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceLevelObjectiveStatus.
func (in *ServiceLevelObjectiveStatus) DeepCopy() *ServiceLevelObjectiveStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceLevelObjectiveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThresholdsObservation) DeepCopyInto(out *ThresholdsObservation) {
	*out = *in
	if in.TargetDisplay != nil {
		in, out := &in.TargetDisplay, &out.TargetDisplay
		*out = new(string)
		**out = **in
	}
	if in.WarningDisplay != nil {
		in, out := &in.WarningDisplay, &out.WarningDisplay
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThresholdsObservation.
func (in *ThresholdsObservation) DeepCopy() *ThresholdsObservation {
	if in == nil {
		return nil
	}
	out := new(ThresholdsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThresholdsParameters) DeepCopyInto(out *ThresholdsParameters) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(float64)
		**out = **in
	}
	if in.Timeframe != nil {
		in, out := &in.Timeframe, &out.Timeframe
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThresholdsParameters.
func (in *ThresholdsParameters) DeepCopy() *ThresholdsParameters {
	if in == nil {
		return nil
	}
	out := new(ThresholdsParameters)
	in.DeepCopyInto(out)
	return out
}
