package impl

import (
	"errors"

	"DewaSRY/sociomile-app/internal/services"
	"DewaSRY/sociomile-app/pkg/dtos/requestdto"
	"DewaSRY/sociomile-app/pkg/dtos/responsedto"
	jwtLib "DewaSRY/sociomile-app/pkg/lib/jwt"
	"DewaSRY/sociomile-app/pkg/models"

	"gorm.io/gorm"
)

type authServiceImpl struct {
	db         *gorm.DB
	jwtService jwtLib.JwtService
}

func NewAuthService(db *gorm.DB, jwtService jwtLib.JwtService) services.AuthService {
	return &authServiceImpl{
		db:         db,
		jwtService: jwtService,
	}
}

func (t *authServiceImpl) Register(req requestdto.RegisterRequest) (*responsedto.AuthResponse, error) {
	var existingUser models.UserModel
	if err := t.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	}

	var guestRole models.UserRoleModel
	if err := t.db.Where("name = ?", models.RoleGuest).First(&guestRole).Error; err != nil {
		guestRole.ID = 4
	}

	user := models.UserModel{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		RoleID:   guestRole.ID,
	}

	if err := t.db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	if err := t.db.
		Preload("Role").
		Preload("Organization").
		First(&user, user.ID).Error; err != nil {
		return nil, errors.New("failed to load user relations")
	}

	token, err := t.jwtService.GenerateToken(user.ID, user.Email, user.RoleID, user.OrganizationID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &responsedto.AuthResponse{
		Token: token,
		User:  t.mappedToUserProfile(&user),
	}, nil
}

func (t *authServiceImpl) Login(req requestdto.LoginRequest) (*responsedto.AuthResponse, error) {
	var user models.UserModel

	if err := t.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, errors.New("failed to fetch user")
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid email or password")
	}

	if err := t.db.
		Preload("Role").
		Preload("Organization").
		First(&user, user.ID).Error; err != nil {
		return nil, errors.New("failed to load user relations")
	}

	token, err := t.jwtService.GenerateToken(user.ID, user.Email, user.RoleID, user.OrganizationID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &responsedto.AuthResponse{
		Token: token,
		User: t.mappedToUserProfile(&user),
	}, nil
}

func (t *authServiceImpl) GetUserByID(userID uint) (*responsedto.UserProfileData, error) {
	var user models.UserModel
	if err := t.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to fetch user")
	}

	if err := t.db.
		Preload("Role").
		Preload("Organization").
		First(&user, user.ID).Error; err != nil {
		return nil, errors.New("failed to load user relations")
	}

	result := t.mappedToUserProfile(&user)
	return &result, nil
}

func (t *authServiceImpl) mappedToUserProfile(user *models.UserModel) responsedto.UserProfileData {
	result := responsedto.UserProfileData{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		RoleName: user.Role.Name,
	}

	if user.Organization != nil {
		result.Organization = &responsedto.OrganizationResponse{
			ID:        user.Organization.ID,
			Name:      user.Organization.Name,
			OwnerID:   user.Organization.OwnerID,
			CreatedAt: user.Organization.CreatedAt,
			UpdatedAt: user.Organization.UpdatedAt,
		}
	}

	return result

}

func (t *authServiceImpl) RefreshToken(tokenString string) (string, error) {
	return t.jwtService.RefreshToken(tokenString)
}
