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

package v1

import (
	"context"
	quizquestionexamplecomv1 "example/build/apis/quizquestion.example.com/v1"
	versioned "example/build/client/clientset/versioned"
	internalinterfaces "example/build/client/informers/externalversions/internalinterfaces"
	v1 "example/build/client/listers/quizquestion.example.com/v1"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// QuizQuestionInformer provides access to a shared informer and lister for
// QuizQuestions.
type QuizQuestionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.QuizQuestionLister
}

type quizQuestionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewQuizQuestionInformer constructs a new informer for QuizQuestion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewQuizQuestionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredQuizQuestionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredQuizQuestionInformer constructs a new informer for QuizQuestion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredQuizQuestionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuizquestionExampleV1().QuizQuestions().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuizquestionExampleV1().QuizQuestions().Watch(context.TODO(), options)
			},
		},
		&quizquestionexamplecomv1.QuizQuestion{},
		resyncPeriod,
		indexers,
	)
}

func (f *quizQuestionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredQuizQuestionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *quizQuestionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quizquestionexamplecomv1.QuizQuestion{}, f.defaultInformer)
}

func (f *quizQuestionInformer) Lister() v1.QuizQuestionLister {
	return v1.NewQuizQuestionLister(f.Informer().GetIndexer())
}
