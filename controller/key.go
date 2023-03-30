package controller

import (
	"net/http"
	"belajar-api/helper"
	"belajar-api/models"

	"github.com/labstack/echo/v4"

)

type KeyController struct {
	model models.KeyModel
}

func (kc *KeyController) SetModel(m models.KeyModel) {
	kc.model = m
}


func (kc *KeyController) AskKey() echo.HandlerFunc{
	return func(c echo.Context) error {
		res, err := kc.model.AddKey()
		if err != nil {
			c.Logger().Error("terjadi kesalah saat add key ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "tidak bisa generate key", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "harap masukkan bagian key pada kolom api-key", res))
	}
}