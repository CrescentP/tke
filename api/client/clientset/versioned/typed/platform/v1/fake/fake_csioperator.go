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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeCSIOperators implements CSIOperatorInterface
type FakeCSIOperators struct {
	Fake *FakePlatformV1
}

var csioperatorsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "csioperators"}

var csioperatorsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "CSIOperator"}

// Get takes name of the cSIOperator, and returns the corresponding cSIOperator object, and an error if there is any.
func (c *FakeCSIOperators) Get(ctx context.Context, name string, options v1.GetOptions) (result *platformv1.CSIOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(csioperatorsResource, name), &platformv1.CSIOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CSIOperator), err
}

// List takes label and field selectors, and returns the list of CSIOperators that match those selectors.
func (c *FakeCSIOperators) List(ctx context.Context, opts v1.ListOptions) (result *platformv1.CSIOperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(csioperatorsResource, csioperatorsKind, opts), &platformv1.CSIOperatorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.CSIOperatorList{ListMeta: obj.(*platformv1.CSIOperatorList).ListMeta}
	for _, item := range obj.(*platformv1.CSIOperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIOperators.
func (c *FakeCSIOperators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(csioperatorsResource, opts))
}

// Create takes the representation of a cSIOperator and creates it.  Returns the server's representation of the cSIOperator, and an error, if there is any.
func (c *FakeCSIOperators) Create(ctx context.Context, cSIOperator *platformv1.CSIOperator, opts v1.CreateOptions) (result *platformv1.CSIOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(csioperatorsResource, cSIOperator), &platformv1.CSIOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CSIOperator), err
}

// Update takes the representation of a cSIOperator and updates it. Returns the server's representation of the cSIOperator, and an error, if there is any.
func (c *FakeCSIOperators) Update(ctx context.Context, cSIOperator *platformv1.CSIOperator, opts v1.UpdateOptions) (result *platformv1.CSIOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(csioperatorsResource, cSIOperator), &platformv1.CSIOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CSIOperator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCSIOperators) UpdateStatus(ctx context.Context, cSIOperator *platformv1.CSIOperator, opts v1.UpdateOptions) (*platformv1.CSIOperator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(csioperatorsResource, "status", cSIOperator), &platformv1.CSIOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CSIOperator), err
}

// Delete takes name of the cSIOperator and deletes it. Returns an error if one occurs.
func (c *FakeCSIOperators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(csioperatorsResource, name), &platformv1.CSIOperator{})
	return err
}

// Patch applies the patch and returns the patched cSIOperator.
func (c *FakeCSIOperators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platformv1.CSIOperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(csioperatorsResource, name, pt, data, subresources...), &platformv1.CSIOperator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.CSIOperator), err
}
