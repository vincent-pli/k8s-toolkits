/*
Copyright 2019 The Knative Authors

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

package fake

import (
	v1alpha1 "github.com/knative/eventing-contrib/camel/source/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCamelSources implements CamelSourceInterface
type FakeCamelSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var camelsourcesResource = schema.GroupVersionResource{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Resource: "camelsources"}

var camelsourcesKind = schema.GroupVersionKind{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Kind: "CamelSource"}

// Get takes name of the camelSource, and returns the corresponding camelSource object, and an error if there is any.
func (c *FakeCamelSources) Get(name string, options v1.GetOptions) (result *v1alpha1.CamelSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(camelsourcesResource, c.ns, name), &v1alpha1.CamelSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CamelSource), err
}

// List takes label and field selectors, and returns the list of CamelSources that match those selectors.
func (c *FakeCamelSources) List(opts v1.ListOptions) (result *v1alpha1.CamelSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(camelsourcesResource, camelsourcesKind, c.ns, opts), &v1alpha1.CamelSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CamelSourceList{ListMeta: obj.(*v1alpha1.CamelSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CamelSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested camelSources.
func (c *FakeCamelSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(camelsourcesResource, c.ns, opts))

}

// Create takes the representation of a camelSource and creates it.  Returns the server's representation of the camelSource, and an error, if there is any.
func (c *FakeCamelSources) Create(camelSource *v1alpha1.CamelSource) (result *v1alpha1.CamelSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(camelsourcesResource, c.ns, camelSource), &v1alpha1.CamelSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CamelSource), err
}

// Update takes the representation of a camelSource and updates it. Returns the server's representation of the camelSource, and an error, if there is any.
func (c *FakeCamelSources) Update(camelSource *v1alpha1.CamelSource) (result *v1alpha1.CamelSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(camelsourcesResource, c.ns, camelSource), &v1alpha1.CamelSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CamelSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCamelSources) UpdateStatus(camelSource *v1alpha1.CamelSource) (*v1alpha1.CamelSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(camelsourcesResource, "status", c.ns, camelSource), &v1alpha1.CamelSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CamelSource), err
}

// Delete takes name of the camelSource and deletes it. Returns an error if one occurs.
func (c *FakeCamelSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(camelsourcesResource, c.ns, name), &v1alpha1.CamelSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCamelSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(camelsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CamelSourceList{})
	return err
}

// Patch applies the patch and returns the patched camelSource.
func (c *FakeCamelSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CamelSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(camelsourcesResource, c.ns, name, data, subresources...), &v1alpha1.CamelSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CamelSource), err
}
