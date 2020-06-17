package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"server/dbmysql"
)

type Player struct {
	UserId    int `json:"user_id" gorm:"primary_key; type:int(28);not null"`
	UserName  string `json:"user_name" gorm:"type:varchar(32);not null;default:''"`
	Level int    `json:"level" gorm:"type:int(7);not null;default:0"`
}

func (info Player) Insert() bool {
	create := dbmysql.Db.Create(&info)
	if create.Error != nil {
		return false
	}
	return true
}

//update
func (info Player) Updates(m map[string]interface{}) bool {
	if info.UserId <= 0 {
		return false
	}
	err := dbmysql.Db.Model(&info).Updates(m).Error
	if err != nil {
		return false
	}
	return true
}

//insert and update
func (info Player) Save() bool {
	if info.UserId <= 0 {
		return false
	}
	create := dbmysql.Db.Save(&info)
	if create.Error != nil {
		return false
	}
	return true
}

//select
func (info *Player) First() (int, error) {
	if info.UserId <= 0 {
		return -1, fmt.Errorf("First:UsernoExist")
	}
	err := dbmysql.Db.First(&info).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0, nil
	} else if err != nil {
		return -1, err
	}
	return 1, nil
}

// Delete
func (info *Player)Delete() bool {
	if info.UserId <= 0 {
		return false
	}
	if err := dbmysql.Db.Delete(info).Error; err != nil {
		return false
	}
	return true
}