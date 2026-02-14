package jwt

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtServiceImp struct {
	jwtSecret string
}

// GetUserFromContext implements JwtService.
func (t *jwtServiceImp) GetUserFromContext(ctx context.Context) (*Claims, bool) {
	user, ok := ctx.Value(UserContextKey).(*Claims)
	return user, ok
}

// GenerateToken implements JwtService.
func (t *jwtServiceImp) GenerateToken(userID uint, email string, roleID uint, OrganizationId *uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RoleID: roleID,
		OrganizationId: OrganizationId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// RefreshToken implements JwtService.
func (t *jwtServiceImp) RefreshToken(tokenString string) (string, error) {
	claims, err := t.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	return t.GenerateToken(claims.UserID, claims.Email, claims.RoleID, claims.OrganizationId)
}

// ValidateToken implements JwtService.
func (t *jwtServiceImp) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(t.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func InstanceJwtService() JwtService {
	secret := os.Getenv("JWT_SECRET")
	return &jwtServiceImp{
		jwtSecret: secret,
	}
}
