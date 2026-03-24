//go:build !nokubernetes

package imports

import (
	// Kubernetes providers
	_ "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd"
	_ "github.com/traefik/traefik/v3/pkg/provider/kubernetes/gateway"
	_ "github.com/traefik/traefik/v3/pkg/provider/kubernetes/ingress"
	_ "github.com/traefik/traefik/v3/pkg/provider/kubernetes/ingress-nginx"
	_ "github.com/traefik/traefik/v3/pkg/provider/kubernetes/knative"
)
