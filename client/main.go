package main;

import (
	"context"
	oidc "github.com/coreos/go-oidc"
)

var (
	clientID = "app"
	clientSecret = "5dc5583d-82d3-4440-acf1-1e7ac65983f0"
)

func main()  {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, issuer: "http://localhost:8080/auth/realms/demo")

	if err != null {
		log.Fatal(err);
	}

	config := oauth2.Config{
		ClientID: clientID
		ClientSecret: clientSecret
		Endpoint: provider.EndPoint()
	}
}
