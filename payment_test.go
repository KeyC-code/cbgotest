package main

import (
	"testing"
)

func TestPaymentValidation(t *testing.T) {

	testCases := []struct {
		name          string
		amount        float64
		currency      string
		rate          float64
		allowed       bool
		expectedError string
	}{
		{
			name:          "100 USD, курс 100, разрешаем",
			amount:        100,
			currency:      "USD",
			rate:          100,
			allowed:       true,
			expectedError: "",
		},
		{
			name:          "200 USD, курс 150, - не разрешаем",
			amount:        200,
			currency:      "USD",
			rate:          150,
			allowed:       false,
			expectedError: "over the limit: 30000.00 RUB",
		},
		{
			name:          "100 EUR, курс 150, разрешаем (ровно в лимит)",
			amount:        100,
			currency:      "EUR",
			rate:          150,
			allowed:       true,
			expectedError: "",
		},
		{
			name:          "Нет такой валюты",
			amount:        1,
			currency:      "UIIAI",
			rate:          10,
			allowed:       false,
			expectedError: "Нет такой валюты, это мем UIIAI",
		},
		{
			name:          "Отрицательная сумма (а вдруг)",
			amount:        -100,
			currency:      "USD",
			rate:          100,
			allowed:       false,
			expectedError: "lower then 0: -10000.00 RUB",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rates := NewRates()

			if tc.rate > 0 {
				rates.SetRate(tc.currency, tc.rate)
			}

			paymentService := NewPaymentService(rates)

			request := PaymentRequest{
				Provider: "test_provider",
				Amount:   tc.amount,
				Currency: tc.currency,
			}

			result := paymentService.ProcessPayment(request)
			if result.Allowed != tc.allowed {
				t.Errorf(
					"Ожидалось разрешение: %v, Получили: %v",
					tc.allowed, result.Allowed,
				)
			}

			if tc.expectedError != "" {
				if result.Message == "" {
					t.Error("Должна быть ошибка, а ее нет...")
				} else if !(len(result.Message) >= len(tc.expectedError) && result.Message[len(result.Message)-len(tc.expectedError):] == tc.expectedError) {
					t.Errorf("Ожидали ошибку: '%s', получили: '%s'",
						tc.expectedError, result.Message)
				}
			} else {
				if result.Message == "" {
					t.Error("Нет информационного сообщения, а хотелось бы")
				}
			}

			if tc.rate > 0 && result.Allowed {
				expectedAmount := tc.amount * tc.rate
				if result.Amount != expectedAmount {
					t.Errorf("Ожидали сумму %.2f, получили %.2f",
						expectedAmount, result.Amount)
				}
			}

			t.Logf("Тест: %s", tc.name)
			t.Logf("Сумма: %.2f, курс: '%v', пропускаем?: %v, исключение:= '%s'",
				tc.amount, tc.currency, tc.allowed, tc.expectedError)
			t.Logf("Результат: allowed=%v, message='%s'",
				result.Allowed, result.Message)
		})
	}
}
