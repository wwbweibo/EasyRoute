package EasyRoute

import (
	dapr "github.com/dapr/go-sdk/client"
	"log"
)

func InitDaprClient() dapr.Client {
	client, err := dapr.NewClient()
	if err != nil {
		log.Panicf("init dapr client err: %s", err)
	}
	return client
}
