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
			expectedError: "over limit",
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
			expectedError: "unknown currency",
		},
		{
			name:          "Отрицательная сумма (а вдруг)",
			amount:        -100,
			currency:      "USD",
			rate:          100,
			allowed:       false,
			expectedError: "negatuive amount",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Тест: %s", tc.name)
			t.Logf("Сумма: %.2f, курс: '%v', пропускаем?: %v, исключение:= '%s'",
				tc.amount, tc.currency, tc.allowed, tc.expectedError)
			t.Skip("Нет реализаци пока")
		})
	}
}
