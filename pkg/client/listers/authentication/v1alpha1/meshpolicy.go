/*
Portions Copyright 2018 The Kubernetes Authors.
Portions Copyright 2018 Aspen Mesh Authors.

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
	v1alpha1 "github.com/aspenmesh/istio-client-go/pkg/apis/authentication/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MeshPolicyLister helps list MeshPolicies.
type MeshPolicyLister interface {
	// List lists all MeshPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MeshPolicy, err error)
	// Get retrieves the MeshPolicy from the index for a given name.
	Get(name string) (*v1alpha1.MeshPolicy, error)
	MeshPolicyListerExpansion
}

// meshPolicyLister implements the MeshPolicyLister interface.
type meshPolicyLister struct {
	indexer cache.Indexer
}

// NewMeshPolicyLister returns a new MeshPolicyLister.
func NewMeshPolicyLister(indexer cache.Indexer) MeshPolicyLister {
	return &meshPolicyLister{indexer: indexer}
}

// List lists all MeshPolicies in the indexer.
func (s *meshPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.MeshPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MeshPolicy))
	})
	return ret, err
}

// Get retrieves the MeshPolicy from the index for a given name.
func (s *meshPolicyLister) Get(name string) (*v1alpha1.MeshPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("meshpolicy"), name)
	}
	return obj.(*v1alpha1.MeshPolicy), nil
}