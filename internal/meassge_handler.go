// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     meassge_handler.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 04:11:47
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"database/sql"
	"encoding/json"

	"github.com/ckf10000/gologger/v3/log"
	"github.com/streadway/amqp"
)

// MessageHandler 处理消息的回调函数
type MessageHandler func(msg *amqp.Delivery, db *sql.DB, log *log.FileLogger) error

// handleMessage 处理消息的具体逻辑
func OrderMessageHandler(msg *amqp.Delivery, db *sql.DB, log *log.FileLogger) error {
	var orderMessage OrderMessage

	// 获取消息的 Properties 数据，内容如下：
	messageHeaders := ConvertHeadersToMessageHeaders(msg, log)
	err := json.Unmarshal(msg.Body, &orderMessage)
	if err != nil {
		log.Error("Error decoding message: %v", err)
		return err
	}

	log.Fatal("The received message is: %v", orderMessage)
	order, err := ConvertOrderMessageToOrder(orderMessage, messageHeaders)
	if err != nil {
		log.Error("message to order failed.")
		return err
	}

	// 插入订单到 MySQL
	err = insertOrder(db, order, log)
	if err != nil {
		log.Error("Error inserting order to MySQL: %v", err)
		return err
	}

	return nil
}
