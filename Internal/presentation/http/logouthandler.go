package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/alireza/identity/internal/application"
)



type LogoutHandler struct{
	logoutService *application.LogOutService
}



func NewLogoutHandler(service *application.LogOutService) LogoutHandler {
	return LogoutHandler{
		logoutService: service,
	}
}

func (l *LogoutHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPost{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	req := NewLogoutRequest()
	if err:= json.NewDecoder(r.Body).Decode(&req) ; err!=nil{
		http.Error(w, "InvalidJson" , http.StatusBadRequest)
		return
	}
	err := l.logoutService.Logout(ctx , req.RefreshToken)
	if err!= nil{
	switch {
	case errors.Is(err, application.ErrInvalidToken):
		http.Error(w, "InvalidToken" , http.StatusUnauthorized)
		return
	default:
		fmt.Println(err)
		http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
		return
	}

	}
	response := NewLogoutResponse(true , "Logout happened succsesfully")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response);err!=nil{
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}	
}
	