package helper 

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"belajar-api/models"
)
type ValidateObj struct {
	DB*gorm.DB
}

func (vo *ValidateObj) Validatekey(key string, c echo.Context) (bool, error) {
	var KeyResult models.Keys
	err := vo.DB.Select("key").Where("`key` = ?", key).First(&KeyResult).Error
	if err != nil {
		return false, err
	}
	return key == KeyResult.Key, nil
}