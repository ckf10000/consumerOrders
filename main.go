// Package main
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     main.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/13 15:42:50
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package main

import (
	"consumerOrders/internal"
	"fmt"

	"github.com/ckf10000/gologger/v3/log"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
)

func main() {
	// 获取可执行文件所在目录的路径
	exeDir := log.GetExecuteFilePath()
	if exeDir == "" {
		return
	}

	// projectPath := "./"
	log := log.NewLogger("debug", exeDir, "comsumer.log", "simple", 50*1024*1024, true, true, true)
	cache := internal.GetApolloCache(log)
	log.Info("开始启动Consumer...")
	// RabbitMQ 连接信息
	amqpURI, _ := cache.Get("amqpURI")
	exchangeName, _ := cache.Get("exchangeName")
	exchangeType, _ := cache.Get("exchangeType")
	queueName, _ := cache.Get("queueName")
	routingKey, _ := cache.Get("routingKey")
	producerExchangeName, _ := cache.Get("producerExchangeName")
	producerExchangeType, _ := cache.Get("producerExchangeType")
	producerQueueName, _ := cache.Get("producerQueueName")
	producerRoutingKey, _ := cache.Get("producerRoutingKey")

	rabbitMQConfig := internal.RabbitMQConfig{
		AMQPURI:              amqpURI.(string),
		ExchangeName:         exchangeName.(string),
		ExchangeType:         exchangeType.(string),
		QueueName:            queueName.(string),
		RoutingKey:           routingKey.(string),
		ProducerQueueName:    producerQueueName.(string),
		ProducerRoutingKey:   producerRoutingKey.(string),
		ProducerExchangeName: producerExchangeName.(string),
		ProducerExchangeType: producerExchangeType.(string),
	}

	// MySQL 连接信息
	mysqlUser, _ := cache.Get("mysqlUser")
	mysqlPassword, _ := cache.Get("mysqlPassword")
	mysqlHost, _ := cache.Get("mysqlHost")

	mysqlPort, _ := cache.Get("mysqlPort")
	mysqlDatabase, _ := cache.Get("mysqlDatabase")
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	err := internal.StartConsumer(rabbitMQConfig, internal.OrderMessageHandler, mysqlURI, log)
	if err != nil {
		log.Error("Consumer encountered an error: %v", err)
	}

}
