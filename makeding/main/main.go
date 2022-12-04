package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
	"makeding"
)



func main(){
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/xoz/qr1/", makeding.Http)
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

