/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func InstanceRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(user User) (User, error) {
	// save to db with return error
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	
	// find by email query
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}