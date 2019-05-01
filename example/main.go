package main

import (
	"github.com/prongbang/nats-mq/example/app1"
	"github.com/prongbang/nats-mq/example/app2"
)

func main() {
	go func() {
		app2.ProvideHandler().SubNotiPromotion()
	}()

	app1.ProvideRoute().Initial("8081")
}
