package main

import (
	"log"
	"net/http"

	app "github.com/alireza/identity/internal/application"
	httpHandler "github.com/alireza/identity/internal/presentation/http"
	"github.com/alireza/identity/internal/infra/memory"
	"github.com/alireza/identity/internal/infra/crypto"
	"github.com/joho/godotenv"
)


func main() {

	godotenv.Load(".env")
	repo := memory.NewInMemoryRepository() //infra
	hasher  := crypto.NewBcryptHasher()   // infra 
	Comparer := crypto.NewBcryptComparer() // infra 
	userService := app.NewUserService(repo , hasher) // application 
	loginservice := app.NewLoginService(repo , Comparer) // application 
	userHandler := httpHandler.NewUserHandler(userService) // presentaion
	loginhandler := httpHandler.NewLoginHandler(loginservice) // presentaion 


	mux := http.NewServeMux()

	mux.HandleFunc("/register",userHandler.Register)
	mux.HandleFunc("/login", loginhandler.Login)
	log.Println("Server listening on:8080")
	log.Fatal(http.ListenAndServe(":8080",mux))
	




}
