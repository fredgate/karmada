// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePropagationBindings implements PropagationBindingInterface
type FakePropagationBindings struct {
	Fake *FakePolicyV1alpha1
	ns   string
}

var propagationbindingsResource = schema.GroupVersionResource{Group: "policy.karmada.io", Version: "v1alpha1", Resource: "propagationbindings"}

var propagationbindingsKind = schema.GroupVersionKind{Group: "policy.karmada.io", Version: "v1alpha1", Kind: "PropagationBinding"}

// Get takes name of the propagationBinding, and returns the corresponding propagationBinding object, and an error if there is any.
func (c *FakePropagationBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PropagationBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(propagationbindingsResource, c.ns, name), &v1alpha1.PropagationBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationBinding), err
}

// List takes label and field selectors, and returns the list of PropagationBindings that match those selectors.
func (c *FakePropagationBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PropagationBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(propagationbindingsResource, propagationbindingsKind, c.ns, opts), &v1alpha1.PropagationBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PropagationBindingList{ListMeta: obj.(*v1alpha1.PropagationBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.PropagationBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested propagationBindings.
func (c *FakePropagationBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(propagationbindingsResource, c.ns, opts))

}

// Create takes the representation of a propagationBinding and creates it.  Returns the server's representation of the propagationBinding, and an error, if there is any.
func (c *FakePropagationBindings) Create(ctx context.Context, propagationBinding *v1alpha1.PropagationBinding, opts v1.CreateOptions) (result *v1alpha1.PropagationBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(propagationbindingsResource, c.ns, propagationBinding), &v1alpha1.PropagationBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationBinding), err
}

// Update takes the representation of a propagationBinding and updates it. Returns the server's representation of the propagationBinding, and an error, if there is any.
func (c *FakePropagationBindings) Update(ctx context.Context, propagationBinding *v1alpha1.PropagationBinding, opts v1.UpdateOptions) (result *v1alpha1.PropagationBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(propagationbindingsResource, c.ns, propagationBinding), &v1alpha1.PropagationBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePropagationBindings) UpdateStatus(ctx context.Context, propagationBinding *v1alpha1.PropagationBinding, opts v1.UpdateOptions) (*v1alpha1.PropagationBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(propagationbindingsResource, "status", c.ns, propagationBinding), &v1alpha1.PropagationBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationBinding), err
}

// Delete takes name of the propagationBinding and deletes it. Returns an error if one occurs.
func (c *FakePropagationBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(propagationbindingsResource, c.ns, name), &v1alpha1.PropagationBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePropagationBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(propagationbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PropagationBindingList{})
	return err
}

// Patch applies the patch and returns the patched propagationBinding.
func (c *FakePropagationBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PropagationBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(propagationbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PropagationBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PropagationBinding), err
}