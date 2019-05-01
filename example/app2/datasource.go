package app2

import nats "github.com/nats-io/go-nats"

// DataSource is the interface
type DataSource interface {
	Subscribe(subject string, cb nats.Handler) (*nats.Subscription, error)
}
