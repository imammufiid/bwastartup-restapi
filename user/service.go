/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	repository Repository // dependency repository
}

func InstanceService(repository Repository) *service {
	return &service{repository: repository}
}

// implemented interface Service
func (s *service) RegisterUser(input RegisterInput) (User, error) {
	var user User = User{}

	// hasing password
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	// 1. mapping from user input to user entity
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Role = "user"
	user.PasswordHash = string(passHash)

	// 2. pass to repository
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	// checking email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// checked if user not found
	if user.ID == 0 {
		return user, errors.New("User not found on that email")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	// checking email
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return false, err
	}

	// check email found
	if user.ID == 0 {
		return true, nil
	}
	
	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	
}
