package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"tz_currency/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// DbInsertCurrencies - Вставляет курсы валют в базу данных. type = 1:USD, 2:RUB, 3:EUR, 4:JPY. date = NOW() и заполняется автоматически.

func DbInsertCurrencies(c models.Currencies) error {
	db, err := sql.Open("mysql", "root:MyPass123!@tcp(127.0.0.1:3306)/tz_currency")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Вставка данных
	insert, err := db.Query("INSERT INTO currencies (type ,date , usd, rub, eur, jpy) VALUES (1, NOW(), ?, ?, ?, ?)", c.USD, c.RUB, c.EUR, c.JPY)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	return nil
}

// DbGetCurrencies - Возвращает последние курсы валют по отношению к переданному type = 1:USD, 2:RUB, 3:EUR, 4:JPY.

func DbGetLastCurrencies(t int) []byte {
	db, err := sql.Open("mysql", "root:MyPass123!@tcp(127.0.0.1:3306)/tz_currency")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Выборка данных
	rows, err := db.Query("SELECT usd, rub, eur, jpy FROM currencies WHERE date = (SELECT MAX(date) FROM currencies) and type = ?", t)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var c models.Currencies
	for rows.Next() {
		err := rows.Scan(&c.USD, &c.RUB, &c.EUR, &c.JPY)
		if err != nil {
			log.Fatal(err)
		}
	}
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// DbGetSpecificDayCurrencies - Возвращает курсы валют по отношению к переданному type = 1:USD, 2:RUB, 3:EUR, 4:JPY в определённый день

// func DbGetSpecificDayCurrencies(t int, d string) []byte {
// 	db, err := sql.Open("mysql", "root:MyPass123!@tcp(127.0.0.1:3306)/tz_currency")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	// Выборка данных
// 	rows, err := db.Query("SELECT usd, rub, eur, jpy FROM currencies WHERE date = ? and type = ?", b, t)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var c models.Currencies
// 	for rows.Next() {
// 		err := rows.Scan(&c.USD, &c.RUB, &c.EUR, &c.JPY)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	b, err := json.Marshal(c)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return b
// }

// func DBGetPairCurrencies(t, b int) []byte {

// 	db, err := sql.Open("mysql", "root:MyPass123!@tcp(127.0.0.1:3306)/tz_currency")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	// Выборка данных
// 	rows, err := db.Query("SELECT usd, rub, eur, jpy FROM currencies WHERE date = (SELECT MAX(date) FROM currencies) and type = ?", t)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var c models.Currencies
// 	for rows.Next() {
// 		err := rows.Scan(&c.USD, &c.RUB)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	return
// }

func Cron() {
	for {
		DbGetLastCurrencies(1)
		DbGetLastCurrencies(2)
		DbGetLastCurrencies(3)
		DbGetLastCurrencies(4)
		time.Sleep(time.Hour * 12)
	}
}
