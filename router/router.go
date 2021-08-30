package router

import (
	"bwastartup/handler"
	"bwastartup/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Router()
}

type router struct {
	userHandler    handler.UserHandler
	authMiddleware middleware.AuthMiddleware
}

func InstanceRouter(userHandler handler.UserHandler, authMiddleware middleware.AuthMiddleware) *router {
	return &router{
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *router) Router() {
	// router
	router := gin.Default()
	// api versioning
	apiV1 := router.Group("/api/v1")
	// route
	apiV1.POST("/users", r.userHandler.RegisterUser)
	apiV1.POST("/sessions", r.userHandler.Login)
	apiV1.POST("/email_checker", r.userHandler.CheckEmailIsAvailable)
	apiV1.POST("/avatars", r.authMiddleware.AuthMiddleware(), r.userHandler.UploadAvatar)
	// running router
	router.Run()
}
