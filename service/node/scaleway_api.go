package node

import scaleway "github.com/scaleway/go-scaleway"

// ScalewayAPI is ...
type ScalewayAPI interface {
}

// ScalewayClient is ...
func ScalewayClient(organization string, token string, userAgent string, region string) ScalewayAPI {
	client, err := scaleway.NewScalewayAPI(organization, token, userAgent, region)
	if err != nil {
		panic(err)
	}
	return client
}
