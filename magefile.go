//go:build mage
// +build mage

package main

import (
	//mage:import
	argoMg "github.com/dm0275/argo-config/magefiles"
)

func init() {
	argoMg.ArgoCDConfig = argoMg.ArgoCdConfig{
		Namespace:       "argocd",
		Version:         "v2.11.3",
		SSHKeyPath:      "~/.ssh/id_rsa",
		PortForwardPort: "8080",
	}
}
