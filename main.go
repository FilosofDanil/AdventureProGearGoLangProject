package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", HelloWorld)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
	//e := echo.New()
	//e.GET("", HelloWorld)
	//go provideLink()
	//s := http.Server
	//if err := e.StartServer(":8080"); err != nil {
	//	log.Fatal(err)
	//}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello world!\n")
}

func provideLink() {
	time.Sleep(1 * time.Second)
	fmt.Println("Link: 104.196.232.237:8080")
}
