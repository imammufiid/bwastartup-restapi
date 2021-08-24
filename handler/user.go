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
		c.JSON(http.StatusBadRequest, nil)
	}

	// 2. map to Register input
	// 3. pass to service
	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	// create response
	response := helper.ApiResponse(
		"Account has been registered",
		http.StatusOK,
		"success",
		user,
	)

	c.JSON(http.StatusCreated, response)
}