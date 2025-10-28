package main

import (
	"fmt"
	"log"
	"mock-cbr/internal/handler"
	"mock-cbr/internal/storage"
	"net/http"
)

func main() {
	fmt.Println("Запускаем Мок ЦБ")

	ratesStorage := storage.NewRates()

	cbrHandler := handler.NewCBRHandler(ratesStorage)

	mux := http.NewServeMux()
	mux.Handle("/scripts/XML_daily.asp", cbrHandler)
	mux.Handle("/cbr", cbrHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Мок сервис ЦБ запущен на localhost:8080")
	log.Printf("Доступные эндпоинты:")
	log.Printf("  - http://localhost:8080/scripts/XML_daily.asp")
	log.Printf("  - http://localhost:8080/cbr")
	log.Printf("Пример: http://localhost:8080/cbr?date_req=02/03/2002")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
