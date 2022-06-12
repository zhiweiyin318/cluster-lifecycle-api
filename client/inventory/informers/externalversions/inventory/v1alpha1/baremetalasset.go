// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/stolostron/cluster-lifecycle-api/client/inventory/clientset/versioned"
	internalinterfaces "github.com/stolostron/cluster-lifecycle-api/client/inventory/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/stolostron/cluster-lifecycle-api/client/inventory/listers/inventory/v1alpha1"
	inventoryv1alpha1 "github.com/stolostron/cluster-lifecycle-api/inventory/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BareMetalAssetInformer provides access to a shared informer and lister for
// BareMetalAssets.
type BareMetalAssetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BareMetalAssetLister
}

type bareMetalAssetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBareMetalAssetInformer constructs a new informer for BareMetalAsset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBareMetalAssetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBareMetalAssetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBareMetalAssetInformer constructs a new informer for BareMetalAsset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBareMetalAssetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InventroyV1alpha1().BareMetalAssets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InventroyV1alpha1().BareMetalAssets(namespace).Watch(context.TODO(), options)
			},
		},
		&inventoryv1alpha1.BareMetalAsset{},
		resyncPeriod,
		indexers,
	)
}

func (f *bareMetalAssetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBareMetalAssetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bareMetalAssetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&inventoryv1alpha1.BareMetalAsset{}, f.defaultInformer)
}

func (f *bareMetalAssetInformer) Lister() v1alpha1.BareMetalAssetLister {
	return v1alpha1.NewBareMetalAssetLister(f.Informer().GetIndexer())
}
