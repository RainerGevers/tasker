package middleware

import (
	"context"
	"fmt"
	"github.com/RainerGevers/tasker/config"
	"github.com/RainerGevers/tasker/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

type authenticationClaims struct {
	Uuid string `json:"uuid"`
	jwt.StandardClaims
}

type AuthenticationMiddleware struct {
	Env *config.Env
}

func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionJwt := r.Header.Get("Authorization")

		token, err := jwt.ParseWithClaims(sessionJwt, &authenticationClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(*authenticationClaims); ok && token.Valid {
			user := models.User{}
			amw.Env.Database.Model(&models.User{}).Select("users.uuid").Joins("INNER JOIN sessions ON users.id = sessions.user_id").Where("sessions.uuid = ? AND sessions.expires_at > ?", claims.Uuid, time.Now()).Scan(&user)
			r = r.WithContext(context.WithValue(r.Context(), "current_user_uuid", user.Uuid))
			next.ServeHTTP(w, r)
		} else {
			fmt.Println(err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

	})
}