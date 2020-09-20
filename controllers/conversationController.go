package controllers

import (
	"Chatter-Api/services"
	"github.com/gin-gonic/gin"
)

var _ConversationRouter *gin.RouterGroup

func init() {

}

func _StartConversationController() {
	_ConversationRouter = _Router.Group("conversation")
	startGetMessagesEndpoint()
}

func startGetMessagesEndpoint() {
	_ConversationRouter.GET("", func(ctx *gin.Context) {
		userId := ctx.Request.Header.Get("userId")
		conversation, err := services.ConversationService.GetConversationsOfUser(userId)
		if err != nil {
			ctx.JSON(200, gin.H{
				"Status": err.Error(),
			})
		}
		ctx.JSON(200, conversation)

	})
}
