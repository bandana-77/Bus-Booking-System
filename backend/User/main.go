package main

import (
	// "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	// initializers "easytripz-main/backend/User/initializers"

// 	"gorm.io/gorm"
 
//   "gorm.io/driver/sqlserver"
	// "github.com/golang-jwt/jwt/v4"
)

func init(){
  LoadEnvVariables()
}
var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"Ashish:AshishDB@tcp(easytripz.cscqq6zfyvxt.ap-south-1.rds.amazonaws.com:3306)/busbookingsystemdb")
	if err != nil {
		panic(err.Error())
	}
}

func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

}

func main(){
	InitDB()
	LoadEnvVariables()
	

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}


//Signup

func Signup(c *gin.Context){
	//Get the email/pass off req body

	var body struct{
		email string
		password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body"
		})
		return

	}

}

//Hash Password
hash,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Failed to hash password"
	})

	return
}

//create user
User_details.Select("email", "password").Create(&body)
