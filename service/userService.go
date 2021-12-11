package service

import (
	"go-learn/model"
)

func CreateUser(name string) (*model.User, error) {
	user := model.User{}
	user.Name = name
	err := model.DbInstance.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
