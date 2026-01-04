package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alireza/identity/internal/application"
)


type TokenRevocationHandler struct {
	tokenRevocationService application.TokenRevocationService
}

func NewTokenRevocationHandler(service application.TokenRevocationService) TokenRevocationHandler {
	return TokenRevocationHandler{
		tokenRevocationService: service,
	}
}

func (t *TokenRevocationHandler) RevokeToken(w http.ResponseWriter, r *http.Request) {{
	ctx := r.Context()
	if r.Method != http.MethodPatch{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	req := NewTokenRevocationRequest()
	if err:= json.NewDecoder(r.Body).Decode(&req) ; err!=nil{
		http.Error(w, "InvalidJson" , http.StatusBadRequest)
		return
	}
	accessToken, recoveryToken, err := t.tokenRevocationService.RevokeToken(ctx , req.RefreshToken)
	if err!= nil{
    switch {
	case errors.Is(err , application.ErrInvalidToken):
		http.Error(w, "InvalidToken" , http.StatusUnauthorized)
		return
	case errors.Is(err , application.ErrTokenAlreadyRevoked):
		w.Header().Set("Content-Type","application/json")
		response := NewTokenRevocationResponse(false , "Token Already Revoked" , "" , "")
		if err := json.NewEncoder(w).Encode(response);err!=nil{
			http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
			return
		}
		return
	default:
		http.Error(w, "Internal Server Error" , http.StatusInternalServerError)
		return
	}
	}
	
	

	response := NewTokenRevocationResponse(true , "Token Revoked Successfully" , accessToken , recoveryToken)	

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response);err!=nil{
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
}
}