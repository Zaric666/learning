package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type Producer struct {
	name           string
	logger         *log.Logger
	connection     *amqp.Connection
	channel        *amqp.Channel
	done           chan bool
	notifyClose    chan *amqp.Error
	notifyConfirm  chan amqp.Confirmation
	isConnected    bool
	isRunReconnect chan bool
}

const (
	reconnectDelay = 5 * time.Second // 连接断开后多久重连
	resendDelay    = 5 * time.Second // 消息发送失败后，多久重发
	resendTime     = 3               // 消息重发次数
)

var (
	errNotConnected  = errors.New("not connected to the producer")
	errAlreadyClosed = errors.New("already closed: not connected to the producer")
)

func NewProducer(name string, addr string) *Producer {
	producer := Producer{
		logger:         log.New(os.Stdout, "", log.LstdFlags),
		name:           name,
		done:           make(chan bool),
		isRunReconnect: make(chan bool),
	}
	go producer.handleReconnect(addr)
	<-producer.isRunReconnect
	return &producer
}

// 如果连接失败会不断重连
// 如果连接断开会重新连接
func (producer *Producer) handleReconnect(addr string) {
	firstRun := true
	for {
		producer.isConnected = false
		log.Println("Attempting to connect")
		for !producer.connect(addr) {
			log.Println("Failed to connect. Retrying...")
			time.Sleep(reconnectDelay)
		}
		if firstRun {
			producer.isRunReconnect <- true
			firstRun = false
		}
		select {
		case <-producer.done:
			return
		case <-producer.notifyClose:
		}
	}
}

// 连接rabbitmq，以生产者的name定义一个队列
func (producer *Producer) connect(addr string) bool {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		return false
	}
	ch.Confirm(false)
	_, err = ch.QueueDeclare(
		producer.name,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return false
	}
	producer.changeConnection(conn, ch)
	producer.isConnected = true
	log.Println("Connected!")
	return true
}

// 监听Rabbit channel的状态
func (producer *Producer) changeConnection(connection *amqp.Connection, channel *amqp.Channel) {
	producer.connection = connection
	producer.channel = channel
	// channels没有必要主动关闭。如果没有协程使用它，它会被垃圾收集器收拾
	producer.notifyClose = make(chan *amqp.Error)
	producer.notifyConfirm = make(chan amqp.Confirmation)
	producer.channel.NotifyClose(producer.notifyClose)
	producer.channel.NotifyPublish(producer.notifyConfirm)
}

// 三次重传的发消息
func (producer *Producer) Push(data []byte) error {
	if !producer.isConnected {
		return errors.New("failed to push push: not connected")
	}
	var currentTime = 0
	for {
		err := producer.UnsafePush(data)
		if err != nil {
			producer.logger.Println("Push failed. Retrying...")
			currentTime += 1
			if currentTime < resendTime {
				continue
			} else {
				return err
			}
		}
		ticker := time.NewTicker(resendDelay)
		select {
		case confirm := <-producer.notifyConfirm:
			if confirm.Ack {
				producer.logger.Println("Push confirmed!")
				return nil
			}
		case <-ticker.C:
		}
		producer.logger.Println("Push didn't confirm. Retrying...")
	}
}

// 发送出去，不管是否接受的到
func (producer *Producer) UnsafePush(data []byte) error {
	if !producer.isConnected {
		return errNotConnected
	}
	return producer.channel.Publish(
		"",            // Exchange
		producer.name, // Routing key
		false,         // Mandatory
		false,         // Immediate
		amqp.Publishing{
			DeliveryMode: 2,
			ContentType:  "application/json",
			Body:         data,
			Timestamp:    time.Now(),
		},
	)
}

// 关闭连接/信道
func (producer *Producer) Close() error {
	if !producer.isConnected {
		return errAlreadyClosed
	}
	err := producer.channel.Close()
	if err != nil {
		return err
	}
	err = producer.connection.Close()
	if err != nil {
		return err
	}
	close(producer.done)
	producer.isConnected = false
	return nil
}

/*func main() {
	producer := NewProducer("test001", "amqp://guest:guest@localhost:/")
	if err := producer.Push([]byte("hello")); err != nil {
		fmt.Print(err)
	}
}*/
