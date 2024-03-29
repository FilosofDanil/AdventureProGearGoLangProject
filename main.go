package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	e := echo.New()
	e.GET("", HelloWorld)
	go provideLink()
	port := os.Getenv("PORT")
	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}

func HelloWorld(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello World!")
}

func provideLink() {
	time.Sleep(1 * time.Second)
	fmt.Println("Link: 104.196.232.237:8080")
}
