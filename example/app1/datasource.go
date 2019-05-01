package app1

import (
	nats "github.com/nats-io/go-nats"
)

// DataSource is the interface
type DataSource interface {
	Publish(subject string, v interface{}) error
	Subscribe(subject string, cb nats.Handler) (*nats.Subscription, error)
}
