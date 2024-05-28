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
	v1 "example/build/apis/runtimeuser.example.com/v1"
	scheme "example/build/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RuntimeUsersGetter has a method to return a RuntimeUserInterface.
// A group's client should implement this interface.
type RuntimeUsersGetter interface {
	RuntimeUsers() RuntimeUserInterface
}

// RuntimeUserInterface has methods to work with RuntimeUser resources.
type RuntimeUserInterface interface {
	Create(ctx context.Context, runtimeUser *v1.RuntimeUser, opts metav1.CreateOptions) (*v1.RuntimeUser, error)
	Update(ctx context.Context, runtimeUser *v1.RuntimeUser, opts metav1.UpdateOptions) (*v1.RuntimeUser, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.RuntimeUser, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.RuntimeUserList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeUser, err error)
	RuntimeUserExpansion
}

// runtimeUsers implements RuntimeUserInterface
type runtimeUsers struct {
	client rest.Interface
}

// newRuntimeUsers returns a RuntimeUsers
func newRuntimeUsers(c *RuntimeuserExampleV1Client) *runtimeUsers {
	return &runtimeUsers{
		client: c.RESTClient(),
	}
}

// Get takes name of the runtimeUser, and returns the corresponding runtimeUser object, and an error if there is any.
func (c *runtimeUsers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RuntimeUser, err error) {
	result = &v1.RuntimeUser{}
	err = c.client.Get().
		Resource("runtimeusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RuntimeUsers that match those selectors.
func (c *runtimeUsers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RuntimeUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RuntimeUserList{}
	err = c.client.Get().
		Resource("runtimeusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested runtimeUsers.
func (c *runtimeUsers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("runtimeusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a runtimeUser and creates it.  Returns the server's representation of the runtimeUser, and an error, if there is any.
func (c *runtimeUsers) Create(ctx context.Context, runtimeUser *v1.RuntimeUser, opts metav1.CreateOptions) (result *v1.RuntimeUser, err error) {
	result = &v1.RuntimeUser{}
	err = c.client.Post().
		Resource("runtimeusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runtimeUser).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a runtimeUser and updates it. Returns the server's representation of the runtimeUser, and an error, if there is any.
func (c *runtimeUsers) Update(ctx context.Context, runtimeUser *v1.RuntimeUser, opts metav1.UpdateOptions) (result *v1.RuntimeUser, err error) {
	result = &v1.RuntimeUser{}
	err = c.client.Put().
		Resource("runtimeusers").
		Name(runtimeUser.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(runtimeUser).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the runtimeUser and deletes it. Returns an error if one occurs.
func (c *runtimeUsers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("runtimeusers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *runtimeUsers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("runtimeusers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched runtimeUser.
func (c *runtimeUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RuntimeUser, err error) {
	result = &v1.RuntimeUser{}
	err = c.client.Patch(pt).
		Resource("runtimeusers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
