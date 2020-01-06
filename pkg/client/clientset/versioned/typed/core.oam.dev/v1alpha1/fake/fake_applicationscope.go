/*

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
	v1alpha1 "github.com/oam-dev/oam-go-sdk/apis/core.oam.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationScopes implements ApplicationScopeInterface
type FakeApplicationScopes struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var applicationscopesResource = schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1alpha1", Resource: "applicationscopes"}

var applicationscopesKind = schema.GroupVersionKind{Group: "core.oam.dev", Version: "v1alpha1", Kind: "ApplicationScope"}

// Get takes name of the applicationScope, and returns the corresponding applicationScope object, and an error if there is any.
func (c *FakeApplicationScopes) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationScope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationscopesResource, c.ns, name), &v1alpha1.ApplicationScope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationScope), err
}

// List takes label and field selectors, and returns the list of ApplicationScopes that match those selectors.
func (c *FakeApplicationScopes) List(opts v1.ListOptions) (result *v1alpha1.ApplicationScopeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationscopesResource, applicationscopesKind, c.ns, opts), &v1alpha1.ApplicationScopeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationScopeList{ListMeta: obj.(*v1alpha1.ApplicationScopeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationScopeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationScopes.
func (c *FakeApplicationScopes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationscopesResource, c.ns, opts))

}

// Create takes the representation of a applicationScope and creates it.  Returns the server's representation of the applicationScope, and an error, if there is any.
func (c *FakeApplicationScopes) Create(applicationScope *v1alpha1.ApplicationScope) (result *v1alpha1.ApplicationScope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationscopesResource, c.ns, applicationScope), &v1alpha1.ApplicationScope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationScope), err
}

// Update takes the representation of a applicationScope and updates it. Returns the server's representation of the applicationScope, and an error, if there is any.
func (c *FakeApplicationScopes) Update(applicationScope *v1alpha1.ApplicationScope) (result *v1alpha1.ApplicationScope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationscopesResource, c.ns, applicationScope), &v1alpha1.ApplicationScope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationScope), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationScopes) UpdateStatus(applicationScope *v1alpha1.ApplicationScope) (*v1alpha1.ApplicationScope, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationscopesResource, "status", c.ns, applicationScope), &v1alpha1.ApplicationScope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationScope), err
}

// Delete takes name of the applicationScope and deletes it. Returns an error if one occurs.
func (c *FakeApplicationScopes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationscopesResource, c.ns, name), &v1alpha1.ApplicationScope{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationScopes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationscopesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationScopeList{})
	return err
}

// Patch applies the patch and returns the patched applicationScope.
func (c *FakeApplicationScopes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationScope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationscopesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationScope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationScope), err
}