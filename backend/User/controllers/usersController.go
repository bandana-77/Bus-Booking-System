package controllers

import "github.com/gin-gonic/gin"

func Signup(c *gin.Context){
	//Get the emai/pass off req body

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