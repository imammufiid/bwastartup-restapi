/*
 * Created on Sun Aug 22 2021
 *
 *  Copyright (c) 2021 Imam Mufiid
 */

package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func Instance(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error // save to db with return error

	if err != nil {
		return user, err
	}
	return user, nil

}