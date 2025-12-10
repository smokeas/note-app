package auth

import (
 "time"

 "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
 UserID int `json:"user_id"`
 jwt.RegisteredClaims
}

func GenerateToken(userID int, secret string, ttl time.Duration) (string, error) {
 claims := Claims{
  UserID: userID,
  RegisteredClaims: jwt.RegisteredClaims{
   ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
   IssuedAt:  jwt.NewNumericDate(time.Now()),
   Subject:   "auth",
  },
 }
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 return token.SignedString([]byte(secret))
}

// ParseToken returns claims or error
func ParseToken(tokenStr, secret string) (*Claims, error) {
 token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
  return []byte(secret), nil
 })
 if err != nil {
  return nil, err
 }
 if claims, ok := token.Claims.(*Claims); ok && token.Valid {
  return claims, nil
 }
 return nil, jwt.ErrTokenInvalidClaims
}

