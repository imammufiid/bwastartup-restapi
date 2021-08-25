/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	Loginuser(input LoginInput) (User, error)
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

func (s *service) Loginuser(input LoginInput) (User, error) {
	
}
