package http


import (
	"net/http"
	"encoding/json"
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
	err := h.services.Register(req.Name , req.Password)
	if err != nil{
		http.Error(w , "could not Register", http.StatusInternalServerError)
		
	}
}