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

package versioned

import (
	"fmt"
	"net/http"
	configexamplev1 "example/build/client/clientset/versioned/typed/config.example.com/v1"
	evaluationexamplev1 "example/build/client/clientset/versioned/typed/evaluation.example.com/v1"
	eventexamplev1 "example/build/client/clientset/versioned/typed/event.example.com/v1"
	interestexamplev1 "example/build/client/clientset/versioned/typed/interest.example.com/v1"
	quizexamplev1 "example/build/client/clientset/versioned/typed/quiz.example.com/v1"
	quizchoiceexamplev1 "example/build/client/clientset/versioned/typed/quizchoice.example.com/v1"
	quizquestionexamplev1 "example/build/client/clientset/versioned/typed/quizquestion.example.com/v1"
	rootexamplev1 "example/build/client/clientset/versioned/typed/root.example.com/v1"
	tenantexamplev1 "example/build/client/clientset/versioned/typed/tenant.example.com/v1"
	userexamplev1 "example/build/client/clientset/versioned/typed/user.example.com/v1"
	wannaexamplev1 "example/build/client/clientset/versioned/typed/wanna.example.com/v1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ConfigExampleV1() configexamplev1.ConfigExampleV1Interface
	EvaluationExampleV1() evaluationexamplev1.EvaluationExampleV1Interface
	EventExampleV1() eventexamplev1.EventExampleV1Interface
	InterestExampleV1() interestexamplev1.InterestExampleV1Interface
	QuizExampleV1() quizexamplev1.QuizExampleV1Interface
	QuizchoiceExampleV1() quizchoiceexamplev1.QuizchoiceExampleV1Interface
	QuizquestionExampleV1() quizquestionexamplev1.QuizquestionExampleV1Interface
	RootExampleV1() rootexamplev1.RootExampleV1Interface
	TenantExampleV1() tenantexamplev1.TenantExampleV1Interface
	UserExampleV1() userexamplev1.UserExampleV1Interface
	WannaExampleV1() wannaexamplev1.WannaExampleV1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	configExampleV1       *configexamplev1.ConfigExampleV1Client
	evaluationExampleV1   *evaluationexamplev1.EvaluationExampleV1Client
	eventExampleV1        *eventexamplev1.EventExampleV1Client
	interestExampleV1     *interestexamplev1.InterestExampleV1Client
	quizExampleV1         *quizexamplev1.QuizExampleV1Client
	quizchoiceExampleV1   *quizchoiceexamplev1.QuizchoiceExampleV1Client
	quizquestionExampleV1 *quizquestionexamplev1.QuizquestionExampleV1Client
	rootExampleV1         *rootexamplev1.RootExampleV1Client
	tenantExampleV1       *tenantexamplev1.TenantExampleV1Client
	userExampleV1         *userexamplev1.UserExampleV1Client
	wannaExampleV1        *wannaexamplev1.WannaExampleV1Client
}

// ConfigExampleV1 retrieves the ConfigExampleV1Client
func (c *Clientset) ConfigExampleV1() configexamplev1.ConfigExampleV1Interface {
	return c.configExampleV1
}

// EvaluationExampleV1 retrieves the EvaluationExampleV1Client
func (c *Clientset) EvaluationExampleV1() evaluationexamplev1.EvaluationExampleV1Interface {
	return c.evaluationExampleV1
}

// EventExampleV1 retrieves the EventExampleV1Client
func (c *Clientset) EventExampleV1() eventexamplev1.EventExampleV1Interface {
	return c.eventExampleV1
}

// InterestExampleV1 retrieves the InterestExampleV1Client
func (c *Clientset) InterestExampleV1() interestexamplev1.InterestExampleV1Interface {
	return c.interestExampleV1
}

// QuizExampleV1 retrieves the QuizExampleV1Client
func (c *Clientset) QuizExampleV1() quizexamplev1.QuizExampleV1Interface {
	return c.quizExampleV1
}

// QuizchoiceExampleV1 retrieves the QuizchoiceExampleV1Client
func (c *Clientset) QuizchoiceExampleV1() quizchoiceexamplev1.QuizchoiceExampleV1Interface {
	return c.quizchoiceExampleV1
}

// QuizquestionExampleV1 retrieves the QuizquestionExampleV1Client
func (c *Clientset) QuizquestionExampleV1() quizquestionexamplev1.QuizquestionExampleV1Interface {
	return c.quizquestionExampleV1
}

// RootExampleV1 retrieves the RootExampleV1Client
func (c *Clientset) RootExampleV1() rootexamplev1.RootExampleV1Interface {
	return c.rootExampleV1
}

// TenantExampleV1 retrieves the TenantExampleV1Client
func (c *Clientset) TenantExampleV1() tenantexamplev1.TenantExampleV1Interface {
	return c.tenantExampleV1
}

// UserExampleV1 retrieves the UserExampleV1Client
func (c *Clientset) UserExampleV1() userexamplev1.UserExampleV1Interface {
	return c.userExampleV1
}

// WannaExampleV1 retrieves the WannaExampleV1Client
func (c *Clientset) WannaExampleV1() wannaexamplev1.WannaExampleV1Interface {
	return c.wannaExampleV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.configExampleV1, err = configexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.evaluationExampleV1, err = evaluationexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.eventExampleV1, err = eventexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.interestExampleV1, err = interestexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.quizExampleV1, err = quizexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.quizchoiceExampleV1, err = quizchoiceexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.quizquestionExampleV1, err = quizquestionexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.rootExampleV1, err = rootexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.tenantExampleV1, err = tenantexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.userExampleV1, err = userexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.wannaExampleV1, err = wannaexamplev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.configExampleV1 = configexamplev1.New(c)
	cs.evaluationExampleV1 = evaluationexamplev1.New(c)
	cs.eventExampleV1 = eventexamplev1.New(c)
	cs.interestExampleV1 = interestexamplev1.New(c)
	cs.quizExampleV1 = quizexamplev1.New(c)
	cs.quizchoiceExampleV1 = quizchoiceexamplev1.New(c)
	cs.quizquestionExampleV1 = quizquestionexamplev1.New(c)
	cs.rootExampleV1 = rootexamplev1.New(c)
	cs.tenantExampleV1 = tenantexamplev1.New(c)
	cs.userExampleV1 = userexamplev1.New(c)
	cs.wannaExampleV1 = wannaexamplev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
