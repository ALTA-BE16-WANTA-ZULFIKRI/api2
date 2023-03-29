package routes

import (
	"belajar-api/controller"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, uc controller.UserController, bc controller.BookController) {
	e.POST("/users", uc.Register)
	e.POST("/login", uc.Login())
	e.GET("/users", uc.GetUser())
	e.PUT("/update", uc.Update())
	// e.GET("/users/:user_id/books")

	e.GET("/books/:bookId", bc.GetBookByID())
	e.GET("/books", bc.GetBook())
	e.POST("/books", bc.AddBook)
	e.PUT("/books", bc.Edit)
	e.DELETE("/books", bc.Delete)
}
