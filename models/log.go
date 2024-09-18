package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Log struct {
	GormModel
	Water       uint   `gorm:"not null" json:"water" form:"water" valid:"required~Value water is required"`
	Wind        uint   `gorm:"not null" json:"wind" form:"wind" valid:"required~Value Water is required"`
	StatusWind  string `json:"status_wind" form:"status_wind"`
	StatusWater string `json:"status_water" form:"status_water"`
}

func (u *Log) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}
	return
}
