package main

import(
	"github.com/bandana-77/user-auth/database"
	"github.com/bandana-77/user-auth/middlewares"
	"github.com/bandana-77/user-auth/controllers"
	"github.com/gin-gonic/gin"
)
func main(){
	// Initialize Database
	database.Connect("Ashish:AshishDB@tcp(easytripz.cscqq6zfyvxt.ap-south-1.rds.amazonaws.com:3306)/busbookingsystemdb?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}