package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/middleware"
	"bwastartup/payment"
	"bwastartup/router"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connection to database
	dbName := helper.GetENV("DB_NAME", "golang-db-name")
	dbUsername := helper.GetENV("DB_USERNAME", "root")

	dsn := fmt.Sprintf("%s:@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbName)
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
	paymentService := payment.InstanceService()
	trsService := transaction.InstanceService(trsRepo, campaignRepo, paymentService)

	// init handler
	authMiddleware := middleware.InstanceAuthMiddleware(userService, authService)
	userHandler := handler.InstanceUserHandler(userService, authService)
	campaignHandler := handler.InstanceCampaignHandler(campaignService)
	trsHandler := handler.InstanceTransactionHandler(trsService)

	// router
	router := router.InstanceRouter(authMiddleware, userHandler, campaignHandler, trsHandler)
	router.Router()

}
