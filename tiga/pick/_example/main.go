package main

import (
	"github.com/liov/tiga/tiga/pick"
	"github.com/liov/tiga/tiga/pick/_example/service"
	"log"
	"net/http"

	_ "github.com/liov/tiga/tiga/pick/_example/service"
)

func main() {
	pick.RegisterService(&service.UserService{}, &service.TestService{}, &service.StaticService{})
	router := pick.New(false, "httptpl")
	router.ServeFiles("/static", "E:/")
	log.Println("visit http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
