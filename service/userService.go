package service

import (
	"errors"
	"go-learn/common"
	"go-learn/model"
	"go-learn/request"
	"time"
)

func SignUp(req request.UserSignUp) (*model.User, error) {
	// 判断账号是否存在
	var users []model.User
	user := model.User{}
	user.Account = req.Account
	user.Password = req.Password
	user.SignUpTime = time.Now()
	err := model.DbInstance.Where("account = ?", req.Account).Limit(1).Find(&users).Error

	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 创建账号
	err = model.DbInstance.Create(&user).Error

	if err != nil {
		return nil, err
	}

	// 返回账号信息
	return &user, nil
}

func SignIn(req request.UserSignUp) (*model.User, string, error) {
	// 判断账号是否存在
	user := model.User{}
	err := model.DbInstance.Where("account = ? AND password = ?", req.Account, req.Password).Limit(1).Find(&user).Error

	if err != nil {
		return nil, "", err
	}

	// 更新登录时间
	user.LastSignInTime = time.Now()

	err = model.DbInstance.Save(&user).Error
	if err != nil {
		return nil, "", err
	}

	// 生成token
	t, err := common.GenerateToken(user.Account, user.Password)

	if err != nil {
		return nil, "", err
	}

	// 返回token和用户相关信息
	return &user, t, nil
}

func Get(req request.UserSignUp) (*model.User, error) {
	user := model.User{}
	err := model.DbInstance.Where("account = ? AND password = ?", req.Account, req.Password).Limit(1).Find(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
