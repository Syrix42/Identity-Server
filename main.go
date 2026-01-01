package main

import (
	"log"
	"net/http"

	app "github.com/alireza/identity/internal/application"
	httpHandler "github.com/alireza/identity/internal/presentation/http"
	db "github.com/alireza/identity/internal/infra/data"
	"github.com/alireza/identity/internal/infra/crypto"
	"github.com/joho/godotenv"
)


func main() {

	godotenv.Load(".env")
	hasher:= crypto.NewBcryptHasher()   // infra 
	Dbrepo := db.NewMYSQLDb()
	Comparer := crypto.NewBcryptComparer() // infra
	userService := app.NewUserService(Dbrepo, hasher) // application 
	loginservice := app.NewLoginService(Dbrepo, Comparer) // application 
	userHandler := httpHandler.NewUserHandler(userService) // presentaion
	loginhandler := httpHandler.NewLoginHandler(loginservice) // presentaion 


	mux := http.NewServeMux()

	mux.HandleFunc("/register",userHandler.Register)
	mux.HandleFunc("/login", loginhandler.Login)
	log.Println("Server listening on:8080")
	log.Fatal(http.ListenAndServe(":8080",mux))
	




}
