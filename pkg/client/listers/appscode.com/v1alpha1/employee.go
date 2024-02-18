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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sheikh-arman/controller-appscode-api/pkg/apis/appscode.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EmployeeLister helps list Employees.
// All objects returned here must be treated as read-only.
type EmployeeLister interface {
	// List lists all Employees in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Employee, err error)
	// Employees returns an object that can list and get Employees.
	Employees(namespace string) EmployeeNamespaceLister
	EmployeeListerExpansion
}

// employeeLister implements the EmployeeLister interface.
type employeeLister struct {
	indexer cache.Indexer
}

// NewEmployeeLister returns a new EmployeeLister.
func NewEmployeeLister(indexer cache.Indexer) EmployeeLister {
	return &employeeLister{indexer: indexer}
}

// List lists all Employees in the indexer.
func (s *employeeLister) List(selector labels.Selector) (ret []*v1alpha1.Employee, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Employee))
	})
	return ret, err
}

// Employees returns an object that can list and get Employees.
func (s *employeeLister) Employees(namespace string) EmployeeNamespaceLister {
	return employeeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EmployeeNamespaceLister helps list and get Employees.
// All objects returned here must be treated as read-only.
type EmployeeNamespaceLister interface {
	// List lists all Employees in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Employee, err error)
	// Get retrieves the Employee from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Employee, error)
	EmployeeNamespaceListerExpansion
}

// employeeNamespaceLister implements the EmployeeNamespaceLister
// interface.
type employeeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Employees in the indexer for a given namespace.
func (s employeeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Employee, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Employee))
	})
	return ret, err
}

// Get retrieves the Employee from the indexer for a given namespace and name.
func (s employeeNamespaceLister) Get(name string) (*v1alpha1.Employee, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("employee"), name)
	}
	return obj.(*v1alpha1.Employee), nil
}