//go:build nokubernetes

package aggregator

import "github.com/traefik/traefik/v3/pkg/config/static"

// addKubernetesProviders is a no-op when Kubernetes support is disabled.
// Build with -tags nokubernetes to exclude Kubernetes providers.
func (p *ProviderAggregator) addKubernetesProviders(conf static.Providers) {
	// Kubernetes providers disabled at build time
}
