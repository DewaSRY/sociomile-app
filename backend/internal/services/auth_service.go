package services

import (
	"errors"

	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtutil "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"

	"gorm.io/gorm"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Register creates a new user account
func (s *AuthService) Register(req requestdto.RegisterRequest) (*responsedto.AuthResponse, error) {
	var existingUser models.UserModel
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	}

	user := models.UserModel{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password, 
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	token, err := jwtutil.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &responsedto.AuthResponse{
		Token: token,
		User: responsedto.UserData{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(req requestdto.LoginRequest) (*responsedto.AuthResponse, error) {
	var user models.UserModel
	
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, errors.New("failed to fetch user")
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwtutil.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &responsedto.AuthResponse{
		Token: token,
		User: responsedto.UserData{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}

func (s *AuthService) GetUserByID(userID uint) (*models.UserModel, error) {
	var user models.UserModel
	if err := database.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to fetch user")
	}
	return &user, nil
}

func (s *AuthService) RefreshToken(tokenString string) (string, error) {
	return jwtutil.RefreshToken(tokenString)
}
