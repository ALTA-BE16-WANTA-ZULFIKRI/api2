package controller

import (
	"net/http"
	"belajar-api/helper"
	"belajar-api/models"
	"github.com/labstack/echo/v4"
)


type UserController struct {
	model models.UserModel
}

func (uc *UserController) SetModel(m models.UserModel) {
	uc.model = m
}

func (uc *UserController) Register(c echo.Context) error {
	input := models.User{}
	if err := c.Bind(&input); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
	}

	res, err := uc.model.Insert(input)

	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
	}

	return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", res))
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		type loginData struct {
			Hp       string `json:"hp"`
			Password string `json:"password"`
		}
		input := loginData{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())

			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		res, err := uc.model.Login(input.Hp, input.Password)

		if err != nil {
			c.Logger().Error("terjadi kesalahan ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses login", res))
	}
}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.model.GetAllUser()

		if err != nil {
			c.Logger().Error("user model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}

}	

func (uc *UserController) Update() echo.HandlerFunc {
	return func (c echo.Context) error {
		res, err := uc.model.GetAllUser()

		if err != nil {
			c.Logger().Error("user model error", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError,"terdapat kesalahan pada server", nil))
		} 
		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}