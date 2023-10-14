package main

import (
	"fmt"
	"log"
	"net/http"

	"tz_currency/internal/pkg/db"
)

type Currencies struct {
	USD float64
	RUB float64
	EUR float64
	JPY float64
}

func main() {

	http.HandleFunc("/currency", currencyHandler)
	http.HandleFunc("/api/health", healthCheckApi)
	http.HandleFunc("/health", healthCheck)

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func currencyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		body:= db.DbGetLastCurrencies(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	fmt.Println("health check")
	return
}

func healthCheckApi(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=USD")
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("health check api ERROR")
	}
	w.Write([]byte("OK"))
	fmt.Println("health check api")

	return
}

// TODO получение валютных пар из бд

// TODO валидация запросов

// TODO логирование запросов в бд

// TODO получение валюты по дате