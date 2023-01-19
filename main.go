package main

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/db"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/server"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect DB", err)
	}

	server.Init()
}
