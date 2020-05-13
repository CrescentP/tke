/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	auth "tkestack.io/tke/api/auth"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
)

// PoliciesGetter has a method to return a PolicyInterface.
// A group's client should implement this interface.
type PoliciesGetter interface {
	Policies() PolicyInterface
}

// PolicyInterface has methods to work with Policy resources.
type PolicyInterface interface {
	Create(ctx context.Context, policy *auth.Policy, opts v1.CreateOptions) (*auth.Policy, error)
	Update(ctx context.Context, policy *auth.Policy, opts v1.UpdateOptions) (*auth.Policy, error)
	UpdateStatus(ctx context.Context, policy *auth.Policy, opts v1.UpdateOptions) (*auth.Policy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*auth.Policy, error)
	List(ctx context.Context, opts v1.ListOptions) (*auth.PolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *auth.Policy, err error)
	PolicyExpansion
}

// policies implements PolicyInterface
type policies struct {
	client rest.Interface
}

// newPolicies returns a Policies
func newPolicies(c *AuthClient) *policies {
	return &policies{
		client: c.RESTClient(),
	}
}

// Get takes name of the policy, and returns the corresponding policy object, and an error if there is any.
func (c *policies) Get(ctx context.Context, name string, options v1.GetOptions) (result *auth.Policy, err error) {
	result = &auth.Policy{}
	err = c.client.Get().
		Resource("policies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Policies that match those selectors.
func (c *policies) List(ctx context.Context, opts v1.ListOptions) (result *auth.PolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &auth.PolicyList{}
	err = c.client.Get().
		Resource("policies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested policies.
func (c *policies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("policies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a policy and creates it.  Returns the server's representation of the policy, and an error, if there is any.
func (c *policies) Create(ctx context.Context, policy *auth.Policy, opts v1.CreateOptions) (result *auth.Policy, err error) {
	result = &auth.Policy{}
	err = c.client.Post().
		Resource("policies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a policy and updates it. Returns the server's representation of the policy, and an error, if there is any.
func (c *policies) Update(ctx context.Context, policy *auth.Policy, opts v1.UpdateOptions) (result *auth.Policy, err error) {
	result = &auth.Policy{}
	err = c.client.Put().
		Resource("policies").
		Name(policy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *policies) UpdateStatus(ctx context.Context, policy *auth.Policy, opts v1.UpdateOptions) (result *auth.Policy, err error) {
	result = &auth.Policy{}
	err = c.client.Put().
		Resource("policies").
		Name(policy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the policy and deletes it. Returns an error if one occurs.
func (c *policies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("policies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *policies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("policies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched policy.
func (c *policies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *auth.Policy, err error) {
	result = &auth.Policy{}
	err = c.client.Patch(pt).
		Resource("policies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
