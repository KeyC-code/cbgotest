package main

import (
	"fmt"
)

type PaymentRequest struct { // запрос на оплату
	Provider string
	Amount   float64
	Currency string
}

type PaymentResponse struct { // ответ на запрос
	Allowed bool
	Amount  float64
	Message string
}

func ProcessPayment(request PaymentRequest) PaymentResponse {

	if request.Currency == "UIIAI" { // Чисто для тестов
		return PaymentResponse{
			Allowed: false,
			Message: "Нет такой валюты, это мем UIIAI",
		}
	}

	rates := map[string]float64{ // Тоже чисто для тестов
		"USD": 90,
		"EUR": 150,
	}

	rate, exists := rates[request.Currency]
	amount := request.Amount * rate

	if !exists {
		return PaymentResponse{
			Allowed: false,
			Message: fmt.Sprintf("currency not supported: %s", request.Currency),
		}
	}

	if amount > 15000 {
		return PaymentResponse{
			Allowed: false,
			Message: fmt.Sprintf("over the limit: %.2f RUB", amount),
		}
	}

	if amount < 0 {
		return PaymentResponse{
			Allowed: false,
			Message: fmt.Sprintf("lower then 0: %.2f RUB", amount),
		}
	}
	return PaymentResponse{
		Allowed: true,
		Message: "payment is allowed!",
	}
}
