package response

import (
	"go-learn/model"
	"time"
)

type UserResponse struct {
	BaseResponse

	Account        string    `json:"account" comment:"账号"`
	Name           string    `json:"name" comment:"名字"`
	Email          string    `json:"email" comment:"邮件"`
	Mobile         string    `json:"mobile" comment:"手机号"`
	SignUpTime     time.Time `json:"sign_up_time" comment:"注册时间"`
	LastSignInTime time.Time `json:"last_sign_in_time" comment:"最近登录时间"`
	PicPath        string    `json:"pic_path" comment:"头像地址"`
}

type SignInResponse struct {
	User  *UserResponse
	Token string `json:"token" comment:"token"`
}

func UserResponseFrom(u *model.User) *UserResponse {
	ur := UserResponse{}

	ur.CreatedAt = u.CreatedAt
	ur.UpdatedAt = u.UpdatedAt

	ur.Account = u.Account
	ur.Name = u.Name
	ur.Email = u.Email
	ur.PicPath = u.PicPath
	ur.SignUpTime = u.SignUpTime
	ur.LastSignInTime = u.LastSignInTime
	return &ur
}

func SignInResponseFrom(u *model.User, token string) *SignInResponse {
	r := SignInResponse{}
	r.User = UserResponseFrom(u)
	r.Token = token

	return &r
}
