package handler

import (
	"encoding/xml"
	"mock-cbr/internal/models"
	"mock-cbr/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCBRHandler(t *testing.T) {
	storage := storage.NewRates()
	storage.SetRate("USD", 30.5)
	storage.SetRate("EUR", 35.2)

	test_handler := NewCBRHandler(storage)

	request := httptest.NewRequest("GET", "/cbr?date_req=02/03/2002", nil)
	w := httptest.NewRecorder()

	test_handler.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Errorf("Статус: %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/xml; charset=utf-8" {
		t.Errorf("Получен Content-Type: %s", contentType)
	}

	var response models.ValCurs
	if err := xml.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Ошибка парсинга XML: %v", err)
	}

	if len(response.Valutes) < 2 {
		t.Errorf("Ожидалось минимум  2 валюты, получили %d", len(response.Valutes))
	}

	foundUSD := false
	for _, valute := range response.Valutes {
		if valute.CharCode == "USD" {
			foundUSD = true
			if valute.Value != "30,5000" {
				t.Errorf("Ожидался курс доллара 30,5000, получили: %s", valute.Value)
			}
			break
		}
	}

	if !foundUSD {
		t.Error("USD нет в ответе")
	}

	t.Logf("Успех! Получен ответ с %d валютами", len(response.Valutes))
	for _, valute := range response.Valutes {
		t.Logf(" %s: %s (%s)", valute.CharCode, valute.Value, valute.Name)
	}
}
