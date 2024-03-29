/*
Copyright The Kubernetes Authors.

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
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/sheikh-arman/controller-appscode-api/pkg/apis/appscode.com/v1alpha1"
	appscodecomv1alpha1 "github.com/sheikh-arman/controller-appscode-api/pkg/client/applyconfiguration/appscode.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEmployees implements EmployeeInterface
type FakeEmployees struct {
	Fake *FakeAppscodeV1alpha1
	ns   string
}

var employeesResource = v1alpha1.SchemeGroupVersion.WithResource("employees")

var employeesKind = v1alpha1.SchemeGroupVersion.WithKind("Employee")

// Get takes name of the employee, and returns the corresponding employee object, and an error if there is any.
func (c *FakeEmployees) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Employee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(employeesResource, c.ns, name), &v1alpha1.Employee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Employee), err
}

// List takes label and field selectors, and returns the list of Employees that match those selectors.
func (c *FakeEmployees) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EmployeeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(employeesResource, employeesKind, c.ns, opts), &v1alpha1.EmployeeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EmployeeList{ListMeta: obj.(*v1alpha1.EmployeeList).ListMeta}
	for _, item := range obj.(*v1alpha1.EmployeeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested employees.
func (c *FakeEmployees) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(employeesResource, c.ns, opts))

}

// Create takes the representation of a employee and creates it.  Returns the server's representation of the employee, and an error, if there is any.
func (c *FakeEmployees) Create(ctx context.Context, employee *v1alpha1.Employee, opts v1.CreateOptions) (result *v1alpha1.Employee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(employeesResource, c.ns, employee), &v1alpha1.Employee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Employee), err
}

// Update takes the representation of a employee and updates it. Returns the server's representation of the employee, and an error, if there is any.
func (c *FakeEmployees) Update(ctx context.Context, employee *v1alpha1.Employee, opts v1.UpdateOptions) (result *v1alpha1.Employee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(employeesResource, c.ns, employee), &v1alpha1.Employee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Employee), err
}

// Delete takes name of the employee and deletes it. Returns an error if one occurs.
func (c *FakeEmployees) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(employeesResource, c.ns, name, opts), &v1alpha1.Employee{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEmployees) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(employeesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EmployeeList{})
	return err
}

// Patch applies the patch and returns the patched employee.
func (c *FakeEmployees) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Employee, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(employeesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Employee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Employee), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied employee.
func (c *FakeEmployees) Apply(ctx context.Context, employee *appscodecomv1alpha1.EmployeeApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Employee, err error) {
	if employee == nil {
		return nil, fmt.Errorf("employee provided to Apply must not be nil")
	}
	data, err := json.Marshal(employee)
	if err != nil {
		return nil, err
	}
	name := employee.Name
	if name == nil {
		return nil, fmt.Errorf("employee.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(employeesResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.Employee{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Employee), err
}
