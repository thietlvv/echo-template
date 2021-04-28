package nats

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

var PubSub *nats.Conn

// InitNasts ...
func InitNats() error {
	natsURL := os.Getenv("NATS_URL")
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
	opts = setupConnOptions(opts)

	natsConn, err := nats.Connect(natsURL, opts...)
	if err != nil {
		fmt.Println("could not connect to NATS server: %w", err)
		return err
	}

	PubSub = natsConn

	// defer natsConn.Close()

	return nil
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal("Exiting, no servers available")
	}))
	return opts
}

// Pub publishes some data to the given topic.
func Pub(topic string, data []byte) error {
	return PubSub.Publish(topic, data)
}

// Sub subscribes the given callback function to the interested topic.
func Sub(topic string, cb func(data []byte)) (err error) {
	// Subscribe
	if _, err := PubSub.Subscribe(topic, func(m *nats.Msg) {
		cb(m.Data)
	}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func PrintMsg(m *nats.Msg) {
	log.Printf("Received on [%s]: '%s'", m.Subject, string(m.Data))
}
