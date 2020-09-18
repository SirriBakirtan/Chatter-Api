package main

import (
	"Chatter-Api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controllers.StartUserController(router)
	controllers.StartConversationController(router)
	if err := router.Run(":12400"); err != nil {
		panic(err)
	}
}
