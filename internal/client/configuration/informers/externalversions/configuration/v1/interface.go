/*
Copyright 2018 The Kong Authors.

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
	internalinterfaces "github.com/kong/kubernetes-ingress-controller/internal/client/configuration/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KongClusterPlugins returns a KongClusterPluginInformer.
	KongClusterPlugins() KongClusterPluginInformer
	// KongConsumers returns a KongConsumerInformer.
	KongConsumers() KongConsumerInformer
	// KongCredentials returns a KongCredentialInformer.
	KongCredentials() KongCredentialInformer
	// KongIngresses returns a KongIngressInformer.
	KongIngresses() KongIngressInformer
	// KongPlugins returns a KongPluginInformer.
	KongPlugins() KongPluginInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// KongClusterPlugins returns a KongClusterPluginInformer.
func (v *version) KongClusterPlugins() KongClusterPluginInformer {
	return &kongClusterPluginInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KongConsumers returns a KongConsumerInformer.
func (v *version) KongConsumers() KongConsumerInformer {
	return &kongConsumerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KongCredentials returns a KongCredentialInformer.
func (v *version) KongCredentials() KongCredentialInformer {
	return &kongCredentialInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KongIngresses returns a KongIngressInformer.
func (v *version) KongIngresses() KongIngressInformer {
	return &kongIngressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// KongPlugins returns a KongPluginInformer.
func (v *version) KongPlugins() KongPluginInformer {
	return &kongPluginInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
