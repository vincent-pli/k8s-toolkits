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

package v1alpha1

import (
	v1alpha1 "github.com/knative/eventing-contrib/contrib/github/pkg/apis/sources/v1alpha1"
	scheme "github.com/knative/eventing-contrib/contrib/github/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GitHubSourcesGetter has a method to return a GitHubSourceInterface.
// A group's client should implement this interface.
type GitHubSourcesGetter interface {
	GitHubSources(namespace string) GitHubSourceInterface
}

// GitHubSourceInterface has methods to work with GitHubSource resources.
type GitHubSourceInterface interface {
	Create(*v1alpha1.GitHubSource) (*v1alpha1.GitHubSource, error)
	Update(*v1alpha1.GitHubSource) (*v1alpha1.GitHubSource, error)
	UpdateStatus(*v1alpha1.GitHubSource) (*v1alpha1.GitHubSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GitHubSource, error)
	List(opts v1.ListOptions) (*v1alpha1.GitHubSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitHubSource, err error)
	GitHubSourceExpansion
}

// gitHubSources implements GitHubSourceInterface
type gitHubSources struct {
	client rest.Interface
	ns     string
}

// newGitHubSources returns a GitHubSources
func newGitHubSources(c *SourcesV1alpha1Client, namespace string) *gitHubSources {
	return &gitHubSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitHubSource, and returns the corresponding gitHubSource object, and an error if there is any.
func (c *gitHubSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("githubsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitHubSources that match those selectors.
func (c *gitHubSources) List(opts v1.ListOptions) (result *v1alpha1.GitHubSourceList, err error) {
	result = &v1alpha1.GitHubSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitHubSources.
func (c *gitHubSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gitHubSource and creates it.  Returns the server's representation of the gitHubSource, and an error, if there is any.
func (c *gitHubSources) Create(gitHubSource *v1alpha1.GitHubSource) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("githubsources").
		Body(gitHubSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gitHubSource and updates it. Returns the server's representation of the gitHubSource, and an error, if there is any.
func (c *gitHubSources) Update(gitHubSource *v1alpha1.GitHubSource) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("githubsources").
		Name(gitHubSource.Name).
		Body(gitHubSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gitHubSources) UpdateStatus(gitHubSource *v1alpha1.GitHubSource) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("githubsources").
		Name(gitHubSource.Name).
		SubResource("status").
		Body(gitHubSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the gitHubSource and deletes it. Returns an error if one occurs.
func (c *gitHubSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("githubsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitHubSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("githubsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gitHubSource.
func (c *gitHubSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitHubSource, err error) {
	result = &v1alpha1.GitHubSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("githubsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
