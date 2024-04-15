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

import "time"

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
	}

	return order, nil
}
