package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	Entities "../entities"
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	secret string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type contextKey string

func (c contextKey) String() string {
	return "mypackage context key " + string(c)
}

const ContextKeyAuthtoken = contextKey("auth-token")

func (a *Auth) NewCookie(user Entities.User) http.Cookie {
	expireTime := time.Now().Add(time.Hour * 1)
	c := Claims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "localhost!",
		}}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(a.secret))

	cookie := http.Cookie{
		Name:     "Auth",
		Value:    token,
		Expires:  expireTime,
		HttpOnly: true}
	return cookie
}

func (a *Auth) Validate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("Auth")
		if err != nil {
			Error(res, http.StatusUnauthorized, "No authorization cookie")
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte(a.secret), nil
		})

		if err != nil {
			Error(res, http.StatusUnauthorized, "Invalid token")
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), ContextKeyAuthtoken, *claims)
			next(res, req.WithContext(ctx))
		} else {
			Error(res, http.StatusUnauthorized, "Unauthorized")
			return
		}
	})
}
