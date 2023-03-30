package models 


import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KeyModel struct {
	db*gorm.DB
}

func (km *KeyModel) SetModel(d *gorm.DB) {
	km.db = d
}


func (km *KeyModel) AddKey() (Keys, error) {
	res := Keys{}
	uid := uuid.New()
	res.Key = uid.String()


	err := km.db.Create(&res).Error
	if err != nil {
		return Keys{}, err
	}
	return res, nil
}