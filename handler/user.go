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
	"github.com/go-playground/validator/v10"
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

		// setup validation
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		// map to response
		errorMessage := gin.H{"errors": errors}

		// create error handling response
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// 2. map to Register input
	// 3. pass to service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		// create error handling response
		response := helper.ApiResponse(
			"Register account failed",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
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