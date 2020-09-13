package controllers

import (
	"../models"
	"../services"
	"github.com/gin-gonic/gin"
)

var userRouter gin.RouterGroup

func StartUserController(router *gin.Engine) {
	userRouter = *router.Group("/user")
	startLoginEndpoint()
	startSignupEndpoint()
}

func startLoginEndpoint() {
	userRouter.POST("/login", func(ctx *gin.Context) {
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
	userRouter.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Signup": "Success",
		})
	})
}
