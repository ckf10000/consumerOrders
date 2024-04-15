// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     producer.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 02:07:44
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"github.com/ckf10000/gologger/v3/log"
	"github.com/streadway/amqp"
)

// SendToExceptionQueue 将消息发送到异常队列
func SendToQueue(msg *amqp.Delivery, rabbitMQConfig RabbitMQConfig, log *log.FileLogger, conn *amqp.Connection) error {
	channel, err := OpenChannel(conn, log)
	if err != nil {
		return err
	}
	DeclareExchange(channel, rabbitMQConfig.ProducerExchangeName, rabbitMQConfig.ProducerExchangeType, log)
	DeclareQueue(channel, rabbitMQConfig.ProducerQueueName, log)
	BindQueue(channel, rabbitMQConfig.ProducerQueueName, rabbitMQConfig.ProducerRoutingKey, rabbitMQConfig.ProducerExchangeName, log)
	err = PublishMessage(channel, rabbitMQConfig.ProducerQueueName, msg, log)
	if err != nil {
		return err
	}
	log.Info("Message sent to exception queue successfully.")

	return nil
}
