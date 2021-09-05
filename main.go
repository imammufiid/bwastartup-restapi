package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/router"
	"bwastartup/user"
	"bwastartup/campaign"
	"bwastartup/transaction"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connection to database
	dsn := "root:@tcp(127.0.0.1:3306)/learn_bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// init user repository
	userRepo := user.InstanceRepository(db)
	campaignRepo := campaign.InstanceRepository(db)
	trsRepo := transaction.InstanceRepository(db)

	// init service
	userService := user.InstanceService(userRepo)
	authService := auth.InstanceService()
	campaignService := campaign.InstanceService(campaignRepo)
	trsService := transaction.InstanceService(trsRepo)

	// init handler
	authMiddleware := middleware.InstanceAuthMiddleware(userService, authService)
	userHandler := handler.InstanceUserHandler(userService, authService)
	campaignHandler := handler.InstanceCampaignHandler(campaignService)
	trsHandler := handler.InstanceTransactionHandler(trsService)

	// router
	router := router.InstanceRouter(authMiddleware, userHandler, campaignHandler, trsHandler)
	router.Router()

}
