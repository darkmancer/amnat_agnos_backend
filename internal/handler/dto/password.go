package dto

type PasswordRequest struct {
	InitPassword string `json:"init_password" binding:"required,min=1,max=40"`
}

type PasswordResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}
