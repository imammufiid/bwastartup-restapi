package router

import (
	h "bwastartup/handler"
	"bwastartup/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Router()
}

type router struct {
	authMiddleware  middleware.AuthMiddleware
	userHandler     h.UserHandler
	campaignHandler h.CampaignHandler
	trsHandler      h.TransactionHandler
}

func InstanceRouter(
	authMiddleware middleware.AuthMiddleware,
	userHandler h.UserHandler,
	campaignHandler h.CampaignHandler,
	trsHandler h.TransactionHandler,
) *router {
	return &router{
		userHandler:     userHandler,
		campaignHandler: campaignHandler,
		authMiddleware:  authMiddleware,
		trsHandler:      trsHandler,
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
	// transaction
	apiV1.GET("/campaigns/:id/transactions", r.authMiddleware.AuthMiddleware(), r.trsHandler.GetCampaignTransactions)
	apiV1.GET("/transactions", r.authMiddleware.AuthMiddleware(), r.trsHandler.GetUserTransactions)
	apiV1.POST("/transactions", r.authMiddleware.AuthMiddleware(), r.trsHandler.CreateTransaction)
	apiV1.POST("/transaction/notification", r.trsHandler.GetNotification)
	// running router
	router.Run()
}
