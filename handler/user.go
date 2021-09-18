/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"bwastartup/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	CheckEmailIsAvailable(c *gin.Context)
	UploadAvatar(c *gin.Context)
	UserTest(c *gin.Context)
}

type userHandler struct {
	userService user.Service
	jwtService auth.Service
}

func InstanceUserHandler(userService user.Service, jwtService auth.Service) *userHandler {
	return &userHandler{userService: userService, jwtService: jwtService}
}

func (h *userHandler) RegisterUser(c *gin.Context)  {
	// 1. get input from user to JSON
	var input user.RegisterInput
	// bind to JSON
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		// map to response
		errorMessage := gin.H{"errors": errors}

		// create error handling response
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// 2. map to Register input
	// 3. pass to service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		// create error handling response
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusUnprocessableEntity,
			"error",
			nil,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// generate token
	token, err := h.jwtService.GenerateToken(newUser.ID)
	if err != nil {
		// map to response
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Failed Generate token",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	// change format reponse user
	formatUser := user.FormatUser(newUser, token)
	
	// create response
	response := helper.ApiResponse(
		"Account has been registered",
		http.StatusOK,
		"success",
		formatUser,
	)

	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	// bind input to JSON
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		// map to response
		errorMessage := gin.H{"errors": errors}
		// create error handling response
		response := helper.ApiResponse(
			"Login Failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// logged in process
	loggedIn, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Login Failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// generate token
	token, err := h.jwtService.GenerateToken(loggedIn.ID)
	if err != nil {
		// map to response
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Failed Generate token",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// mapping to format user response
	responseUser := user.FormatUser(loggedIn, token)
	response := helper.ApiResponse(
		"Successfuly loggedin",
		http.StatusOK,
		"success",
		responseUser,
	)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailIsAvailable(c *gin.Context) {
	// get input from user
	var input user.CheckEmailInput

	// bind to JSON
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		// map to response
		errorMessage := gin.H{"errors": errors}
		// create error handling response
		response := helper.ApiResponse(
			"Email checking failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// check email from service
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		// create error handling response
		response := helper.ApiResponse(
			"Login Failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// map data is available
	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email is available"

	if !isEmailAvailable {
		metaMessage = "Email address has been registered"
	}

	response := helper.ApiResponse(
		metaMessage,
		http.StatusOK,
		"success",
		data,
	)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// get file input user
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		// create error handling response
		response := helper.ApiResponse(
			"Failed to upload avatar image",
			http.StatusUnprocessableEntity,
			"error",
			data,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User) // get context with key currentUser
	userID := currentUser.ID

	// asign path file save
	path := fmt.Sprintf("assets/images/avatar/%d-%s_%s", userID, file.Filename, helper.TimeNowMilli())

	// save file to path
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		// create error handling response
		response := helper.ApiResponse(
			"Failed to save avatar image into path",
			http.StatusUnprocessableEntity,
			"error",
			data,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		// create error handling response
		response := helper.ApiResponse(
			"Failed to save avatar image into database",
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
		"Avatar successfuly uploaded",
		http.StatusOK,
		"success",
		data,
	)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UserTest(c *gin.Context) {
	response := helper.ApiResponse(
		"Testing",
		http.StatusOK,
		"error",
		nil,
	)
	c.JSON(http.StatusOK, response)
}