package main

import (
	"fmt"
	"github.com/Zaric666/learning/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	// 1. 创建RabbitMQ连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 创建channel
	ch, err := conn.Channel()
	rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 建立confirm监听
	ch.Confirm(false)
	confirmChan := ch.NotifyPublish(make(chan amqp.Confirmation))

	// 4. 声明
	exchange := "test_confirm_exchange"
	routingKey := "confirm.save"
	body := "Hello RabbitMQ Send confirm message!"

	// 5. 发送消息
	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	rabbitmq.FailOnError(err, "Failed to publish a message")

	// 6. 监听抵达Broker无误后的确认信息,设置5秒超时
	ticker := time.NewTicker(5 * time.Second)
	select {
	case confirm := <-confirmChan:
		if confirm.Ack {
			log.Printf("Push confirmed!")
		} else {
			log.Printf("Push failed!")
		}
	case <-ticker.C:
		fmt.Println("out of limit time!")
	}

}
