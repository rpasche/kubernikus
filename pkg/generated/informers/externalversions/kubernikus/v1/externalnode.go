// This file was automatically generated by informer-gen

package v1

import (
	kubernikus_v1 "github.com/sapcc/kubernikus/pkg/apis/kubernikus/v1"
	clientset "github.com/sapcc/kubernikus/pkg/generated/clientset"
	internalinterfaces "github.com/sapcc/kubernikus/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/sapcc/kubernikus/pkg/generated/listers/kubernikus/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ExternalNodeInformer provides access to a shared informer and lister for
// ExternalNodes.
type ExternalNodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ExternalNodeLister
}

type externalNodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewExternalNodeInformer constructs a new informer for ExternalNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExternalNodeInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExternalNodeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredExternalNodeInformer constructs a new informer for ExternalNode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExternalNodeInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubernikusV1().ExternalNodes(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubernikusV1().ExternalNodes(namespace).Watch(options)
			},
		},
		&kubernikus_v1.ExternalNode{},
		resyncPeriod,
		indexers,
	)
}

func (f *externalNodeInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExternalNodeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *externalNodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubernikus_v1.ExternalNode{}, f.defaultInformer)
}

func (f *externalNodeInformer) Lister() v1.ExternalNodeLister {
	return v1.NewExternalNodeLister(f.Informer().GetIndexer())
}