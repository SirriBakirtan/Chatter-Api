package controllers

import "github.com/gin-gonic/gin"

var _Router *gin.Engine

func StartControllers() {
	_Router = gin.Default()
	_StartUserController()
	_StartConversationController()
}
