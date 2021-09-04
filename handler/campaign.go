package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler interface {
	GetCampaigns(c *gin.Context)
	GetCampaign(c *gin.Context)
	CreateCampaign(c *gin.Context)
	UpdateCampaign(c *gin.Context)
	UploadImage(c *gin.Context)
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

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get current user from context
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// call service
	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse(
		"Success create campaign",
		http.StatusOK,
		"success",
		campaign.FormatCampaign(newCampaign),
	)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// get input ID from URI
	var inputID campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&inputID)
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

	// get data from body json
	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get current user by jwt
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.campaignService.UpdateCampaign(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse(
		"Success updated campaign",
		http.StatusOK,
		"success",
		campaign.FormatCampaign(updatedCampaign),
	)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImage

	// get input from form data
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	// get form data file
	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// set path
	path := fmt.Sprintf("assets/images/campaign/%d-%s_%s", currentUser.ID, file.Filename, helper.TimeNowMilli())

	// save file in path
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to save upload campaign image into path", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.campaignService.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		// create error handling response
		response := helper.ApiResponse(
			"Failed to upload campaign image into database",
			http.StatusUnprocessableEntity,
			"error",
			data,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	// create error handling response
	response := helper.ApiResponse(
		"Campaign image successfuly uploaded",
		http.StatusOK,
		"success",
		data,
	)
	c.JSON(http.StatusOK, response)
}