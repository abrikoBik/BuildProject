package main

import (
	"database/sql"
	"log"
	"mp1/db"
	"mp1/util"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	util.Check(godotenv.Load())
	var (
		database *sql.DB
		err      error
	)
	for {
		database, err = db.Open("localhost", 5432, "main")
		if err == nil {
			break
		}
		log.Println(err)
		log.Println("retrying in 5 seconds")
		time.Sleep(time.Second * 5)
	}
	defer database.Close()
	log.Println("connected to database")
}
