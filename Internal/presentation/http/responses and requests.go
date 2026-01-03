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
	Name string `json:"username"`
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
	JWTRecoveryToken  string `json:"jwtRecoveryToken"`
	JWTAccessToken    string `json:"jwtAccessToken"`

}

func NewLoginResponse(sucess bool , messege string , accessToken , recoveryToken string) LoginResponse{
	return LoginResponse{
		Success: sucess,
		Messege: messege,
		JWTAccessToken: accessToken,
		JWTRecoveryToken: recoveryToken,
	}
}

type TokenRevocationRequest struct{
	RefreshToken string `json:"refreshToken"`
}

func NewTokenRevocationRequest() TokenRevocationRequest{
	return TokenRevocationRequest{
		RefreshToken: "",
	}
}

type TokenRevocationResponse struct{
	Success bool `json:"success"`
	Messege string `json:"Messege"`
	JWTRecoveryToken  string `json:"jwtRecoveryToken"`
	JWTAccessToken    string `json:"jwtAccessToken"`

}

func NewTokenRevocationResponse(sucess bool , messege string , accessToken , recoveryToken string) TokenRevocationResponse{
	return TokenRevocationResponse{
		Success: sucess,
		Messege: messege,
		JWTAccessToken: accessToken,
		JWTRecoveryToken: recoveryToken,
	}
}