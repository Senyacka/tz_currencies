package api

import (
	"io"
	"log"
	"net/http"
)

func GetCurrencyFromApiUSD(w http.ResponseWriter, r *http.Request) []byte {
	resp, err := http.Get("https://api.freecurrencyapi.com/v1/latest?apikey=fca_live_Ewku7vJ9godvShu9T3u5qT3nW1FA8LoJObX0jQ93&currencies=EUR%2CJPY%2CRUB%2CUSD")
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

func GetCurrencyFromApiRUB(w http.ResponseWriter, r *http.Request) []byte {
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

func GetCurrencyFromApiEUR(w http.ResponseWriter, r *http.Request) []byte {
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

func GetCurrencyFromApiJPY(w http.ResponseWriter, r *http.Request) []byte {
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