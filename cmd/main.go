package main

import "log"
import "appword-api/dbconnect"
import "appword-api/handlers"

func main() {
	log.Println("Service starting")

	dbconnect.DbInit()
	dbconnect.InitialMigration()
	handlers.Run()

	log.Println("service ending")
}
