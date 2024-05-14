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
	clientset "example/build/client/clientset/versioned"
	configexamplev1 "example/build/client/clientset/versioned/typed/config.example.com/v1"
	fakeconfigexamplev1 "example/build/client/clientset/versioned/typed/config.example.com/v1/fake"
	eventexamplev1 "example/build/client/clientset/versioned/typed/event.example.com/v1"
	fakeeventexamplev1 "example/build/client/clientset/versioned/typed/event.example.com/v1/fake"
	interestexamplev1 "example/build/client/clientset/versioned/typed/interest.example.com/v1"
	fakeinterestexamplev1 "example/build/client/clientset/versioned/typed/interest.example.com/v1/fake"
	rootexamplev1 "example/build/client/clientset/versioned/typed/root.example.com/v1"
	fakerootexamplev1 "example/build/client/clientset/versioned/typed/root.example.com/v1/fake"
	tenantexamplev1 "example/build/client/clientset/versioned/typed/tenant.example.com/v1"
	faketenantexamplev1 "example/build/client/clientset/versioned/typed/tenant.example.com/v1/fake"
	userexamplev1 "example/build/client/clientset/versioned/typed/user.example.com/v1"
	fakeuserexamplev1 "example/build/client/clientset/versioned/typed/user.example.com/v1/fake"
	wannaexamplev1 "example/build/client/clientset/versioned/typed/wanna.example.com/v1"
	fakewannaexamplev1 "example/build/client/clientset/versioned/typed/wanna.example.com/v1/fake"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ConfigExampleV1 retrieves the ConfigExampleV1Client
func (c *Clientset) ConfigExampleV1() configexamplev1.ConfigExampleV1Interface {
	return &fakeconfigexamplev1.FakeConfigExampleV1{Fake: &c.Fake}
}

// EventExampleV1 retrieves the EventExampleV1Client
func (c *Clientset) EventExampleV1() eventexamplev1.EventExampleV1Interface {
	return &fakeeventexamplev1.FakeEventExampleV1{Fake: &c.Fake}
}

// InterestExampleV1 retrieves the InterestExampleV1Client
func (c *Clientset) InterestExampleV1() interestexamplev1.InterestExampleV1Interface {
	return &fakeinterestexamplev1.FakeInterestExampleV1{Fake: &c.Fake}
}

// RootExampleV1 retrieves the RootExampleV1Client
func (c *Clientset) RootExampleV1() rootexamplev1.RootExampleV1Interface {
	return &fakerootexamplev1.FakeRootExampleV1{Fake: &c.Fake}
}

// TenantExampleV1 retrieves the TenantExampleV1Client
func (c *Clientset) TenantExampleV1() tenantexamplev1.TenantExampleV1Interface {
	return &faketenantexamplev1.FakeTenantExampleV1{Fake: &c.Fake}
}

// UserExampleV1 retrieves the UserExampleV1Client
func (c *Clientset) UserExampleV1() userexamplev1.UserExampleV1Interface {
	return &fakeuserexamplev1.FakeUserExampleV1{Fake: &c.Fake}
}

// WannaExampleV1 retrieves the WannaExampleV1Client
func (c *Clientset) WannaExampleV1() wannaexamplev1.WannaExampleV1Interface {
	return &fakewannaexamplev1.FakeWannaExampleV1{Fake: &c.Fake}
}
