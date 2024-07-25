/*
Copyright 2023 The Knative Authors

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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "knative.dev/operator/pkg/apis/operator/v1beta1"
	scheme "knative.dev/operator/pkg/client/clientset/versioned/scheme"
)

// KnativeServingsGetter has a method to return a KnativeServingInterface.
// A group's client should implement this interface.
type KnativeServingsGetter interface {
	KnativeServings(namespace string) KnativeServingInterface
}

// KnativeServingInterface has methods to work with KnativeServing resources.
type KnativeServingInterface interface {
	Create(ctx context.Context, knativeServing *v1beta1.KnativeServing, opts v1.CreateOptions) (*v1beta1.KnativeServing, error)
	Update(ctx context.Context, knativeServing *v1beta1.KnativeServing, opts v1.UpdateOptions) (*v1beta1.KnativeServing, error)
	UpdateStatus(ctx context.Context, knativeServing *v1beta1.KnativeServing, opts v1.UpdateOptions) (*v1beta1.KnativeServing, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.KnativeServing, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.KnativeServingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KnativeServing, err error)
	KnativeServingExpansion
}

// knativeServings implements KnativeServingInterface
type knativeServings struct {
	client rest.Interface
	ns     string
}

// newKnativeServings returns a KnativeServings
func newKnativeServings(c *OperatorV1beta1Client, namespace string) *knativeServings {
	return &knativeServings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the knativeServing, and returns the corresponding knativeServing object, and an error if there is any.
func (c *knativeServings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KnativeServing, err error) {
	result = &v1beta1.KnativeServing{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knativeservings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KnativeServings that match those selectors.
func (c *knativeServings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KnativeServingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.KnativeServingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knativeservings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested knativeServings.
func (c *knativeServings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("knativeservings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a knativeServing and creates it.  Returns the server's representation of the knativeServing, and an error, if there is any.
func (c *knativeServings) Create(ctx context.Context, knativeServing *v1beta1.KnativeServing, opts v1.CreateOptions) (result *v1beta1.KnativeServing, err error) {
	result = &v1beta1.KnativeServing{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("knativeservings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(knativeServing).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a knativeServing and updates it. Returns the server's representation of the knativeServing, and an error, if there is any.
func (c *knativeServings) Update(ctx context.Context, knativeServing *v1beta1.KnativeServing, opts v1.UpdateOptions) (result *v1beta1.KnativeServing, err error) {
	result = &v1beta1.KnativeServing{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knativeservings").
		Name(knativeServing.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(knativeServing).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *knativeServings) UpdateStatus(ctx context.Context, knativeServing *v1beta1.KnativeServing, opts v1.UpdateOptions) (result *v1beta1.KnativeServing, err error) {
	result = &v1beta1.KnativeServing{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knativeservings").
		Name(knativeServing.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(knativeServing).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the knativeServing and deletes it. Returns an error if one occurs.
func (c *knativeServings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knativeservings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *knativeServings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knativeservings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched knativeServing.
func (c *knativeServings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KnativeServing, err error) {
	result = &v1beta1.KnativeServing{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("knativeservings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
