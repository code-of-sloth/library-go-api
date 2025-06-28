package main

import (
	"LibraryGo/src/config"
	"LibraryGo/src/models"
	"LibraryGo/src/routes"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	err := config.InitCore()
	if err != nil {
		log.Fatal(err)
	}

	defer config.InitClose()
	e := echo.New()

	e.Validator = new(models.Validator)
	routes.LoadEndpoints(e)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
