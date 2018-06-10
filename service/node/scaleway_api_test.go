package node

import (
	"testing"

	scaleway "github.com/scaleway/go-scaleway"
	"github.com/scaleway/go-scaleway/types"
)

func TestAPI(t *testing.T) {
	t.Run("", func(t *testing.T) {
		organization := ""
		token := ""
		userAgent := ""
		region := "par1"

		api, err := scaleway.NewScalewayAPI(organization, token, userAgent, region)
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}

		imageIdentifier, err := api.GetImageID("image:ubuntu-mini-xenial-25g", "x86_64")
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}

		def := types.ScalewayServerDefinition{
			Name:           "test",
			CommercialType: "START1-XS",
			Image:          &imageIdentifier.Identifier,
		}
		identifier, err := api.PostServer(def)
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}

		server, err := api.GetServerID(identifier)
		if err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}

		if err = api.PostServerAction(server, "poweron"); err != nil {
			t.Errorf("got: %v\nwant: %v", err, nil)
		}

		t.Logf("identifier: %v", identifier)
	})
}
