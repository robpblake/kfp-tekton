/*
Copyright 2020 The Knative Authors

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
	"context"

	v1alpha1 "github.com/kubeflow/kfp-tekton/tekton-catalog/pipeline-loops/pkg/apis/pipelineloop/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelineLoops implements PipelineLoopInterface
type FakePipelineLoops struct {
	Fake *FakeCustomV1alpha1
	ns   string
}

var pipelineloopsResource = schema.GroupVersionResource{Group: "custom.tekton.dev", Version: "v1alpha1", Resource: "pipelineloops"}

var pipelineloopsKind = schema.GroupVersionKind{Group: "custom.tekton.dev", Version: "v1alpha1", Kind: "PipelineLoop"}

// Get takes name of the pipelineLoop, and returns the corresponding pipelineLoop object, and an error if there is any.
func (c *FakePipelineLoops) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PipelineLoop, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelineloopsResource, c.ns, name), &v1alpha1.PipelineLoop{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineLoop), err
}

// List takes label and field selectors, and returns the list of PipelineLoops that match those selectors.
func (c *FakePipelineLoops) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PipelineLoopList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelineloopsResource, pipelineloopsKind, c.ns, opts), &v1alpha1.PipelineLoopList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PipelineLoopList{ListMeta: obj.(*v1alpha1.PipelineLoopList).ListMeta}
	for _, item := range obj.(*v1alpha1.PipelineLoopList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelineLoops.
func (c *FakePipelineLoops) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelineloopsResource, c.ns, opts))

}

// Create takes the representation of a pipelineLoop and creates it.  Returns the server's representation of the pipelineLoop, and an error, if there is any.
func (c *FakePipelineLoops) Create(ctx context.Context, pipelineLoop *v1alpha1.PipelineLoop, opts v1.CreateOptions) (result *v1alpha1.PipelineLoop, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelineloopsResource, c.ns, pipelineLoop), &v1alpha1.PipelineLoop{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineLoop), err
}

// Update takes the representation of a pipelineLoop and updates it. Returns the server's representation of the pipelineLoop, and an error, if there is any.
func (c *FakePipelineLoops) Update(ctx context.Context, pipelineLoop *v1alpha1.PipelineLoop, opts v1.UpdateOptions) (result *v1alpha1.PipelineLoop, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelineloopsResource, c.ns, pipelineLoop), &v1alpha1.PipelineLoop{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineLoop), err
}

// Delete takes name of the pipelineLoop and deletes it. Returns an error if one occurs.
func (c *FakePipelineLoops) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelineloopsResource, c.ns, name), &v1alpha1.PipelineLoop{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelineLoops) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelineloopsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PipelineLoopList{})
	return err
}

// Patch applies the patch and returns the patched pipelineLoop.
func (c *FakePipelineLoops) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PipelineLoop, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelineloopsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PipelineLoop{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PipelineLoop), err
}