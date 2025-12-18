package http

import (
	"encoding/json"
	"net/http"
	"errors"
	"github.com/alireza/identity/application"
)



type UserHandler struct{
	services *application.UserService
}

func NewUserHandler (service *application.UserService) *UserHandler{
	return  &UserHandler{services: service}
}

func (h *UserHandler) Register(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed" , http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err!= nil{
		http.Error(w, "Invalid json" , http.StatusBadRequest)
		return
	}
	
	err :=  h.services.Register(req.Name , req.Password)
	if err != nil{
		switch {
		case errors.Is(err , application.ErrUserAlreadyExists):
			Response:= NewRegistrationResponse(false , "User Already Exists" , "")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(Response); err != nil {
        http.Error(w, "Server Problem", http.StatusInternalServerError)
    }

		default:
			http.Error(w, "Internal Server Error " , http.StatusInternalServerError)
		}
	
		
	}else{

		Response := NewRegistrationResponse(true, "User Registered Successfully", "")

    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    if err := json.NewEncoder(w).Encode(Response); err != nil {
        http.Error(w, "Server Problem", http.StatusInternalServerError)
    }

	}
}