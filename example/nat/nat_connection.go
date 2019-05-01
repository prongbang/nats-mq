package nat

import nats "github.com/nats-io/go-nats"

// NatConnection is a model
type NatConnection struct {
	Conn *nats.Conn
}
