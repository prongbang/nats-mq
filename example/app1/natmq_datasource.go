package app1

import (
	nats "github.com/nats-io/go-nats"
)

type natMqDataSource struct {
	MqConn *nats.EncodedConn
}

// NewNatMqDataSource is new instance
func NewNatMqDataSource(mqConn *nats.EncodedConn) DataSource {
	return &natMqDataSource{
		MqConn: mqConn,
	}
}

func (natMq *natMqDataSource) Publish(subject string, v interface{}) error {
	return natMq.MqConn.Publish(subject, v)
}

func (natMq *natMqDataSource) Subscribe(subject string, cb nats.Handler) (*nats.Subscription, error) {
	return natMq.MqConn.Subscribe(subject, cb)
}
