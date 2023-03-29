package main

import (
	"belajar-api/config"
	"belajar-api/controller"
	"belajar-api/models"
	"belajar-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitSQL()
	cfg.AutoMigrate(models.User{})
	cfg.AutoMigrate(models.Book{})

	mdl := models.UserModel{}
	mdl.SetDB(cfg)

	bookMdl := models.BookModel{}
	bookMdl.SetDBB(cfg)
	bookCtl := controller.BookController{}
	bookCtl.SetModel(bookMdl)
	bookCt2 := controller.UserController{}
	bookCt2.SetModel(mdl)

	// ROUTING
	routes.Route(e,bookCt2,bookCtl)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}
