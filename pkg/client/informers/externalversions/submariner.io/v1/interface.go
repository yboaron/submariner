/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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
	internalinterfaces "github.com/submariner-io/submariner/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ClusterGlobalEgressIPs returns a ClusterGlobalEgressIPInformer.
	ClusterGlobalEgressIPs() ClusterGlobalEgressIPInformer
	// Endpoints returns a EndpointInformer.
	Endpoints() EndpointInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// GatewayRoutes returns a GatewayRouteInformer.
	GatewayRoutes() GatewayRouteInformer
	// GlobalEgressIPs returns a GlobalEgressIPInformer.
	GlobalEgressIPs() GlobalEgressIPInformer
	// GlobalIngressIPs returns a GlobalIngressIPInformer.
	GlobalIngressIPs() GlobalIngressIPInformer
	// NonGatewayRoutes returns a NonGatewayRouteInformer.
	NonGatewayRoutes() NonGatewayRouteInformer
	// RouteAgents returns a RouteAgentInformer.
	RouteAgents() RouteAgentInformer
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

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterGlobalEgressIPs returns a ClusterGlobalEgressIPInformer.
func (v *version) ClusterGlobalEgressIPs() ClusterGlobalEgressIPInformer {
	return &clusterGlobalEgressIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Endpoints returns a EndpointInformer.
func (v *version) Endpoints() EndpointInformer {
	return &endpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GatewayRoutes returns a GatewayRouteInformer.
func (v *version) GatewayRoutes() GatewayRouteInformer {
	return &gatewayRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GlobalEgressIPs returns a GlobalEgressIPInformer.
func (v *version) GlobalEgressIPs() GlobalEgressIPInformer {
	return &globalEgressIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GlobalIngressIPs returns a GlobalIngressIPInformer.
func (v *version) GlobalIngressIPs() GlobalIngressIPInformer {
	return &globalIngressIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NonGatewayRoutes returns a NonGatewayRouteInformer.
func (v *version) NonGatewayRoutes() NonGatewayRouteInformer {
	return &nonGatewayRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteAgents returns a RouteAgentInformer.
func (v *version) RouteAgents() RouteAgentInformer {
	return &routeAgentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
