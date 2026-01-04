package main

import (
	"context"
	"log"
	"net/http"

	app "github.com/alireza/identity/internal/application"
	"github.com/alireza/identity/internal/infra/crypto"
	db "github.com/alireza/identity/internal/infra/data"
	httpHandler "github.com/alireza/identity/internal/presentation/http"
	"github.com/joho/godotenv"
	"github.com/alireza/identity/internal/infra/revokedtoken"
)


func main() {

	godotenv.Load(".env")
	hasher:= crypto.NewBcryptHasher()   // infra 
	Dbrepo := db.NewMYSQLDb()
	Tokenrepo := db.NewTokenDataRepository() // infra
	tokenRevocationService := app.NewTokenRevocationService(Tokenrepo , Dbrepo) // application
	tokenRevocationHandler := httpHandler.NewTokenRevocationHandler(tokenRevocationService) // presentation
	Comparer := crypto.NewBcryptComparer() // infra
	userService := app.NewUserService(Dbrepo, hasher) // application 
	loginservice := app.NewLoginService(Dbrepo, Comparer) // application 
	userHandler := httpHandler.NewUserHandler(userService) // presentaion
	loginhandler := httpHandler.NewLoginHandler(loginservice) // presentaion 
	logoutService := app.NewLogOutService(Dbrepo ,Tokenrepo)
	logoutHandler := httpHandler.NewLogoutHandler(logoutService)

	ctx := context.Background()

	revokedtoken.StartDailyCleanup(ctx , 7)

	mux := http.NewServeMux()

	mux.HandleFunc("/register",userHandler.Register)
	mux.HandleFunc("/login", loginhandler.Login)
	mux.HandleFunc("/revocation", tokenRevocationHandler.RevokeToken)
	mux.HandleFunc("/logout" , logoutHandler.Logout)
	log.Println("Server listening on:8080")
	log.Fatal(http.ListenAndServe(":8080",mux))
	
}

