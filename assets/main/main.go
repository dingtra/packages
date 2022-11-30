package main

import (
	"net/http"
	"assets"
	"os"
	"log"
	"fmt"
)


func main () {
	AssetsRoutes := []string{"/assets/", "/branch/"}

	for _, item := range AssetsRoutes {
		http.HandleFunc(item, assets.Http)
	}

	log.Fatal(http.ListenAndServe(Getenver(), nil))
}


func Getenver() string {
	var port = os.Getenv("PORT")

	if port == ""{
		port = "7070"
		fmt.Println("INFO: no port detected")
	}else{
		fmt.Println("HEROKU PORT FOUND "+port)
	}

	return ":"+port
}