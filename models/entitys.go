package models

import(
	
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Books []Book `gorm:"foreignKey:UserID"`
}

type Book struct {
	gorm.Model
	Judul    string `json:"judul"`
	Tahun    string `json:"tahun"`
	Penerbit string `json:"penerbit"`
	UserID   uint   `json:"user_id"`
}


type Response struct {
	gorm.Model
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Keys struct {
	gorm.Model
	Key string
}