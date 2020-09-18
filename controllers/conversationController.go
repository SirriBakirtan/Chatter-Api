package controllers

import (
	"Chatter-Api/models"
	"Chatter-Api/services"
	"github.com/gin-gonic/gin"
)

var conversationRouter *gin.RouterGroup

func init() {

}

func StartConversationController(router *gin.Engine) {
	conversationRouter = router.Group("conversation")
	startGetMessagesEndpoint()
}

func startGetMessagesEndpoint() {
	conversationRouter.GET("", func(ctx *gin.Context) {
		userId := ctx.Request.Header.Get("userId")
		conversationId := ctx.Request.Header.Get("conversationId")
		conversation, err := services.ConversationService.GetConversation(userId, models.Conversation{
			ID: conversationId,
		})
		if err != nil {
			var errorMessage string
			switch err.Error() {
			case "failed_to_get_messages":
				errorMessage = "Failed to get messages!"
			case "user_not_included":
				errorMessage = "Looks like you don't belong to this chat!"
			default:
				errorMessage = "Something happened!"
			}
			ctx.JSON(400, gin.H{
				"Status": errorMessage,
			})
			return
		}
		ctx.JSON(200, conversation)

	})
}
