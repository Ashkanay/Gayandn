package configration

import (
	"fmt"
	"gayandn/model"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to db
func ConnectDB() {
	var err error

	//Connection to Prodection
	//dsn := os.Getenv("DB_PRO_URL")

	//Connection to Devlopment
	//dsn := os.Getenv("DB_DEV_URL")

	dsn := "host=localhost user=postgres password=barez dbname=Gayandn port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed connection *************************************************")
		panic(err)
	}

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
	DB = db
	fmt.Println(DB)

}
