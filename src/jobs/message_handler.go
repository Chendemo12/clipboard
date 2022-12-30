package jobs

import (
	"github.com/Shopify/sarama"
	"gitlab.cowave.com/gogo/functools/zaplog"
)

var console = zaplog.ConsoleLogger{}

func MessageHandler(msg *sarama.ConsumerMessage) {
	console.Info("Topic: " + string(msg.Topic) + "-Key: " + string(msg.Key) + "-Value: " + string(msg.Value))
}
