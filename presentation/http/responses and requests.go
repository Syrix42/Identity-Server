package http

type RegistrationResponse struct{
	Success bool `json:"success"`
	Messege string `json:"messege"`
	Data any `json"data,omitempty"`
}


func NewRegistrationResponse(sucess bool , Messege string , Data any)RegistrationResponse{
	return  RegistrationResponse{
		Success: sucess,
		Messege: Messege,
		Data: Data,
	}

}

type LoginRequest struct{
	Name string `json:"Name"`
	Password string `json:"Password"`
}
func NewLoginRequest ()LoginRequest{
	return LoginRequest{
		Name: "",
		Password: "",
	}
}

type LoginResponse struct{
	Success bool `json:"success"`
	Messege string `json:"Messege"`

}

func NewLoginResponse(sucess bool , messege string) LoginResponse{
	return LoginResponse{
		Success: sucess,
		Messege: messege,
	}
}