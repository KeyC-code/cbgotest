package handler

import (
	"encoding/xml"
	"fmt"
	"mock-cbr/internal/models"
	"mock-cbr/internal/storage"
	"net/http"
	"strings"
	"time"
)

type CBRHandler struct {
	storage *storage.Rates
}

func NewCBRHandler(storage *storage.Rates) *CBRHandler {
	return &CBRHandler{
		storage: storage,
	}
}

func (h *CBRHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("ЦБ Мок: получен запрос %s %s\n", r.Method, r.URL.String())
	if r.Method != "GET" {
		http.Error(w, "Метод не доступен", http.StatusMethodNotAllowed)
		return
	}

	dateReq := r.URL.Query().Get("date_req")
	if dateReq != "" {
		fmt.Printf("ЦБ Мок: запрошена дата %s\n", dateReq)
	}

	rates := h.storage.GetAllRates()

	response := h.buildCBRResponse(rates)

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	encoder := xml.NewEncoder(w)
	encoder.Indent("", " ")
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Не удалось расшифровать XML", http.StatusInternalServerError)
		return
	}

	fmt.Printf("ЦБ Мок: отправлен ответ по валютам: %d\n", len(response.Valutes))

}

func (h *CBRHandler) buildCBRResponse(rates map[string]float64) models.ValCurs {
	valCurs := models.ValCurs{
		Date: time.Now().Format("02.01.2006"),
		Name: "Foreign Currency Market",
	}

	for currencyCode, rate := range rates {
		if mapping, exists := models.CurrencyMapping[currencyCode]; exists {
			valute := models.Valute{
				ID:       mapping.ID,
				NumCode:  mapping.NumCode,
				CharCode: currencyCode,
				Nominal:  1,
				Name:     mapping.Name,
				Value:    h.formatRate(rate),
			}
			valCurs.Valutes = append(valCurs.Valutes, valute)
		} else {
			valute := models.Valute{
				ID:       "R00000",
				NumCode:  "000",
				CharCode: currencyCode,
				Nominal:  1,
				Name:     currencyCode,
				Value:    h.formatRate(rate),
			}
			valCurs.Valutes = append(valCurs.Valutes, valute)
		}
	}
	return valCurs
}

func (h *CBRHandler) formatRate(rate float64) string {
	formatted := fmt.Sprintf("%.4f", rate)
	return strings.Replace(formatted, ".", ",", 1)
}
