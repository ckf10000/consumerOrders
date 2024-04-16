// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     message.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 02:15:58
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import "time"

// headers  内容
// {
// 	"app_id": "smartIssueTickets",
// 	"user_id": "ticket",
// 	"timestamp": 1712979477811,
// 	"message_id": "e9c8bcb8-56d4-4ce4-93c7-67a087f269de",
// 	"delivery_mode": 2,
// 	"content_encoding": "utf-8",
// 	"content_type": "application/json"
// }

type MessageHeaders struct {
	AppId           string    `json:"app_id"`
	UserId          string    `json:"user_id"`
	Timestamp       time.Time `json:"timestamp"`
	MessageId       string    `json:"message_id"`
	DeliveryMode    int       `json:"delivery_mode"`
	ContentEncoding string    `json:"content_encoding"`
	ContentType     string    `json:"content_type"`
}

type OrderMessage struct {
	PreOrderID        int    `json:"pre_order_id"`
	DepartureCity     string `json:"departure_city"`
	ArriveCity        string `json:"arrive_city"`
	DepartureTime     string `json:"departure_time"`
	PreSaleAmount     string `json:"pre_sale_amount"`
	Flight            string `json:"flight"`
	Passenger         string `json:"passenger"`
	AgeStage          string `json:"age_stage"`
	CardType          string `json:"card_type"`
	CardID            string `json:"card_id"`
	InternalPhone     string `json:"internal_phone"`
	PassengerPhone    string `json:"passenger_phone"`
	CTripOrderID      string `json:"ctrip_order_id"`
	PaymentAmount     string `json:"payment_amount"`
	PaymentMethod     string `json:"payment_method"`
	ItineraryID       string `json:"itinerary_id"`
	DepartureCityName string `json:"departure_city_name"`
	ArriveCityName    string `json:"arrive_city_name"`
	ArriveTime        string `json:"arrive_time"`
	CTripUsername     string `json:"ctrip_username"`
	UserPass          string `json:"user_pass"`
	OutPf             string `json:"out_pf"`
	OutTicketAccount  string `json:"out_ticket_account"`
	PayAccountType    string `json:"pay_account_type"`
	PayAccount        string `json:"pay_account"`
	Oper              string `json:"oper"`
}
