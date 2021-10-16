package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/therealbahodir/rest-chat/handlers"
	"os"
	"log"
)


func main () {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("APP_HTTP_PORT")

	httpRouter := gin.Default()

	httpRouter.POST("/signup", handlers.SignupApi)

	httpRouter.Run(port)

}