package dto

type UserRegistrationRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string        `json:"token"`
	User  UserLoginInfo `json:"user"`
}

type UserLoginInfo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
}
