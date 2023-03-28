package main 

import (
	
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type Book struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Published string `json:"published"`
}

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}



var DataStoreUsers = make(map[string]User, 0)
var DataStoreBooks = make(map[string][]Book, 0)


func CR(code int, data interface{}, status, message string) *Response {
	return &Response{Status: status, Code: code, Message: message, Data: data}
}

func registerUser(c echo.Context) error {
	var req User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CR(http.StatusBadRequest, nil, "BAD REQUEST", "Username/Password Salah"))
	}
	if data, ok := DataStoreUsers[req.Phone]; ok {
		if data.Password == req.Password {
			return c.JSON(http.StatusOK, CR(http.StatusOK, DataStoreUsers[req.Phone], "Success", "Berhasil Login"))
		}
	}
	return c.JSON(http.StatusBadRequest, CR(http.StatusBadRequest, nil, "BAD REQUEST", "Username/Password Salah"))
}

func loginUser(c echo.Context) error {
	var req User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CR(http.StatusBadRequest, nil, "BAD REQUEST", "Gagal Mendaftar"))
	}
	DataStoreUsers[req.Phone] = req
	return c.JSON(http.StatusOK, CR(http.StatusOK, nil, "Success", "selamat anda telah terdaftar"))
}

func addBook(c echo.Context) error {
	var req Book
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CR(http.StatusBadRequest, nil, "BAD REQUEST", "Gagal Menambahkan Data Buku"))
	}
	DataStoreBooks["books"] = append(DataStoreBooks["books"], req)
	return c.JSON(http.StatusOK, CR(http.StatusOK, nil, "Success", "data buku yang berhasil diinputkan"))
}

func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, CR(http.StatusOK, DataStoreBooks["books"], "Success", "data buku yang berhasil diinputkan"))
}

func main () {
	e := echo.New()
	
	//register user
	e.POST("/user", registerUser)

	//login user
	e.POST("login", loginUser)

	//add book 
	e.POST("/book", addBook)

	// get all books
	e.GET("/books", getBooks)

	// start server 
	e.Logger.Fatal(e.Start(":1323"))
}