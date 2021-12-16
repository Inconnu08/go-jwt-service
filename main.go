package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"socialmedia-auth/internal/handlers"
	"socialmedia-auth/internal/service"
	"socialmedia-auth/internal/service/token"
)

func main() {
	prvKey, err := ioutil.ReadFile("internal/cert/app.rsa")
	if err != nil {
		log.Fatalln(err)
	}

	pubKey, err := ioutil.ReadFile("internal/cert/app.rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := token.NewJWT(prvKey, pubKey)
	s := service.New(jwtToken)
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", 8080),
		Handler:           handlers.New(s),
		ReadHeaderTimeout: time.Second * 5,
		ReadTimeout:       time.Second * 15,
	}
	log.Println("accepting connections on port 8080\nstarting server at http://localhost:8080")
	if err = server.ListenAndServe(); err != nil {
		log.Fatalf("could not listen and serve: %v", err)
		return
	}
}
