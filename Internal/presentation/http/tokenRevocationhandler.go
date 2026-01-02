package http

import (
	"encoding/json"
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
	
	
}
