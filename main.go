package main

import (
	"log"
	"net/http"

	app "github.com/alireza/identity/application"
	httpHandler "github.com/alireza/identity/presentation/http"
	"github.com/alireza/identity/infra/memory"
	"github.com/alireza/identity/infra/crypto"
)


func main() {
	repo := memory.NewInMemoryRepository()
	hasher  := crypto.NewBcryptHasher()
	userService := app.NewUserService(repo , hasher)
	userHandler := httpHandler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("/register",userHandler.Register)
	log.Println("Server listening on:8080")
	log.Fatal(http.ListenAndServe(":8080",mux))
	




}
