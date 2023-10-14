package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// _ "github.com/go-sql-driver/mysql"
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
		body := getCurrencyFromApiUSD(w, r)
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
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=EUR%2CJPY%2CRUB")
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("health check api ERROR")
	}
	fmt.Println("health check api")

	return
}

func getCurrencyFromApiUSD(w http.ResponseWriter, r *http.Request) []byte {
	var m Currencies
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=EUR%2CJPY%2CRUB%2CUSD")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	c := string(body[8 : len(body)-1]) //TODO переделать костыль
	a := json.Unmarshal([]byte(c), &m)
	if a != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	fmt.Printf("USD: %v\n", m.USD)
	fmt.Printf("RUB: %v\n", m.RUB)
	fmt.Printf("EUR: %v\n", m.EUR)
	fmt.Printf("JPY: %v\n", m.JPY)

	return body
}

func getCurrencyFromApiRUB(w http.ResponseWriter, r *http.Request) []byte {
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=EUR%2CUSD%2CJPY%2CRUB&base_currency=RUB")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getCurrencyFromApiEUR(w http.ResponseWriter, r *http.Request) []byte {
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=EUR%2CUSD%2CJPY%2CRUB&base_currency=EUR")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getCurrencyFromApiJPY(w http.ResponseWriter, r *http.Request) []byte {
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=EUR%2CUSD%2CJPY%2CRUB&base_currency=JPY")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}


// TODO отправка данных в бд, получение

// TODO получение валютных пар из бд

// TODO обновление валюты по крону 

// TODO валидация запросов

// TODO догирование запросов в бд

// TODO получение валюты по дате