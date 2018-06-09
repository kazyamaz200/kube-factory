package dns

import cloudflare "github.com/cloudflare/cloudflare-go"

// CloudflareAPI is ...
type CloudflareAPI interface {
}

// CloudflareClient is ...
func CloudflareClient(key string, email string) CloudflareAPI {
	client, err := cloudflare.New(key, email)
	if err != nil {
		panic(err)
	}
	return client
}
