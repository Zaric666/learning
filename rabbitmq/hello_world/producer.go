package main

import (
	"github.com/Zaric666/learning/rabbitmq"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	rabbitmq.FailOnError(err, "连接失败")

	ch, err := conn.Channel()
	rabbitmq.FailOnError(err, "打开通道失败")

	q, err := ch.QueueDeclare("hello_queue", true, false, false, false, nil)
	rabbitmq.FailOnError(err, "声明队列失败")

	body := "hello world!"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{Body: []byte(body), ContentType: "text/plain"})
	rabbitmq.FailOnError(err, "发送消息失败")
}
