package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/riteshco/Feasto/pkg/types"

	"github.com/golang-jwt/jwt/v5"
)



func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenString string
			authHeader := r.Header.Get("Authorization")
			if strings.Trim(authHeader, " \n") != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = authHeader[7:]
			}
			if tokenString == "" {
				cookie, err := r.Cookie("auth_token")
				if err != nil {
					http.Error(w, "No Token present", http.StatusUnauthorized);
					return
				}
				tokenString = cookie.Value
		}

        secret := os.Getenv("JWT_SECRET")
        token, err := jwt.ParseWithClaims(tokenString, &types.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
            if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return []byte(secret), nil
        })
		fmt.Println(err)
        if err != nil || !token.Valid {
            http.Error(w, "Invalid Token", http.StatusUnauthorized); return
        }


		if claims, ok := token.Claims.(*types.MyClaims); ok {
				ctx := r.Context()
				ctx = context.WithValue(ctx, "id" , claims.ID)
				ctx = context.WithValue(ctx, "username", claims.Username)
				ctx = context.WithValue(ctx, "email" , claims.Email)
				ctx = context.WithValue(ctx, "user_role", claims.UserRole)

				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
				return
		}

		log.Printf("Error processing token claims")
		http.Error(w, "Failed to process token claims", http.StatusInternalServerError)

    })
}
