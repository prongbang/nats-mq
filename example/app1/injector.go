package app1

import (
	nats "github.com/nats-io/go-nats"
	"github.com/prongbang/nats-mq/example/nat"
)

// ProvideNatMq is provide instance
func ProvideNatMq() nat.NatMQ {
	return nat.NewNatMQ()
}

func ProvideMqConnection() *nats.EncodedConn {
	con, _ := ProvideNatMq().EncodedConnect(nats.DefaultURL)
	return con
}

func ProvideDataSource() DataSource {
	return NewNatMqDataSource(ProvideMqConnection())
}

func ProvideRepository() Repository {
	return NewRepository(ProvideDataSource())
}

func ProvideDomain() UseCase {
	return NewUseCase(ProvideRepository())
}

func ProvideHandler() Handler {
	return NewHandler(ProvideDomain())
}

func ProvideRoute() Route {
	return NewRoute(ProvideHandler())
}
