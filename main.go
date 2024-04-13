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
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ckf10000/gologger/v2/log"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
	"github.com/streadway/amqp"
)

type OrderMessage struct {
	PreOrderID        int    `json:"pre_order_id"`
	DepartureCity     string `json:"departure_city"`
	ArriveCity        string `json:"arrive_city"`
	DepartureTime     string `json:"departure_time"`
	PreSaleAmount     string `json:"pre_sale_amount"`
	Flight            string `json:"flight"`
	Passenger         string `json:"passenger"`
	AgeStage          string `json:"age_stage"`
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
}

// ConvertOrderMessageToOrder converts an OrderMessage to an Order
func ConvertOrderMessageToOrder(msg OrderMessage) (Order, error) {
	departureTime, err := time.Parse("2006-01-02 15:04:05", msg.DepartureTime)
	if err != nil {
		return Order{}, err
	}

	arriveTime, err := time.Parse("2006-01-02 15:04:05", msg.ArriveTime)
	if err != nil {
		return Order{}, err
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
	}

	return order, nil
}

func main() {
	// 获取可执行文件所在目录的路径
	exeDir := log.GetExecuteFilePath()
	if exeDir == "" {
		return
	}

	// projectPath := "./"
	log := log.NewLogger("debug", exeDir, "comsumer.log", 50*1024*1024, true, true, true)
	// RabbitMQ 连接信息
	amqpURI := "amqp://ticket:Admin%40123@192.168.3.232:5672/smartIssueTickets"
	exchangeName := "amq.fanout"
	exchangeType := "fanout"
	queueName := "order.flight.ctrip"
	routingKey := "order.flight.ctrip"

	// MySQL 连接信息
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "Admin@123", "192.168.3.232", "3306", "cloudBeanOrder")

	// 连接 RabbitMQ
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Error("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	// 声明 Exchange
	err = channel.ExchangeDeclare(
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
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to declare a queue: %v", err)
	}

	// 绑定 Queue 到 Exchange
	err = channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Error("Failed to bind queue to exchange: %v", err)
	}

	// 连接 MySQL
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Error("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

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
	}

	for msg := range msgs {
		var orderMessage OrderMessage
		err := json.Unmarshal(msg.Body, &orderMessage)
		if err != nil {
			log.Error("Error decoding message: %v", err)
			continue
		}
		log.Fatal("The received message is: %v", orderMessage)
		order, err := ConvertOrderMessageToOrder(orderMessage)
		if err != nil {
			log.Error("message to order failed.")
			continue
		}
		// 插入订单到 MySQL
		err = insertOrder(db, order, log)
		if err != nil {
			log.Error("Error inserting order to MySQL: %v", err)
			continue
		}
	}
}

// insertOrder 将订单插入到 MySQL
func insertOrder(db *sql.DB, order Order, log *log.FileLogger) error {
	stmt, err := db.Prepare(`
		INSERT INTO t_orders (
			pre_order_id, departure_city, arrive_city, departure_time,
			pre_sale_amount, flight, passenger, age_stage, card_id, internal_phone,
			passenger_phone, ctrip_order_id, payment_amount, payment_method, 
			itinerary_id, departure_city_name, arrive_city_name, arrive_time, 
			ctrip_username, user_pass, out_pf, out_ticket_account, pay_account_type, 
			pay_account, oper
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		order.PreOrderID, order.DepartureCity, order.ArriveCity, order.DepartureTime,
		order.PreSaleAmount, order.Flight, order.Passenger, order.AgeStage, order.CardID,
		order.InternalPhone, order.PassengerPhone, order.CTripOrderID, order.PaymentAmount,
		order.PaymentMethod, order.ItineraryID, order.DepartureCityName, order.ArriveCityName,
		order.ArriveTime, order.CTripUsername, order.UserPass, order.OutPf, order.OutTicketAccount,
		order.PayAccountType, order.PayAccount, order.Oper,
	)
	if err != nil {
		return err
	}

	log.Info("Order inserted successfully: PreOrderID %d", order.PreOrderID)
	return nil
}
