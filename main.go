package main

import (
	"context"
	"flag"
	"log"
	"os"
	"terraform-provider-onsched/internal/provider"
	"terraform-provider-onsched/onsched"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func mainS() {
	client_id := os.Getenv("ONSCHED_CLIENT_ID")
	client_secret := os.Getenv("ONSCHED_CLIENT_SECRET")
	c := onsched.NewClient(onsched.Sandbox, client_id, client_secret)
	comp, _ := c.GetCompany()
	comp.City = "Tagbilaran"
	comp, _ = c.UpdateCompany(comp)
	log.Println(comp)

}

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "level2.com/level2/onsched",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New("v1"), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
