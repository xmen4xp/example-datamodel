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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"
	v1 "example/build/apis/config.example.com/v1"
	interestexamplecomv1 "example/build/apis/interest.example.com/v1"
	rootexamplecomv1 "example/build/apis/root.example.com/v1"
	tenantexamplecomv1 "example/build/apis/tenant.example.com/v1"
	userexamplecomv1 "example/build/apis/user.example.com/v1"
	wannaexamplecomv1 "example/build/apis/wanna.example.com/v1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=config.example.com, Version=v1
	case v1.SchemeGroupVersion.WithResource("configs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.ConfigExample().V1().Configs().Informer()}, nil

		// Group=interest.example.com, Version=v1
	case interestexamplecomv1.SchemeGroupVersion.WithResource("interests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.InterestExample().V1().Interests().Informer()}, nil

		// Group=root.example.com, Version=v1
	case rootexamplecomv1.SchemeGroupVersion.WithResource("roots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.RootExample().V1().Roots().Informer()}, nil

		// Group=tenant.example.com, Version=v1
	case tenantexamplecomv1.SchemeGroupVersion.WithResource("tenants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.TenantExample().V1().Tenants().Informer()}, nil

		// Group=user.example.com, Version=v1
	case userexamplecomv1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.UserExample().V1().Users().Informer()}, nil

		// Group=wanna.example.com, Version=v1
	case wannaexamplecomv1.SchemeGroupVersion.WithResource("wannas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.WannaExample().V1().Wannas().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
