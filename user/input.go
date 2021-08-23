/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package user

// struct for mapping input from user
type RegisterInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}
