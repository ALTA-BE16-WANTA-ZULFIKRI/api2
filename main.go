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
	cfg.AutoMigrate(models.Keys{})

	mdl := models.UserModel{}
	mdl.SetDB(cfg)
	ctl := controller.UserController{}
	ctl.SetModel(mdl)

	bookMdl := models.BookModel{}
	bookMdl.SetDBB(cfg)
	bookCtl := controller.BookController{}
	bookCtl.SetModel(bookMdl)
	bookCt2 := controller.UserController{}
	bookCt2.SetModel(mdl)


	keyMdl := models.KeyModel{}
	keyMdl.SetModel(cfg)
	keyCtl := controller.KeyController{}
	keyCtl.SetModel(keyMdl)


	// ROUTING
	routes.Route(e,ctl,bookCtl,keyCtl, cfg)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}
