package http

import (
	"net/http"
	"encoding/json"
	"github.com/alireza/identity/internal/application"
	"errors"
	"os"
)

type LoginHandler struct{
	LogService *application.LoginSerivce
}

func NewLoginHandler(logservice *application.LoginSerivce) LoginHandler{
	return LoginHandler{
		LogService: logservice,
	}
}

func (l *LoginHandler) Login(w http.ResponseWriter , r *http.Request){
	ctx := r.Context()
	switch r.Method{
	case http.MethodGet:
		http.ServeFile(w , r , os.Getenv("LOGIN_PAGE"))
	case http.MethodPost:

		req := NewLoginRequest()
	if err:= json.NewDecoder(r.Body).Decode(&req) ; err!=nil{
		http.Error(w, "InvalidJson" , http.StatusBadRequest)
		return
	} 
	 accessToken , recoveryToken , err := l.LogService.Login(ctx ,req.Name , req.Password)
	if err != nil{
		switch {
		case errors.Is(err , application.ErrUserNotFound):
			    respond := NewLoginResponse(false , "User Not Found", "" , "")
				w.Header().Set("Content/Type" , "application/json")
				w.WriteHeader(http.StatusCreated)
				if err:= json.NewEncoder(w).Encode(&respond) ; err!= nil{
					http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
					return
				}
		case errors.Is(err , application.ErrCanNotAuthenticate):
				respond := NewLoginResponse(false , "To many Active Sessions" , "" , "")
				w.Header().Set("Content/Type" , "application/json")
				w.WriteHeader(http.StatusCreated)
				if err:= json.NewEncoder(w).Encode(&respond) ; err!= nil{
					http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
					return
				}
			default:
				respond := NewLoginResponse(false , "Invalid Password" , "" , "")
				w.Header().Set("Content/Type" , "application/ json")
				w.WriteHeader(http.StatusCreated)
				if err:= json.NewEncoder(w).Encode(&respond) ; err!= nil{
					http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
					return
				}
				return


		}
	}

	
		respond := NewLoginResponse(true , "" , accessToken , recoveryToken )
		w.Header().Set("Content/Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(&respond); err!=nil{
			http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
		}
	default:
		http.Error(w , "Method not allowed", http.StatusMethodNotAllowed)
	}
	
	}
	
	




	

