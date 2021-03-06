// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	policyv1alpha1 "github.com/openshift-talks/k8s-go/controller/pkg/apis/policy/v1alpha1"
	versioned "github.com/openshift-talks/k8s-go/controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift-talks/k8s-go/controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift-talks/k8s-go/controller/pkg/generated/listers/policy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HealthCheckPolicyInformer provides access to a shared informer and lister for
// HealthCheckPolicies.
type HealthCheckPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HealthCheckPolicyLister
}

type healthCheckPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHealthCheckPolicyInformer constructs a new informer for HealthCheckPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHealthCheckPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHealthCheckPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHealthCheckPolicyInformer constructs a new informer for HealthCheckPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHealthCheckPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().HealthCheckPolicies(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PolicyV1alpha1().HealthCheckPolicies(namespace).Watch(options)
			},
		},
		&policyv1alpha1.HealthCheckPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *healthCheckPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHealthCheckPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *healthCheckPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&policyv1alpha1.HealthCheckPolicy{}, f.defaultInformer)
}

func (f *healthCheckPolicyInformer) Lister() v1alpha1.HealthCheckPolicyLister {
	return v1alpha1.NewHealthCheckPolicyLister(f.Informer().GetIndexer())
}
