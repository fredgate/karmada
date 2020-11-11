// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	propagationstrategyv1alpha1 "github.com/huawei-cloudnative/karmada/pkg/apis/propagationstrategy/v1alpha1"
	versioned "github.com/huawei-cloudnative/karmada/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/huawei-cloudnative/karmada/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/huawei-cloudnative/karmada/pkg/generated/listers/propagationstrategy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PropagationPolicyInformer provides access to a shared informer and lister for
// PropagationPolicies.
type PropagationPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PropagationPolicyLister
}

type propagationPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPropagationPolicyInformer constructs a new informer for PropagationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPropagationPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPropagationPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPropagationPolicyInformer constructs a new informer for PropagationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPropagationPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PropagationstrategyV1alpha1().PropagationPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PropagationstrategyV1alpha1().PropagationPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&propagationstrategyv1alpha1.PropagationPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *propagationPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPropagationPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *propagationPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&propagationstrategyv1alpha1.PropagationPolicy{}, f.defaultInformer)
}

func (f *propagationPolicyInformer) Lister() v1alpha1.PropagationPolicyLister {
	return v1alpha1.NewPropagationPolicyLister(f.Informer().GetIndexer())
}
