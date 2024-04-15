// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     mq.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 02:08:04
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"github.com/ckf10000/gologger/v3/log"
	"github.com/streadway/amqp"
)

// RabbitMQConfig 存储 RabbitMQ 连接信息
type RabbitMQConfig struct {
	AMQPURI              string
	ExchangeName         string
	ExchangeType         string
	QueueName            string
	RoutingKey           string
	ProducerQueueName    string
	ProducerExchangeName string
	ProducerExchangeType string
	ProducerRoutingKey   string
}

func ConnectMQ(rabbitMQConfig RabbitMQConfig, log *log.FileLogger) (*amqp.Connection, error) {
	log.Info("开始连接MQ：%s", rabbitMQConfig.AMQPURI)
	// 连接 RabbitMQ
	conn, err := amqp.Dial(rabbitMQConfig.AMQPURI)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	} else {
		log.Info("连接MQ正常.")
	}
	return conn, nil
}

func OpenChannel(conn *amqp.Connection, log *log.FileLogger) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		log.Error("Failed to open a channel: %v", err)
		return nil, err
	} else {
		log.Info("MQ通道已打开.")
	}
	return channel, nil
}

func DeclareExchange(channel *amqp.Channel, exchangeName, exchangeType string, log *log.FileLogger) error {
	// 声明 Exchange
	err := channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to declare an exchange: %v", err)
		return err
	} else {
		log.Info("MQ声明exchange: %s 完成.", exchangeName)
	}
	return nil
}

func DeclareQueue(channel *amqp.Channel, queueName string, log *log.FileLogger) error {
	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to declare a queue: %v", err)
		return err
	} else {
		log.Info("MQ声明队列: %s 完成.", queueName)
	}
	return nil
}

func BindQueue(channel *amqp.Channel, queueName, routingKey, exchangeName string, log *log.FileLogger) error {
	// 绑定 Queue 到 Exchange
	err := channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to bind queue to exchange: %v", err)
		return err
	} else {
		log.Info("MQ绑定队列: %s 至exchange: %s 完成.", queueName, exchangeName)
	}
	return nil
}

func ConsumeMessage(channel *amqp.Channel, queueName string, log *log.FileLogger) (<-chan amqp.Delivery, error) {
	// 消费 RabbitMQ 消息
	msgs, err := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to register a consumer: %v", err)
		return nil, err
	} else {
		log.Info("开始接收队列: %s 中的消息.", queueName)
	}
	return msgs, nil
}

func PublishMessage(channel *amqp.Channel, queueName string, msg *amqp.Delivery, log *log.FileLogger) error {
	err := channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg.Body,
		},
	)
	if err != nil {
		log.Error("msg publish to mq failed. %s", err)
		return err
	}
	return nil
}
