package impl

import (
	"errors"

	"DewaSRY/sociomile-app/internal/database"
	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"

	"gorm.io/gorm"
)


type authServiceImpl struct{
	jwtService jwtLib.JwtService
}


func InstanceAuthService() services.AuthService {
	return &authServiceImpl{
		jwtService: jwtLib.InstanceJwtService(),
	}
}

func (t *authServiceImpl) Register(req requestdto.RegisterRequest) (*responsedto.AuthResponse, error) {
	var existingUser models.UserModel
	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	}

	var guestRole models.UserRoleModel
	if err := database.DB.Where("name = ?", models.RoleGuest).First(&guestRole).Error; err != nil {
		guestRole.ID = 4
	}

	user := models.UserModel{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		RoleID:   guestRole.ID,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	token, err := t.jwtService.GenerateToken(user.ID, user.Email, user.RoleID)
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

func (t *authServiceImpl) Login(req requestdto.LoginRequest) (*responsedto.AuthResponse, error) {
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

	token, err := t.jwtService.GenerateToken(user.ID, user.Email, user.RoleID)
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

func (s *authServiceImpl) GetUserByID(userID uint) (*models.UserModel, error) {
	var user models.UserModel
	if err := database.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to fetch user")
	}
	return &user, nil
}

func (t *authServiceImpl) RefreshToken(tokenString string) (string, error) {
	return t.jwtService.RefreshToken(tokenString)
}
