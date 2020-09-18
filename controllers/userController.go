package controllers

import (
	"Chatter-Api/models"
	"Chatter-Api/services"
	"github.com/gin-gonic/gin"
)

var _UserRouter *gin.RouterGroup

func _StartUserController() {
	_UserRouter = _Router.Group("/user")
	startLoginEndpoint()
	startSignupEndpoint()
}

func init() {

}

func startLoginEndpoint() {
	_UserRouter.POST("/login", func(ctx *gin.Context) {
		username := ctx.Request.Header.Get("username")
		hashedPassword := ctx.Request.Header.Get("hashedPassword")
		if username == "" || hashedPassword == "" {
			ctx.JSON(400, gin.H{
				"Status": "Missing credentials!",
			})
			return
		}
		if services.UserService.Login(models.User{
			Username:       username,
			HashedPassword: hashedPassword,
		}) {
			ctx.JSON(200, gin.H{
				"Status": "Success!",
			})
			return
		} else {
			ctx.JSON(401, gin.H{
				"Status": "Invalid credentials!",
			})
			return
		}
	})
}

func startSignupEndpoint() {
	_UserRouter.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Signup": "Success",
		})
	})
}
