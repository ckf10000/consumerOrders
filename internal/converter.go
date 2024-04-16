// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     converter.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 02:17:03
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"encoding/json"
	"time"

	"github.com/ckf10000/gologger/v3/log"
	"github.com/streadway/amqp"
)

func ConvertUtcToTime(utc int) time.Time {
	// 以毫秒表示的时间戳
	milliseconds := int64(utc)

	// 将毫秒转换为秒和纳秒
	seconds := milliseconds / 1000                 // 毫秒转换为秒
	nanoseconds := (milliseconds % 1000) * 1000000 // 毫秒的余数转换为纳秒

	// 使用 time.Unix() 将秒数和纳秒数转换为 time.Time 对象
	t := time.Unix(seconds, nanoseconds)
	return t
}

func ConvertHeadersToMessageHeaders(msg *amqp.Delivery, log *log.FileLogger) MessageHeaders {
	// 获取消息的 Headers
	headers := msg.Headers
	// 创建 MessageHeaders 结构体实例
	messageHeaders := MessageHeaders{}
	if headers == nil {
		log.Error("Headers are nil")
		return MessageHeaders{}
	}

	// 将 Headers 转换为 map[string]interface{}
	headerMap := make(map[string]interface{})
	for k, v := range headers {
		if k == "timestamp" {
			headerMap[k] = ConvertUtcToTime(v.(int))
		} else {
			headerMap[k] = v
		}
	}

	// 将 headerMap 转换为 JSON 字符串
	jsonData, err := json.Marshal(headerMap)
	if err != nil {
		log.Error("Failed to marshal header map to JSON: %v", err)
	}

	// 将 JSON 字符串解析到 MessageHeaders 结构体中
	err = json.Unmarshal(jsonData, &messageHeaders)
	if err != nil {
		log.Error("Failed to unmarshal JSON to MessageHeaders: %v", err)
	}
	return messageHeaders
}

// ConvertOrderMessageToOrder converts an OrderMessage to an Order
func ConvertOrderMessageToOrder(msg OrderMessage, messageHeaders MessageHeaders) (Order, error) {
	departureTime, err := time.Parse("2006-01-02 15:04:05", msg.DepartureTime)
	if err != nil {
		return Order{}, err
	}

	arriveTime, err := time.Parse("2006-01-02 15:04:05", msg.ArriveTime)
	if err != nil {
		return Order{}, err
	}

	messageTime := messageHeaders.Timestamp
	if messageTime == (time.Time{}) {
		// 给 messageTime 赋值为当前时间
		messageTime = time.Now()
	}
	order := Order{
		PreOrderID:        msg.PreOrderID,
		DepartureCity:     msg.DepartureCity,
		ArriveCity:        msg.ArriveCity,
		DepartureTime:     departureTime,
		PreSaleAmount:     msg.PreSaleAmount,
		Flight:            msg.Flight,
		Passenger:         msg.Passenger,
		AgeStage:          msg.AgeStage,
		CardType:          msg.CardType,
		CardID:            msg.CardID,
		InternalPhone:     msg.InternalPhone,
		PassengerPhone:    msg.PassengerPhone,
		CTripOrderID:      msg.CTripOrderID,
		PaymentAmount:     msg.PaymentAmount,
		PaymentMethod:     msg.PaymentMethod,
		ItineraryID:       msg.ItineraryID,
		DepartureCityName: msg.DepartureCityName,
		ArriveCityName:    msg.ArriveCityName,
		ArriveTime:        arriveTime,
		CTripUsername:     msg.CTripUsername,
		UserPass:          msg.UserPass,
		OutPf:             msg.OutPf,
		OutTicketAccount:  msg.OutTicketAccount,
		PayAccountType:    msg.PayAccountType,
		PayAccount:        msg.PayAccount,
		Oper:              msg.Oper,
		CreateTime:        messageTime,
		UpdateTime:        messageTime,
	}

	return order, nil
}
