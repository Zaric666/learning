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
	conn, err := amqp.Dial("amqp://guest:guest@localhost:/")
	rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 创建channel
	ch, err := conn.Channel()
	rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 建立confirm监听
	ch.Confirm(false)
	confirmChan := ch.NotifyPublish(make(chan amqp.Confirmation))

	// 4. 建立return监听
	go func(ch *amqp.Channel) {
		// NotifyReturn 会导致整个程序最后无法完全退出，所以使用goroutine
		returnChan := ch.NotifyReturn(make(chan amqp.Return))
		for re := range returnChan {
			fmt.Println("---------handle  return----------")
			fmt.Printf("replyCode: %d\n", re.ReplyCode)
			fmt.Println("replyText: " + re.ReplyText)
			fmt.Println("exchange: " + re.Exchange)
			fmt.Println("routingKey: " + re.RoutingKey)
			//fmt.Printf("properties: %d\n", re.Priority)
			fmt.Println("body: " + string(re.Body))
		}
	}(ch)

	// 5. 声明
	exchange := "test_return_exchange"
	//routingKey := "return.save"
	routingKey := "abc.save"
	body := "Hello RabbitMQ Send Return message!"

	// 6. 发送消息
	// 注意：mandatory设置为true
	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		true,       // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	rabbitmq.FailOnError(err, "Failed to publish a message")
	// 7. 监听抵达Broker无误后的确认信息,设置5秒超时
	ticker := time.NewTicker(2 * time.Second)
	select {
	case confirm := <-confirmChan:
		if confirm.Ack {
			fmt.Println("Push confirmed!")
		} else {
			fmt.Println("Push failed!")
		}
	case <-ticker.C:
		fmt.Println("out of limit time!")
		// 当同时开启 confirm 和 return 监听时，此时return失败，会导致confirm超时。
		// 由此也可以在这里进行校验
	}
}
