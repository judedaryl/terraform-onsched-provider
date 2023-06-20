package main

import (
	"context"
	"flag"
	"log"
	"terraform-provider-onsched/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/judedaryl/onsched",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New("v1"), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
