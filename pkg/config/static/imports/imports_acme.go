//go:build !noacme

package imports

import (
	// ACME certificate resolver
	_ "github.com/traefik/traefik/v3/pkg/provider/acme"
)
