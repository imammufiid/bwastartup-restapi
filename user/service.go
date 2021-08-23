/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package user

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
}

type service struct {
	repository Repository // dependency repository
}

func InstanceService(repository Repository) *service  {
	return &service{repository: repository}
}

// implemented interface Service
func (s *service) RegisterUser(input RegisterInput) (User, error) {
	
}
