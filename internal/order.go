// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     order.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 02:13:01
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"database/sql"
	"time"

	"github.com/ckf10000/gologger/v3/log"
)

// Order 结构体定义
type Order struct {
	PreOrderID        int       `json:"pre_order_id"`
	DepartureCity     string    `json:"departure_city"`
	ArriveCity        string    `json:"arrive_city"`
	DepartureTime     time.Time `json:"departure_time"`
	PreSaleAmount     string    `json:"pre_sale_amount"`
	Flight            string    `json:"flight"`
	Passenger         string    `json:"passenger"`
	AgeStage          string    `json:"age_stage"`
	CardType          string    `json:"card_type"`
	CardID            string    `json:"card_id"`
	InternalPhone     string    `json:"internal_phone"`
	PassengerPhone    string    `json:"passenger_phone"`
	CTripOrderID      string    `json:"ctrip_order_id"`
	PaymentAmount     string    `json:"payment_amount"`
	PaymentMethod     string    `json:"payment_method"`
	ItineraryID       string    `json:"itinerary_id"`
	DepartureCityName string    `json:"departure_city_name"`
	ArriveCityName    string    `json:"arrive_city_name"`
	ArriveTime        time.Time `json:"arrive_time"`
	CTripUsername     string    `json:"ctrip_username"`
	UserPass          string    `json:"user_pass"`
	OutPf             string    `json:"out_pf"`
	OutTicketAccount  string    `json:"out_ticket_account"`
	PayAccountType    string    `json:"pay_account_type"`
	PayAccount        string    `json:"pay_account"`
	Oper              string    `json:"oper"`
	PaymentTime       time.Time `json:"payment_time"`
}

// insertOrder 将订单插入到 MySQL
func insertOrder(db *sql.DB, order Order, log *log.FileLogger) error {
	stmt, err := db.Prepare(`
		INSERT INTO t_orders (
			pre_order_id, departure_city, arrive_city, departure_time,
			pre_sale_amount, flight, passenger, age_stage, card_type, card_id, 
			internal_phone, passenger_phone, ctrip_order_id, payment_amount, 
			payment_method, itinerary_id, departure_city_name, arrive_city_name, 
			arrive_time, ctrip_username, user_pass, out_pf, out_ticket_account, 
			pay_account_type, pay_account, oper, payment_time
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		order.PreOrderID, order.DepartureCity, order.ArriveCity, order.DepartureTime,
		order.PreSaleAmount, order.Flight, order.Passenger, order.AgeStage, order.CardType,
		order.CardID, order.InternalPhone, order.PassengerPhone, order.CTripOrderID,
		order.PaymentAmount, order.PaymentMethod, order.ItineraryID, order.DepartureCityName,
		order.ArriveCityName, order.ArriveTime, order.CTripUsername, order.UserPass, order.OutPf,
		order.OutTicketAccount, order.PayAccountType, order.PayAccount, order.Oper, order.PaymentTime,
	)
	if err != nil {
		return err
	}

	log.Info("Order inserted successfully: PreOrderID %d", order.PreOrderID)
	return nil
}
