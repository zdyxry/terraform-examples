package main
import (
	"context"
	"terraform-provider-hashicups-pf/hashicups"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider document generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name hashicups

func main() {
	providerserver.Serve(context.Background(), hashicups.New, providerserver.ServeOpts{
        Address: "hashicorp.com/edu/hashicups-pf",
	})
}