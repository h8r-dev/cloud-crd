/*
Copyright 2022.

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

package v1alpha1

import (
	"context"
	time "time"

	cloudv1alpha1 "github.com/h8r-dev/cloud-crd/api/cloud/v1alpha1"
	versioned "github.com/h8r-dev/cloud-crd/generated/cloud/clientset/versioned"
	internalinterfaces "github.com/h8r-dev/cloud-crd/generated/cloud/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/h8r-dev/cloud-crd/generated/cloud/listers/cloud/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EnvironmentInformer provides access to a shared informer and lister for
// Environments.
type EnvironmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EnvironmentLister
}

type environmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEnvironmentInformer constructs a new informer for Environment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEnvironmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEnvironmentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEnvironmentInformer constructs a new informer for Environment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEnvironmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().Environments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().Environments(namespace).Watch(context.TODO(), options)
			},
		},
		&cloudv1alpha1.Environment{},
		resyncPeriod,
		indexers,
	)
}

func (f *environmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEnvironmentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *environmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudv1alpha1.Environment{}, f.defaultInformer)
}

func (f *environmentInformer) Lister() v1alpha1.EnvironmentLister {
	return v1alpha1.NewEnvironmentLister(f.Informer().GetIndexer())
}
