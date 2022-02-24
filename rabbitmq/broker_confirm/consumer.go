package main

import (
	"github.com/Zaric666/learning/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	// 1. 建立RabbitMQ连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 创建channel
	ch, err := conn.Channel()
	rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 声明exchange,routing key,queue name
	exchange := "test_confirm_exchange"
	routingKey := "confirm.#"
	queueName := "test_confirm_queue"

	// 4. 声明（创建）一个交换机
	err = ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare an exchange")

	// 5. 声明（创建）一个队列
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	rabbitmq.FailOnError(err, "Failed to declare a queue")

	// 6. 队列绑定
	err = ch.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil)
	rabbitmq.FailOnError(err, "Failed to bind a queue")

	// 7. RMQ Server主动把消息推给消费者
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	rabbitmq.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
