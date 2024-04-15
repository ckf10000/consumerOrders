// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     consumer.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 02:06:30
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"github.com/ckf10000/gologger/v3/log"
)

// StartConsumer 启动消费者
func StartConsumer(rabbitMQConfig RabbitMQConfig, handler MessageHandler, mysqlURI string, log *log.FileLogger) error {
	conn, err := ConnectMQ(rabbitMQConfig, log)
	if err != nil {
		return err
	}
	defer conn.Close()

	channel, err := OpenChannel(conn, log)
	if err != nil {
		return err
	}
	defer channel.Close()

	DeclareExchange(channel, rabbitMQConfig.ExchangeName, rabbitMQConfig.ExchangeType, log)
	DeclareQueue(channel, rabbitMQConfig.QueueName, log)
	BindQueue(channel, rabbitMQConfig.QueueName, rabbitMQConfig.RoutingKey, rabbitMQConfig.ExchangeName, log)

	db, err := ConnectMysql(mysqlURI, log)
	if err != nil {
		return err
	}
	defer db.Close()

	// 消费 RabbitMQ 消息
	msgs, err := ConsumeMessage(channel, rabbitMQConfig.QueueName, log)
	if err != nil {
		return err
	}

	for msg := range msgs {
		err := handler(&msg, db, log)
		if err != nil {
			// 发送消息到异常队列
			err = SendToQueue(&msg, rabbitMQConfig, log, conn)
			if err != nil {
				log.Error("Failed to send message to producer queue: %v", err)
			}
		}
	}

	return nil
}
