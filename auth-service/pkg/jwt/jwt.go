package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
    secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
    return &JWTManager{secretKey: secretKey}
}

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func (j *JWTManager) GenerateToken(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.secretKey))
}
