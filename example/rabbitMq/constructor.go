package rabbitMq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const MQURL = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// routing Key
	RoutingKey string
	// MQ链接字符串
	Mqurl string
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}

func NewRabbitMQ(queueName, exchange, routingKey string) (rabbitMQ *RabbitMQ, err error) {
	rabbitMQ = &RabbitMQ{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
		Mqurl:      MQURL,
	}
	//创建 rabbitmq 连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "创建channel失败")

	return rabbitMQ, err
}

// 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *RabbitMQ) ReleaseRes() {
	mq.Conn.Close()
	mq.Channel.Close()
}
