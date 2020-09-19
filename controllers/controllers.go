package controllers

import "github.com/gin-gonic/gin"

var _Router *gin.Engine

func StartControllers() {
	_Router = gin.Default()
	_StartUserController()
	_StartConversationController()

	if err := _Router.Run(":12400"); err != nil {
		panic(err)
	}
}
