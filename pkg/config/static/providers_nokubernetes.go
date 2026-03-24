//go:build nokubernetes

package static

// kubernetesProviders is empty when Kubernetes support is disabled.
type kubernetesProviders struct{}

// setKubernetesEffectiveConfiguration is a no-op when Kubernetes support is disabled.
func (c *Configuration) setKubernetesEffectiveConfiguration() {}

// validateKubernetesProviders is a no-op when Kubernetes support is disabled.
func (c *Configuration) validateKubernetesProviders() error { return nil }
