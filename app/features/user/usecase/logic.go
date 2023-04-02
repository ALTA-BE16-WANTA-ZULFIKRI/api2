package usecase

import (
	"belajar-api/app/features/user"
	"errors"

	"github.com/labstack/gommon/log"
)

type UserLogic struct {
	m user.Repository
}

func New(r user.Repository) user.UseCase {
	return &UserLogic{
		m: r,
	}
}


func (ul *UserLogic) Login(hp string, password string) (user.Core, error) {
	result, err := ul.m.Login(hp, password)
	if err != nil {
		return user.Core{}, errors.New("terdapat permasalahan pada server")
	}

	return result, nil
}

func (ul *UserLogic) Register(newUser user.Core) error {
	_, err := ul.m.Insert(newUser)
	if err != nil {
		log.Error("register logic error:", err.Error())
		return errors.New("terjadi kesalahan pada server")
	}

	return nil
}
