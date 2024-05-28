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
	evaluationexamplev1 "example/build/client/clientset/versioned/typed/evaluation.example.com/v1"
	fakeevaluationexamplev1 "example/build/client/clientset/versioned/typed/evaluation.example.com/v1/fake"
	eventexamplev1 "example/build/client/clientset/versioned/typed/event.example.com/v1"
	fakeeventexamplev1 "example/build/client/clientset/versioned/typed/event.example.com/v1/fake"
	interestexamplev1 "example/build/client/clientset/versioned/typed/interest.example.com/v1"
	fakeinterestexamplev1 "example/build/client/clientset/versioned/typed/interest.example.com/v1/fake"
	quizexamplev1 "example/build/client/clientset/versioned/typed/quiz.example.com/v1"
	fakequizexamplev1 "example/build/client/clientset/versioned/typed/quiz.example.com/v1/fake"
	quizchoiceexamplev1 "example/build/client/clientset/versioned/typed/quizchoice.example.com/v1"
	fakequizchoiceexamplev1 "example/build/client/clientset/versioned/typed/quizchoice.example.com/v1/fake"
	quizquestionexamplev1 "example/build/client/clientset/versioned/typed/quizquestion.example.com/v1"
	fakequizquestionexamplev1 "example/build/client/clientset/versioned/typed/quizquestion.example.com/v1/fake"
	rootexamplev1 "example/build/client/clientset/versioned/typed/root.example.com/v1"
	fakerootexamplev1 "example/build/client/clientset/versioned/typed/root.example.com/v1/fake"
	runtimeexamplev1 "example/build/client/clientset/versioned/typed/runtime.example.com/v1"
	fakeruntimeexamplev1 "example/build/client/clientset/versioned/typed/runtime.example.com/v1/fake"
	runtimeanswerexamplev1 "example/build/client/clientset/versioned/typed/runtimeanswer.example.com/v1"
	fakeruntimeanswerexamplev1 "example/build/client/clientset/versioned/typed/runtimeanswer.example.com/v1/fake"
	runtimeevaluationexamplev1 "example/build/client/clientset/versioned/typed/runtimeevaluation.example.com/v1"
	fakeruntimeevaluationexamplev1 "example/build/client/clientset/versioned/typed/runtimeevaluation.example.com/v1/fake"
	runtimequizexamplev1 "example/build/client/clientset/versioned/typed/runtimequiz.example.com/v1"
	fakeruntimequizexamplev1 "example/build/client/clientset/versioned/typed/runtimequiz.example.com/v1/fake"
	runtimeuserexamplev1 "example/build/client/clientset/versioned/typed/runtimeuser.example.com/v1"
	fakeruntimeuserexamplev1 "example/build/client/clientset/versioned/typed/runtimeuser.example.com/v1/fake"
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

// EvaluationExampleV1 retrieves the EvaluationExampleV1Client
func (c *Clientset) EvaluationExampleV1() evaluationexamplev1.EvaluationExampleV1Interface {
	return &fakeevaluationexamplev1.FakeEvaluationExampleV1{Fake: &c.Fake}
}

// EventExampleV1 retrieves the EventExampleV1Client
func (c *Clientset) EventExampleV1() eventexamplev1.EventExampleV1Interface {
	return &fakeeventexamplev1.FakeEventExampleV1{Fake: &c.Fake}
}

// InterestExampleV1 retrieves the InterestExampleV1Client
func (c *Clientset) InterestExampleV1() interestexamplev1.InterestExampleV1Interface {
	return &fakeinterestexamplev1.FakeInterestExampleV1{Fake: &c.Fake}
}

// QuizExampleV1 retrieves the QuizExampleV1Client
func (c *Clientset) QuizExampleV1() quizexamplev1.QuizExampleV1Interface {
	return &fakequizexamplev1.FakeQuizExampleV1{Fake: &c.Fake}
}

// QuizchoiceExampleV1 retrieves the QuizchoiceExampleV1Client
func (c *Clientset) QuizchoiceExampleV1() quizchoiceexamplev1.QuizchoiceExampleV1Interface {
	return &fakequizchoiceexamplev1.FakeQuizchoiceExampleV1{Fake: &c.Fake}
}

// QuizquestionExampleV1 retrieves the QuizquestionExampleV1Client
func (c *Clientset) QuizquestionExampleV1() quizquestionexamplev1.QuizquestionExampleV1Interface {
	return &fakequizquestionexamplev1.FakeQuizquestionExampleV1{Fake: &c.Fake}
}

// RootExampleV1 retrieves the RootExampleV1Client
func (c *Clientset) RootExampleV1() rootexamplev1.RootExampleV1Interface {
	return &fakerootexamplev1.FakeRootExampleV1{Fake: &c.Fake}
}

// RuntimeExampleV1 retrieves the RuntimeExampleV1Client
func (c *Clientset) RuntimeExampleV1() runtimeexamplev1.RuntimeExampleV1Interface {
	return &fakeruntimeexamplev1.FakeRuntimeExampleV1{Fake: &c.Fake}
}

// RuntimeanswerExampleV1 retrieves the RuntimeanswerExampleV1Client
func (c *Clientset) RuntimeanswerExampleV1() runtimeanswerexamplev1.RuntimeanswerExampleV1Interface {
	return &fakeruntimeanswerexamplev1.FakeRuntimeanswerExampleV1{Fake: &c.Fake}
}

// RuntimeevaluationExampleV1 retrieves the RuntimeevaluationExampleV1Client
func (c *Clientset) RuntimeevaluationExampleV1() runtimeevaluationexamplev1.RuntimeevaluationExampleV1Interface {
	return &fakeruntimeevaluationexamplev1.FakeRuntimeevaluationExampleV1{Fake: &c.Fake}
}

// RuntimequizExampleV1 retrieves the RuntimequizExampleV1Client
func (c *Clientset) RuntimequizExampleV1() runtimequizexamplev1.RuntimequizExampleV1Interface {
	return &fakeruntimequizexamplev1.FakeRuntimequizExampleV1{Fake: &c.Fake}
}

// RuntimeuserExampleV1 retrieves the RuntimeuserExampleV1Client
func (c *Clientset) RuntimeuserExampleV1() runtimeuserexamplev1.RuntimeuserExampleV1Interface {
	return &fakeruntimeuserexamplev1.FakeRuntimeuserExampleV1{Fake: &c.Fake}
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
