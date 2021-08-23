/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package handler

import "bwastartup/user"

type userHandler struct {
	userService user.Service
}

func InstanceUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}