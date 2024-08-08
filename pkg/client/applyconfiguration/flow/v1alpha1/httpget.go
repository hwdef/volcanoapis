/*
Copyright 2021 The Volcano Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// HttpGetApplyConfiguration represents an declarative configuration of the HttpGet type for use
// with apply.
type HttpGetApplyConfiguration struct {
	TaskName   *string        `json:"taskName,omitempty"`
	Path       *string        `json:"path,omitempty"`
	Port       *int           `json:"port,omitempty"`
	HTTPHeader *v1.HTTPHeader `json:"httpHeader,omitempty"`
}

// HttpGetApplyConfiguration constructs an declarative configuration of the HttpGet type for use with
// apply.
func HttpGet() *HttpGetApplyConfiguration {
	return &HttpGetApplyConfiguration{}
}

// WithTaskName sets the TaskName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TaskName field is set to the value of the last call.
func (b *HttpGetApplyConfiguration) WithTaskName(value string) *HttpGetApplyConfiguration {
	b.TaskName = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *HttpGetApplyConfiguration) WithPath(value string) *HttpGetApplyConfiguration {
	b.Path = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *HttpGetApplyConfiguration) WithPort(value int) *HttpGetApplyConfiguration {
	b.Port = &value
	return b
}

// WithHTTPHeader sets the HTTPHeader field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPHeader field is set to the value of the last call.
func (b *HttpGetApplyConfiguration) WithHTTPHeader(value v1.HTTPHeader) *HttpGetApplyConfiguration {
	b.HTTPHeader = &value
	return b
}
