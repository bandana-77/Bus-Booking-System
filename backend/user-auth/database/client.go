package database

import(
	"github.com/bandana-77/user-auth/models"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var Instance *gorm.DB
var dbError error
func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}