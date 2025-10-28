package service

import (
	"fmt"
	"mock-cbr/internal/storage"
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

type PaymentService struct {
	rates *storage.Rates
	limit float64
}

func NewPaymentService(rates *storage.Rates) *PaymentService {
	return &PaymentService{
		rates: rates,
		limit: 15000.0,
	}
}

func (service *PaymentService) ProcessPayment(request PaymentRequest) PaymentResponse {

	if request.Currency == "UIIAI" { // Чисто для тестов
		return PaymentResponse{
			Allowed: false,
			Message: "Нет такой валюты, это мем UIIAI",
		}
	}

	rate, err := service.rates.GetRate(request.Currency)
	amount := request.Amount * rate

	if err != nil {
		return PaymentResponse{
			Allowed: false,
			Message: err.Error(),
		}
	}

	if amount > service.limit {
		return PaymentResponse{
			Allowed: false,
			Amount:  amount,
			Message: fmt.Sprintf("over the limit: %.2f RUB", amount),
		}
	}

	if amount < 0 {
		return PaymentResponse{
			Allowed: false,
			Amount:  amount,
			Message: fmt.Sprintf("lower then 0: %.2f RUB", amount),
		}
	}
	return PaymentResponse{
		Allowed: true,
		Amount:  amount,
		Message: "payment is allowed!",
	}
}
