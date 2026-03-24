//go:build !nokubernetes

package static

import (
	"errors"

	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd"
	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/gateway"
	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/ingress"
	ingressnginx "github.com/traefik/traefik/v3/pkg/provider/kubernetes/ingress-nginx"
	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/knative"
)

// kubernetesProviders contains Kubernetes-specific provider configuration.
// This is embedded in Providers struct when Kubernetes support is enabled.
type kubernetesProviders struct {
	KubernetesIngress      *ingress.Provider      `description:"Enables Kubernetes Ingress provider." json:"kubernetesIngress,omitempty" toml:"kubernetesIngress,omitempty" yaml:"kubernetesIngress,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	KubernetesIngressNGINX *ingressnginx.Provider `description:"Enables Kubernetes Ingress NGINX provider." json:"kubernetesIngressNGINX,omitempty" toml:"kubernetesIngressNGINX,omitempty" yaml:"kubernetesIngressNGINX,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	KubernetesCRD          *crd.Provider          `description:"Enables Kubernetes CRD provider." json:"kubernetesCRD,omitempty" toml:"kubernetesCRD,omitempty" yaml:"kubernetesCRD,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	KubernetesGateway      *gateway.Provider      `description:"Enables Kubernetes Gateway API provider." json:"kubernetesGateway,omitempty" toml:"kubernetesGateway,omitempty" yaml:"kubernetesGateway,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
	Knative                *knative.Provider      `description:"Enables Knative provider." json:"knative,omitempty" toml:"knative,omitempty" yaml:"knative,omitempty" label:"allowEmpty" file:"allowEmpty" export:"true"`
}

// setKubernetesEffectiveConfiguration handles Kubernetes-specific configuration.
func (c *Configuration) setKubernetesEffectiveConfiguration() {
	// Configure Gateway API provider
	if c.Providers.KubernetesGateway != nil {
		entryPoints := make(map[string]gateway.Entrypoint)
		for epName, entryPoint := range c.EntryPoints {
			entryPoints[epName] = gateway.Entrypoint{Address: entryPoint.GetAddress(), HasHTTPTLSConf: entryPoint.HTTP.TLS != nil}
		}

		if c.Providers.KubernetesCRD != nil {
			c.Providers.KubernetesCRD.FillExtensionBuilderRegistry(c.Providers.KubernetesGateway)
		}

		c.Providers.KubernetesGateway.EntryPoints = entryPoints
	}

	// Configure Ingress NGINX provider.
	if c.Providers.KubernetesIngressNGINX != nil {
		var nonTLSEntryPoints []string
		for epName, entryPoint := range c.EntryPoints {
			if entryPoint.HTTP.TLS == nil {
				nonTLSEntryPoints = append(nonTLSEntryPoints, epName)
			}
		}

		c.Providers.KubernetesIngressNGINX.NonTLSEntryPoints = nonTLSEntryPoints
	}

	// Defines the default rule syntax for the Kubernetes Ingress Provider.
	if c.Core != nil && c.Providers.KubernetesIngress != nil {
		c.Providers.KubernetesIngress.DefaultRuleSyntax = c.Core.DefaultRuleSyntax
	}
}

// validateKubernetesProviders validates Kubernetes provider configuration.
func (c *Configuration) validateKubernetesProviders() error {
	if c.Providers != nil && c.Providers.KubernetesIngressNGINX != nil {
		if c.Providers.KubernetesIngressNGINX.WatchNamespace != "" && c.Providers.KubernetesIngressNGINX.WatchNamespaceSelector != "" {
			return errors.New("watchNamespace and watchNamespaceSelector options are mutually exclusive")
		}
	}

	if c.Providers != nil && c.Providers.Knative != nil {
		if c.Experimental == nil || !c.Experimental.Knative {
			return errors.New("the experimental Knative feature must be enabled to use the Knative provider")
		}
	}

	return nil
}
