package handlers 

import (
	"github.com/gin-gonic/gin"
)


func SignupApi (ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message" : "Hello World",
	},)

}