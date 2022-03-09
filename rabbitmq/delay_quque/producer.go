package delay_quque

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/streadway/amqp"
)

// Config 链接配置
type Config struct {
	Addr, Exchange, Queue, RoutingKey string
	AutoDelete                        bool
}

// Producer rabbitmq 生产者
type Producer struct {
	conn       *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
	config     Config
	done       chan bool
	connErr    chan error
	channelErr chan *amqp.Error
}

// NewProducer 创建生产者
func NewProducer(config Config) *Producer {
	return &Producer{
		config:     config,
		done:       make(chan bool),
		connErr:    make(chan error),
		channelErr: make(chan *amqp.Error),
	}
}

// Connect 链接到 MQ 服务器
func (c *Producer) Connect() error {
	var err error
	if c.conn, err = amqp.Dial(c.config.Addr); err != nil {
		return err
	}

	if c.Channel, err = c.conn.Channel(); err != nil {
		_ = c.Close()
		return err
	}

	// watching tcp connect
	go c.WatchConnect()
	return nil
}

// Close to close remote mq server connection
func (c *Producer) Close() error {
	close(c.done)

	if !c.conn.IsClosed() {
		if err := c.conn.Close(); err != nil {
			log.Error("rabbitmq producer - connection close failed: ", err)
			return err
		}
	}
	return nil
}

// Publish 发送消息至mq
func (c *Producer) Publish(body []byte, delay int64) error {
	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}
	if delay >= 0 {
		publishing.Headers = amqp.Table{
			"x-delay": delay,
		}
	}
	err := c.Channel.Publish(c.config.Exchange, c.config.RoutingKey, true, false, publishing)
	if err != nil {
		target := &amqp.Error{}
		if errors.As(err, target) {
			c.channelErr <- target
		} else {
			c.connErr <- err
		}
	}
	return err
}

// PublishJSON 将对象JSON格式化后发送消息
func (c *Producer) PublishJSON(body interface{}, delay int64) error {
	if data, err := json.Marshal(body); err != nil {
		return err
	} else {
		return c.Publish(data, delay)
	}
}

// WatchConnect 监控 MQ 的链接状态
func (c *Producer) WatchConnect() {
	ticker := time.NewTicker(30 * time.Second) // every 30 second
	defer ticker.Stop()

	for {
		select {
		case err := <-c.connErr:
			log.Errorf("rabbitmq producer - connection notify close: %s", err.Error())
			c.ReConnect()

		case err := <-c.channelErr:
			log.Errorf("rabbitmq producer - channel notify close: %s", err.Error())
			c.ReConnect()

		case <-ticker.C:
			c.ReConnect()

		case <-c.done:
			log.Debug("auto detect connection is done")
			return
		}
	}
}

// ReConnect 根据当前链接状态判断是否需要重新连接，如果连接异常则尝试重新连接
func (c *Producer) ReConnect() {
	if c.conn == nil || (c.conn != nil && c.conn.IsClosed()) {
		log.Errorf("rabbitmq connection is closed try to reconnect")
		if err := c.Connect(); err != nil {
			log.Errorf("rabbitmq reconnect failed: %s", err.Error())
		} else {
			log.Infof("rabbitmq reconnect succeeded")
		}
	}
}
