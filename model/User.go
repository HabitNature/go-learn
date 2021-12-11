package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Name string `json:"name" gorm:"column:name;type:varchar(16);comment:'名字';default:NULL;"`

	BirthTime time.Time `json:"birth_time" gorm:"column:birth_time;comment:'出生时间';default:NULL;"`
	DeathTime time.Time `json:"death_time" gorm:"column:death_time;comment:'死亡时间';default:NULL;"`
}
