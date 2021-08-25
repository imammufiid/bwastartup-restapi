/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func InstanceUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
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

	// change format reponse user
	formatUser := user.FormatUser(newUser, "token123")
	
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

	// mapping to format user response
	responseUser := user.FormatUser(loggedIn, "token123")
	response := helper.ApiResponse(
		"Successfuly loggedin",
		http.StatusOK,
		"success",
		responseUser,
	)
	c.JSON(http.StatusOK, response)
}