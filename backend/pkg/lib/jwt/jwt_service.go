package jwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	OrganizationId *uint `json:"organization_id"`
	jwt.RegisteredClaims
}

type JwtService interface{
	GenerateToken(userID uint, email string, roleID uint, OrganizationId *uint) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
	RefreshToken(tokenString string) (string, error) 
	GetUserFromContext(ctx context.Context) (*Claims, bool)
}


const (
	UserContextKey = "user"
)