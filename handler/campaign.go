package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
	GetCampaigns(c *gin.Context)
}

type campaignHandler struct {
	campaignService campaign.Service
}

func InstanceCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// get query param user_id
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		// map to response
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Error to get campaigns",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse(
		"List of Campaigns",
		http.StatusOK,
		"success",
		campaigns,
	)
	c.JSON(http.StatusOK, response)
}