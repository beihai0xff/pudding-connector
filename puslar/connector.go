package pulsar

import (
	"context"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"

	"github.com/beihai0xff/pudding/api/gen/pudding/types/v1"
	type2 "github.com/beihai0xff/pudding/app/broker/pkg/types"
)

// RealTimeQueue impl the connector interface
type RealTimeQueue struct {
	pulsar *Client
}

// NewRealTimeQueue create a pulsar connector
func NewRealTimeQueue(client *Client) *RealTimeQueue {
	return &RealTimeQueue{
		pulsar: client,
	}
}

// Produce produce a Message to the queue in real time
func (q *RealTimeQueue) Produce(ctx context.Context, msg *types.Message) error {
	if msg.Payload == nil || len(msg.Payload) == 0 {
		return fmt.Errorf("message payload can not be empty")
	}
	return q.pulsar.Produce(ctx, msg.Topic, convertToPulsarProducerMessage(msg))
}

// NewConsumer consume Messages from the queue in real time
func (q *RealTimeQueue) NewConsumer(topic, group string, batchSize int, fn type2.HandleMessage) error {
	var f = func(ctx context.Context, msg pulsar.Message) error {
		return fn(ctx, convertPulsarMessageToDelayMessage(msg))
	}

	return q.pulsar.NewConsumer(topic, group, f)
}

// Close the queue
func (q *RealTimeQueue) Close() error {
	return nil
}
