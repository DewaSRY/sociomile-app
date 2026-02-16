package services

import (
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
)


type AuthService interface{
	Register(req requestdto.RegisterRequest) (*responsedto.AuthResponse, error)
	Login(req requestdto.LoginRequest) (*responsedto.AuthResponse, error)
	GetUserByID(userID uint) (*responsedto.UserProfileData, error)
	RefreshToken(tokenString string) (string, error)
}


