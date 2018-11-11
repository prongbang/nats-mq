package main

import (
	"fmt"
	"net/http"
	"time"

	nats "github.com/nats-io/go-nats"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func basicHandler(w http.ResponseWriter, r *http.Request) {

	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}

func encodedHandler(w http.ResponseWriter, r *http.Request) {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	// Simple Publisher
	c.Publish("foo", "Hello World")

	// Simple Async Subscriber
	c.Subscribe("foo", func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})

	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name    string
		Address string
		Age     int
	}

	// Go type Subscriber
	c.Subscribe("hello", func(p *person) {
		fmt.Printf("Received a person: %+v\n", p)
	})

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	// Go type Publisher
	c.Publish("hello", me)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/basic", basicHandler)
	http.HandleFunc("/encoded", encodedHandler)
	http.ListenAndServe(":8080", nil)
}
