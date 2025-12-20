package http

import (
	"net/http"
	"encoding/json"
	"github.com/alireza/identity/application"
	"errors"
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
	if r.Method != http.MethodPost{
		http.Error(w, "Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	req := NewLoginRequest()
	if err:= json.NewDecoder(r.Body).Decode(&req) ; err!=nil{
		http.Error(w, "InvalidJson" , http.StatusBadRequest)
		return
	} 
	 err := l.LogService.Login(req.Name , req.Password)
	if err != nil{
		switch {
		case errors.Is(err , application.ErrUserNotFound):
			    respond := NewLoginResponse(false , "User Not Found")
				w.Header().Set("Content/Type" , "application/json")
				w.WriteHeader(http.StatusCreated)
				if err:= json.NewEncoder(w).Encode(&respond) ; err!= nil{
					http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
					return
				}
			default:
				respond := NewLoginResponse(false , "Invalid Password")
				w.Header().Set("Content/Type" , "application/ json")
				w.WriteHeader(http.StatusCreated)
				if err:= json.NewEncoder(w).Encode(&respond) ; err!= nil{
					http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
					return
				}


		}
	}

	
		respond := NewLoginResponse(true , "")
		w.Header().Set("Content/Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(&respond); err!=nil{
			http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
		}
	}
	
	




	

