/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
	submarineriov1 "github.com/submariner-io/submariner/pkg/client/applyconfiguration/submariner.io/v1"
	scheme "github.com/submariner-io/submariner/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalEgressIPsGetter has a method to return a GlobalEgressIPInterface.
// A group's client should implement this interface.
type GlobalEgressIPsGetter interface {
	GlobalEgressIPs(namespace string) GlobalEgressIPInterface
}

// GlobalEgressIPInterface has methods to work with GlobalEgressIP resources.
type GlobalEgressIPInterface interface {
	Create(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.CreateOptions) (*v1.GlobalEgressIP, error)
	Update(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.UpdateOptions) (*v1.GlobalEgressIP, error)
	UpdateStatus(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.UpdateOptions) (*v1.GlobalEgressIP, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GlobalEgressIP, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GlobalEgressIPList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GlobalEgressIP, err error)
	Apply(ctx context.Context, globalEgressIP *submarineriov1.GlobalEgressIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GlobalEgressIP, err error)
	ApplyStatus(ctx context.Context, globalEgressIP *submarineriov1.GlobalEgressIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GlobalEgressIP, err error)
	GlobalEgressIPExpansion
}

// globalEgressIPs implements GlobalEgressIPInterface
type globalEgressIPs struct {
	client rest.Interface
	ns     string
}

// newGlobalEgressIPs returns a GlobalEgressIPs
func newGlobalEgressIPs(c *SubmarinerV1Client, namespace string) *globalEgressIPs {
	return &globalEgressIPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the globalEgressIP, and returns the corresponding globalEgressIP object, and an error if there is any.
func (c *globalEgressIPs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GlobalEgressIP, err error) {
	result = &v1.GlobalEgressIP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalegressips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalEgressIPs that match those selectors.
func (c *globalEgressIPs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GlobalEgressIPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GlobalEgressIPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalegressips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalEgressIPs.
func (c *globalEgressIPs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("globalegressips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalEgressIP and creates it.  Returns the server's representation of the globalEgressIP, and an error, if there is any.
func (c *globalEgressIPs) Create(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.CreateOptions) (result *v1.GlobalEgressIP, err error) {
	result = &v1.GlobalEgressIP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("globalegressips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalEgressIP).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalEgressIP and updates it. Returns the server's representation of the globalEgressIP, and an error, if there is any.
func (c *globalEgressIPs) Update(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.UpdateOptions) (result *v1.GlobalEgressIP, err error) {
	result = &v1.GlobalEgressIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalegressips").
		Name(globalEgressIP.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalEgressIP).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *globalEgressIPs) UpdateStatus(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.UpdateOptions) (result *v1.GlobalEgressIP, err error) {
	result = &v1.GlobalEgressIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalegressips").
		Name(globalEgressIP.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalEgressIP).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalEgressIP and deletes it. Returns an error if one occurs.
func (c *globalEgressIPs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalegressips").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalEgressIPs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalegressips").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalEgressIP.
func (c *globalEgressIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GlobalEgressIP, err error) {
	result = &v1.GlobalEgressIP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("globalegressips").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied globalEgressIP.
func (c *globalEgressIPs) Apply(ctx context.Context, globalEgressIP *submarineriov1.GlobalEgressIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GlobalEgressIP, err error) {
	if globalEgressIP == nil {
		return nil, fmt.Errorf("globalEgressIP provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(globalEgressIP)
	if err != nil {
		return nil, err
	}
	name := globalEgressIP.Name
	if name == nil {
		return nil, fmt.Errorf("globalEgressIP.Name must be provided to Apply")
	}
	result = &v1.GlobalEgressIP{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("globalegressips").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *globalEgressIPs) ApplyStatus(ctx context.Context, globalEgressIP *submarineriov1.GlobalEgressIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GlobalEgressIP, err error) {
	if globalEgressIP == nil {
		return nil, fmt.Errorf("globalEgressIP provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(globalEgressIP)
	if err != nil {
		return nil, err
	}

	name := globalEgressIP.Name
	if name == nil {
		return nil, fmt.Errorf("globalEgressIP.Name must be provided to Apply")
	}

	result = &v1.GlobalEgressIP{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("globalegressips").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
