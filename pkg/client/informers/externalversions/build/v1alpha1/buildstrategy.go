// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	versioned "github.com/shipwright-io/build/pkg/client/clientset/versioned"
	internalinterfaces "github.com/shipwright-io/build/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/shipwright-io/build/pkg/client/listers/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BuildStrategyInformer provides access to a shared informer and lister for
// BuildStrategies.
type BuildStrategyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BuildStrategyLister
}

type buildStrategyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBuildStrategyInformer constructs a new informer for BuildStrategy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBuildStrategyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBuildStrategyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBuildStrategyInformer constructs a new informer for BuildStrategy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBuildStrategyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ShipwrightV1alpha1().BuildStrategies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ShipwrightV1alpha1().BuildStrategies(namespace).Watch(context.TODO(), options)
			},
		},
		&buildv1alpha1.BuildStrategy{},
		resyncPeriod,
		indexers,
	)
}

func (f *buildStrategyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBuildStrategyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *buildStrategyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&buildv1alpha1.BuildStrategy{}, f.defaultInformer)
}

func (f *buildStrategyInformer) Lister() v1alpha1.BuildStrategyLister {
	return v1alpha1.NewBuildStrategyLister(f.Informer().GetIndexer())
}