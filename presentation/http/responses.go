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