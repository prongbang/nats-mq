package nat

import nats "github.com/nats-io/go-nats"

// NatMQ is the interface
type NatMQ interface {
	Connect(url string, options ...nats.Option) (*NatConnection, error)
	EncodedConnect(url string, options ...nats.Option) (*nats.EncodedConn, error)
}

type natMQ struct {
}

// NewNatMQ is new instance
func NewNatMQ() NatMQ {
	return &natMQ{}
}

func (n *natMQ) Connect(url string, options ...nats.Option) (*NatConnection, error) {

	nc, err := nats.Connect(nats.DefaultURL, options...)

	return &NatConnection{Conn: nc}, err
}

func (n *natMQ) EncodedConnect(url string, options ...nats.Option) (*nats.EncodedConn, error) {

	nc, _ := n.Connect(url, options...)

	return nats.NewEncodedConn(nc.Conn, nats.JSON_ENCODER)
}
