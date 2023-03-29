package models

import (
	"errors"
	"log"
	"gorm.io/gorm"
)

type UserModel struct {
	db*gorm.DB
}

func (um *UserModel) SetDB(db*gorm.DB) {
	um.db = db
}

func (um *UserModel) Insert(newUser User) (User, error) {
	if err := um.db.Create(&newUser).Error; err != nil {
		log.Println("Terjadi error saat create user", err.Error())
		return User{}, err
	}

	return newUser,nil
}


func (um *UserModel) Login(hp, password string) (User, error) {
	res := User{}
	if err := um.db.Where("hp = ? and password = ?", hp, password).Find(&res).Error; err != nil {
		log.Println("Terjadi error saat create user", err.Error())
		return User{}, err
	}

	if res.Phone == "" {
		log.Println("Data tidak ditemukan")
		return User{}, errors.New("data tidak ditemukan")
	}
	return res, nil
}

func (um *UserModel) GetAllUser() ([]User, error) {
	res := []User{}

	if err := um.db.Select("hp, nama, id").Find(&res).Error; err != nil {
		log.Println("Terjadi error saat select user", err.Error())
		return nil, err
	}
	return res, nil
}

func (um *UserModel)Update() ([]User, error) {
	res := []User{}
	if err := um.db.Save("name,phone,password").Find(&res).Error; err != nil {
		log.Println("Terjadi kesalahan saat update user", err.Error())
		return nil, err
	}
	return res, nil
}