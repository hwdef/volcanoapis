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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
	flowv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/flow/v1alpha1"
	scheme "volcano.sh/apis/pkg/client/clientset/versioned/scheme"
)

// JobTemplatesGetter has a method to return a JobTemplateInterface.
// A group's client should implement this interface.
type JobTemplatesGetter interface {
	JobTemplates(namespace string) JobTemplateInterface
}

// JobTemplateInterface has methods to work with JobTemplate resources.
type JobTemplateInterface interface {
	Create(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.CreateOptions) (*v1alpha1.JobTemplate, error)
	Update(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (*v1alpha1.JobTemplate, error)
	UpdateStatus(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (*v1alpha1.JobTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.JobTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.JobTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobTemplate, err error)
	Apply(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobTemplate, err error)
	ApplyStatus(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobTemplate, err error)
	JobTemplateExpansion
}

// jobTemplates implements JobTemplateInterface
type jobTemplates struct {
	client rest.Interface
	ns     string
}

// newJobTemplates returns a JobTemplates
func newJobTemplates(c *FlowV1alpha1Client, namespace string) *jobTemplates {
	return &jobTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the jobTemplate, and returns the corresponding jobTemplate object, and an error if there is any.
func (c *jobTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JobTemplate, err error) {
	result = &v1alpha1.JobTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of JobTemplates that match those selectors.
func (c *jobTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JobTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.JobTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("jobtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jobTemplates.
func (c *jobTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("jobtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a jobTemplate and creates it.  Returns the server's representation of the jobTemplate, and an error, if there is any.
func (c *jobTemplates) Create(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.CreateOptions) (result *v1alpha1.JobTemplate, err error) {
	result = &v1alpha1.JobTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("jobtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jobTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a jobTemplate and updates it. Returns the server's representation of the jobTemplate, and an error, if there is any.
func (c *jobTemplates) Update(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (result *v1alpha1.JobTemplate, err error) {
	result = &v1alpha1.JobTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(jobTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jobTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *jobTemplates) UpdateStatus(ctx context.Context, jobTemplate *v1alpha1.JobTemplate, opts v1.UpdateOptions) (result *v1alpha1.JobTemplate, err error) {
	result = &v1alpha1.JobTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(jobTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(jobTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the jobTemplate and deletes it. Returns an error if one occurs.
func (c *jobTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *jobTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("jobtemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched jobTemplate.
func (c *jobTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JobTemplate, err error) {
	result = &v1alpha1.JobTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied jobTemplate.
func (c *jobTemplates) Apply(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobTemplate, err error) {
	if jobTemplate == nil {
		return nil, fmt.Errorf("jobTemplate provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(jobTemplate)
	if err != nil {
		return nil, err
	}
	name := jobTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("jobTemplate.Name must be provided to Apply")
	}
	result = &v1alpha1.JobTemplate{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *jobTemplates) ApplyStatus(ctx context.Context, jobTemplate *flowv1alpha1.JobTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.JobTemplate, err error) {
	if jobTemplate == nil {
		return nil, fmt.Errorf("jobTemplate provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(jobTemplate)
	if err != nil {
		return nil, err
	}

	name := jobTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("jobTemplate.Name must be provided to Apply")
	}

	result = &v1alpha1.JobTemplate{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("jobtemplates").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
