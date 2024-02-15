package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("", HelloWorld)
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Link: http://localhost:8080")
}

func HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello World!")
}
