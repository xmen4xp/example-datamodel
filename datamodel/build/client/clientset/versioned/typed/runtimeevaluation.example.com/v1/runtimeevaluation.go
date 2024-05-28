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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	v1 "example/build/apis/runtimeevaluation.example.com/v1"
	scheme "example/build/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RuntimeEvaluationsGetter has a method to return a RuntimeEvaluationInterface.
// A group's client should implement this interface.
type RuntimeEvaluationsGetter interface {
	RuntimeEvaluations() RuntimeEvaluationInterface
}

// RuntimeEvaluationInterface has methods to work with RuntimeEvaluation resources.
type RuntimeEvaluationInterface interface {
	Create(ctx context.Context, runtimeEvaluation *v1.RuntimeEvaluation, opts metav1.CreateOptions) (*v1.RuntimeEvaluation, error)
	Update(ctx context.Context, runtimeEvaluation *v1.RuntimeEvaluation, opts metav1.UpdateOptions) (*v1.RuntimeEvaluation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RuntimeEvaluation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RuntimeEvaluationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeEvaluation, err error)
	RuntimeEvaluationExpansion
}

// runtimeEvaluations implements RuntimeEvaluationInterface
type runtimeEvaluations struct {
	client rest.Interface
}

// newRuntimeEvaluations returns a RuntimeEvaluations
func newRuntimeEvaluations(c *RuntimeevaluationExampleV1Client) *runtimeEvaluations {
	return &runtimeEvaluations{
		client: c.RESTClient(),
	}
}

// Get takes name of the runtimeEvaluation, and returns the corresponding runtimeEvaluation object, and an error if there is any.
func (c *runtimeEvaluations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RuntimeEvaluation, err error) {
	result = &v1.RuntimeEvaluation{}
	err = c.client.Get().
		Resource("runtimeevaluations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RuntimeEvaluations that match those selectors.
func (c *runtimeEvaluations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RuntimeEvaluationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RuntimeEvaluationList{}
	err = c.client.Get().
		Resource("runtimeevaluations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested runtimeEvaluations.
func (c *runtimeEvaluations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("runtimeevaluations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a runtimeEvaluation and creates it.  Returns the server's representation of the runtimeEvaluation, and an error, if there is any.
func (c *runtimeEvaluations) Create(ctx context.Context, runtimeEvaluation *v1.RuntimeEvaluation, opts metav1.CreateOptions) (result *v1.RuntimeEvaluation, err error) {
	result = &v1.RuntimeEvaluation{}
	err = c.client.Post().
		Resource("runtimeevaluations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runtimeEvaluation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a runtimeEvaluation and updates it. Returns the server's representation of the runtimeEvaluation, and an error, if there is any.
func (c *runtimeEvaluations) Update(ctx context.Context, runtimeEvaluation *v1.RuntimeEvaluation, opts metav1.UpdateOptions) (result *v1.RuntimeEvaluation, err error) {
	result = &v1.RuntimeEvaluation{}
	err = c.client.Put().
		Resource("runtimeevaluations").
		Name(runtimeEvaluation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runtimeEvaluation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the runtimeEvaluation and deletes it. Returns an error if one occurs.
func (c *runtimeEvaluations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("runtimeevaluations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *runtimeEvaluations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("runtimeevaluations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched runtimeEvaluation.
func (c *runtimeEvaluations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeEvaluation, err error) {
	result = &v1.RuntimeEvaluation{}
	err = c.client.Patch(pt).
		Resource("runtimeevaluations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}