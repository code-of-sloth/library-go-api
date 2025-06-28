package routes

import (
	"LibraryGo/src/handler"

	"github.com/labstack/echo/v4"
)

func LoadEndpoints(e *echo.Echo) {
	apiGrp := e.Group("/api")

	loadVerGrp(apiGrp)
}

func loadVerGrp(apiGrp *echo.Group) {
	verGrp := apiGrp.Group("/v1")

	verGrp.POST("/user", handler.CreateUser)
	verGrp.GET("/user", handler.FetchUser) // query parameter userid
	verGrp.PUT("/user", handler.UpdateUser)
	verGrp.DELETE("/user", handler.DeleteUser)

	verGrp.POST("/book", handler.AddBook)
	verGrp.DELETE("/book", handler.RemoveBook)
	verGrp.POST("/allbook", handler.FetchBooks) //query parameter bookid

	verGrp.POST("rent", handler.RentBook)
	verGrp.POST("return", handler.ReturnBook)

}
