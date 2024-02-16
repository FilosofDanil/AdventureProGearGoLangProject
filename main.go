package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", HelloWorld)
	go provideLink()
	if err := e.StartAutoTLS(":8443"); err != nil {
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
