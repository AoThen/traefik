//go:build !nokubernetes

package aggregator

import "github.com/traefik/traefik/v3/pkg/config/static"

// addKubernetesProviders adds Kubernetes providers to the aggregator.
// This file is excluded when building with the nokubernetes tag.
func (p *ProviderAggregator) addKubernetesProviders(conf static.Providers) {
	if conf.KubernetesIngress != nil {
		p.quietAddProvider(conf.KubernetesIngress)
	}

	if conf.KubernetesIngressNGINX != nil {
		p.quietAddProvider(conf.KubernetesIngressNGINX)
	}

	if conf.KubernetesCRD != nil {
		p.quietAddProvider(conf.KubernetesCRD)
	}

	if conf.Knative != nil {
		p.quietAddProvider(conf.Knative)
	}

	if conf.KubernetesGateway != nil {
		p.quietAddProvider(conf.KubernetesGateway)
	}
}
