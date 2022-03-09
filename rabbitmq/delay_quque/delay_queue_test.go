package delay_quque

import (
	"github.com/Zaric666/learning/rabbitmq"
	"log"
	"testing"
)

const (
	ADDR      = "amqp://guest:guest@localhost:5672/"
	EXCHANGE  = "x-delayed-message"
	QUEUENAME = "delay_queue"
)

func TestSend(t *testing.T) {
	producer := NewProducer(Config{ADDR, EXCHANGE, QUEUENAME, "", false})
	err := producer.Connect()
	rabbitmq.FailOnError(err, "connect fail")

	err = producer.Publish([]byte("order delay to push"), 10000)
	rabbitmq.FailOnError(err, "publish fail")
}

func TestReceive(t *testing.T) {
	ch := make(chan bool)
	consumer := NewConsumer(ADDR, EXCHANGE, QUEUENAME, "", false, func(msg []byte) error {
		log.Printf("Received a message: %s", string(msg))
		return nil
	})
	consumer.Start()
	<-ch
}
