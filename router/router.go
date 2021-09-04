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
	userHandler     handler.UserHandler
	campaignHandler handler.CampaignHandler
	authMiddleware  middleware.AuthMiddleware
}

func InstanceRouter(
	userHandler handler.UserHandler,
	campaignHandler handler.CampaignHandler,
	authMiddleware middleware.AuthMiddleware,
) *router {
	return &router{
		userHandler:     userHandler,
		campaignHandler: campaignHandler,
		authMiddleware:  authMiddleware,
	}
}

func (r *router) Router() {
	// router
	router := gin.Default()
	// setting static route for file
	router.Static("/image/avatar", "./assets/images/avatar")
	// api versioning
	apiV1 := router.Group("/api/v1")
	// route
	apiV1.POST("/users", r.userHandler.RegisterUser)
	apiV1.POST("/sessions", r.userHandler.Login)
	apiV1.POST("/email_checker", r.userHandler.CheckEmailIsAvailable)
	apiV1.POST("/avatars", r.authMiddleware.AuthMiddleware(), r.userHandler.UploadAvatar)
	// campaign
	apiV1.GET("/campaigns", r.campaignHandler.GetCampaigns)
	apiV1.GET("/campaigns/:id", r.campaignHandler.GetCampaign)
	apiV1.POST("/campaigns", r.authMiddleware.AuthMiddleware(), r.campaignHandler.CreateCampaign)
	apiV1.PUT("/campaigns/:id", r.authMiddleware.AuthMiddleware(), r.campaignHandler.UpdateCampaign)
	apiV1.POST("/campaign-images", r.authMiddleware.AuthMiddleware(), r.campaignHandler.UploadImage)
	// running router
	router.Run()
}
