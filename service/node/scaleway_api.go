package node

import (
	scaleway "github.com/scaleway/go-scaleway"
	"github.com/scaleway/go-scaleway/types"
)

// ScalewayAPI is ...
type ScalewayAPI interface {
	GetImageID(needle, arch string) (*ScalewayImageIdentifier, error)
	PostServer(definition ScalewayServerDefinition) (string, error)
	GetServerID(needle string) (string, error)
	PostServerAction(serverID, action string) error
}

// ScalewayImageIdentifier is ...
type ScalewayImageIdentifier = types.ScalewayImageIdentifier

// ScalewayServerDefinition is ...
type ScalewayServerDefinition = types.ScalewayServerDefinition

// ScalewayClient is ...
func ScalewayClient(organization string, token string, userAgent string, region string) ScalewayAPI {
	client, err := scaleway.NewScalewayAPI(organization, token, userAgent, region)
	if err != nil {
		panic(err)
	}
	return client
}
