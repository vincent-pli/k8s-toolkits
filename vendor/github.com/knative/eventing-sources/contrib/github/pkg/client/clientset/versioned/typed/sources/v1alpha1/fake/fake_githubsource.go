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
	v1alpha1 "github.com/knative/eventing-contrib/contrib/github/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGitHubSources implements GitHubSourceInterface
type FakeGitHubSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var githubsourcesResource = schema.GroupVersionResource{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Resource: "githubsources"}

var githubsourcesKind = schema.GroupVersionKind{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Kind: "GitHubSource"}

// Get takes name of the gitHubSource, and returns the corresponding gitHubSource object, and an error if there is any.
func (c *FakeGitHubSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GitHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(githubsourcesResource, c.ns, name), &v1alpha1.GitHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubSource), err
}

// List takes label and field selectors, and returns the list of GitHubSources that match those selectors.
func (c *FakeGitHubSources) List(opts v1.ListOptions) (result *v1alpha1.GitHubSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(githubsourcesResource, githubsourcesKind, c.ns, opts), &v1alpha1.GitHubSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GitHubSourceList{ListMeta: obj.(*v1alpha1.GitHubSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GitHubSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gitHubSources.
func (c *FakeGitHubSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(githubsourcesResource, c.ns, opts))

}

// Create takes the representation of a gitHubSource and creates it.  Returns the server's representation of the gitHubSource, and an error, if there is any.
func (c *FakeGitHubSources) Create(gitHubSource *v1alpha1.GitHubSource) (result *v1alpha1.GitHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(githubsourcesResource, c.ns, gitHubSource), &v1alpha1.GitHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubSource), err
}

// Update takes the representation of a gitHubSource and updates it. Returns the server's representation of the gitHubSource, and an error, if there is any.
func (c *FakeGitHubSources) Update(gitHubSource *v1alpha1.GitHubSource) (result *v1alpha1.GitHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(githubsourcesResource, c.ns, gitHubSource), &v1alpha1.GitHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGitHubSources) UpdateStatus(gitHubSource *v1alpha1.GitHubSource) (*v1alpha1.GitHubSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(githubsourcesResource, "status", c.ns, gitHubSource), &v1alpha1.GitHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubSource), err
}

// Delete takes name of the gitHubSource and deletes it. Returns an error if one occurs.
func (c *FakeGitHubSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(githubsourcesResource, c.ns, name), &v1alpha1.GitHubSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGitHubSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(githubsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GitHubSourceList{})
	return err
}

// Patch applies the patch and returns the patched gitHubSource.
func (c *FakeGitHubSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitHubSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(githubsourcesResource, c.ns, name, data, subresources...), &v1alpha1.GitHubSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GitHubSource), err
}
