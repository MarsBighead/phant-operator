// Authors: Marsbighead <duanhmhy@126.com>
//
// Copyright (c) 2024 Marsbighead
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	v1alpha1 "phant-operator/pkg/apis/phant/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePostgresClusters implements PostgresClusterInterface
type FakePostgresClusters struct {
	Fake *FakePhantV1alpha1
	ns   string
}

var postgresclustersResource = v1alpha1.SchemeGroupVersion.WithResource("postgresclusters")

var postgresclustersKind = v1alpha1.SchemeGroupVersion.WithKind("PostgresCluster")

// Get takes name of the postgresCluster, and returns the corresponding postgresCluster object, and an error if there is any.
func (c *FakePostgresClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PostgresCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postgresclustersResource, c.ns, name), &v1alpha1.PostgresCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresCluster), err
}

// List takes label and field selectors, and returns the list of PostgresClusters that match those selectors.
func (c *FakePostgresClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PostgresClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postgresclustersResource, postgresclustersKind, c.ns, opts), &v1alpha1.PostgresClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PostgresClusterList{ListMeta: obj.(*v1alpha1.PostgresClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.PostgresClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postgresClusters.
func (c *FakePostgresClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postgresclustersResource, c.ns, opts))

}

// Create takes the representation of a postgresCluster and creates it.  Returns the server's representation of the postgresCluster, and an error, if there is any.
func (c *FakePostgresClusters) Create(ctx context.Context, postgresCluster *v1alpha1.PostgresCluster, opts v1.CreateOptions) (result *v1alpha1.PostgresCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postgresclustersResource, c.ns, postgresCluster), &v1alpha1.PostgresCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresCluster), err
}

// Update takes the representation of a postgresCluster and updates it. Returns the server's representation of the postgresCluster, and an error, if there is any.
func (c *FakePostgresClusters) Update(ctx context.Context, postgresCluster *v1alpha1.PostgresCluster, opts v1.UpdateOptions) (result *v1alpha1.PostgresCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postgresclustersResource, c.ns, postgresCluster), &v1alpha1.PostgresCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePostgresClusters) UpdateStatus(ctx context.Context, postgresCluster *v1alpha1.PostgresCluster, opts v1.UpdateOptions) (*v1alpha1.PostgresCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(postgresclustersResource, "status", c.ns, postgresCluster), &v1alpha1.PostgresCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresCluster), err
}

// Delete takes name of the postgresCluster and deletes it. Returns an error if one occurs.
func (c *FakePostgresClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(postgresclustersResource, c.ns, name, opts), &v1alpha1.PostgresCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostgresClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postgresclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PostgresClusterList{})
	return err
}

// Patch applies the patch and returns the patched postgresCluster.
func (c *FakePostgresClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PostgresCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postgresclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.PostgresCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresCluster), err
}
