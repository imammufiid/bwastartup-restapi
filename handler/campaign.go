package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
	GetCampaigns(c *gin.Context)
	GetCampaign(c *gin.Context)
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

	responseCampaign := campaign.FormatCampaigns(campaigns)

	response := helper.ApiResponse(
		"List of Campaigns",
		http.StatusOK,
		"success",
		responseCampaign,
	)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// 1. Handle user input from url and mapping to struct input
	// 2. passing struct input to service and call repo
	// 3. get campaign by id 
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		// map to response
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Failed to bind URI",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaign(input)
	if err != nil {
		// map to response
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Failed to get campaign",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	message := fmt.Sprintf("Campaign detail with by ID %d", input.ID)
	response := helper.ApiResponse(
		message,
		http.StatusOK,
		"success",
		campaign.FormatDetailCampaign(campaignDetail),
	)
	c.JSON(http.StatusOK, response)
}