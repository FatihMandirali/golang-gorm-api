package dbconnect

import . "appword-api/models"

func InitialMigration() {
	connection := DbInit()
	defer CloseDatabase(connection)
	connection.AutoMigrate(&User{})
}
