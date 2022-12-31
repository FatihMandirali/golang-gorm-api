package dbconnect

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //bu kütüphaneyi kurmamız gerekmektedir...
	"log"
)

func DbInit() *gorm.DB {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//connString1 := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"
	connection, err := gorm.Open("postgres", connString)
	//db.SetMaxIdleConns(5)
	//db.SetMaxOpenConns(5)
	//db.SetConnMaxIdleTime(1*time.Second)
	//db.SetConnMaxLifetime(1*time.Second)
	if err != nil {
		log.Fatalln("db ye bağlanamadı", err.Error())
	}
	return connection
}

// closes database connection
func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
