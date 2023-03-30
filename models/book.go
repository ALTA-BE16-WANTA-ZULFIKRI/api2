package models

import (
	"log"

	"gorm.io/gorm"
)

type BookModel struct {
	db *gorm.DB
}

func (bm BookModel) SetDBB(db *gorm.DB) {
       bm.db = db
}

func (bm BookModel) Insert(newBook Book) (Book, error) {
	if err := bm.db.Create(&newBook).Error; err != nil {
		log.Println("Terjadi error saat create Book", err.Error())
		return Book{}, err
	}
	return newBook, nil
}

func (bm *BookModel) GetAllBook(userID uint) (any, error) {
	type ExpectedRespond struct {
		Judul string `json:"judul"`
		Tahun string `json:"tahun"`
		Nama  string `json:"nama"`
	}
	res := []ExpectedRespond{}
	var err error
	if userID != 0 {
		err = bm.db.Table("books").Select("books.judul as judul, books.tahun, users.nama as nama").Joins("JOIN users on users.id = books.user_id").Where("books.user_id = ?", userID).Scan(&res).Error
	} else {
		err = bm.db.Table("books").Select("books.judul as judul, books.tahun, users.nama as nama").Joins("JOIN users on users.id = books.user_id").Scan(&res).Error
	}

	if err != nil {
		log.Println("Terjadi error saat select Book ", err.Error())
		return nil, err
	}

	return res, nil
}

func (bm *BookModel) GetBookByID(id uint) (Book, error) {
	res := Book{}

	if err := bm.db.Find(&res, id).Error; err != nil {
		log.Println("Terjadi error saat select Book ", err.Error())
		return Book{}, err
	}

	return res, nil
}

func (bm *BookModel) Edit(id uint) (Book, error) {
	res := Book{}

	if err := bm.db.Save(&res.UpdatedAt).Error; err != nil {
		log.Println("Terjadi error saat edit Book ", err.Error())
		return Book{}, err
	}
	return res, nil
}

func (bm *BookModel) Delete(id uint) (Book, error) {
	res := Book{}

	if err := bm.db.Delete(&res.UserID).Error; err != nil {
		log.Println("Terjadi error saat delete Book ", err.Error())
		return Book{}, err
	}
	return res, nil
}
