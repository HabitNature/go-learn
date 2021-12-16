package model

import (
	"gorm.io/gorm"
	"time"
)

// 参见 : https://blog.csdn.net/gglinux/article/details/68948901
// 参见 : https://www.cnblogs.com/jiqing9006/p/5937733.html

type User struct {
	gorm.Model
	Account        string    `json:"account" gorm:"column:account;type:varchar(16);comment:'账号';default:NULL;"`
	Name           string    `json:"name" gorm:"column:name;type:varchar(16);comment:'名字';default:NULL;"`
	Password       string    `json:"password" gorm:"column:password;type:varchar(32);comment:'密码';default:NULL;"`
	Email          string    `json:"email" gorm:"column:email;type:varchar(64);comment:'邮件';default:NULL;"`
	Mobile         string    `json:"mobile" gorm:"column:mobile;type:varchar(64);comment:'邮件';default:NULL;"`
	SignUpTime     time.Time `json:"sign_up_time" gorm:"column:sign_up_time;type:datetime;comment:'注册时间';default:NULL;"`
	LastSignInTime time.Time `json:"last_sign_in_time" gorm:"column:last_sign_in_time;type:datetime;comment:'最近登录时间';default:NULL;"`
	PicPath        string    `json:"pic_path" gorm:"column:pic_path;type:varchar(128);comment:'头像地址';default:NULL;"`
}
