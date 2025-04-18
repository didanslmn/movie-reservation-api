package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"slices"

	"github.com/didanslmn/movie-reservation-api/internal/user/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type contextKey string

const userContextKey contextKey = "user"

// middleware untuk otentikasi dengan JWT
func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// cek apakah algoritma valid
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid sub claim"})
			return
		}
		email, _ := claims["email"].(string)
		name, _ := claims["name"].(string)
		roleStr, _ := claims["role"].(string)

		user := &model.User{
			Model: gorm.Model{
				ID: uint(sub),
			},
			Name:  name,
			Email: email,
			Role:  model.Role(roleStr),
		}

		fmt.Println("User from token:", user)

		ctx := context.WithValue(c.Request.Context(), userContextKey, user)
		c.Request = c.Request.WithContext(ctx)
		c.Set("userID", user.ID)

		c.Next()
	}
}

func RoleBasedAccess(allowedRoles ...model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Request.Context().Value(userContextKey).(*model.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		fmt.Println("User in RoleBasedAccess:", user)

		// cek apakah role user termasuk dalam allowedRoles
		if slices.Contains(allowedRoles, user.Role) {
			c.Next()
			return
		}

		// jika role tidak diizinkan
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
	}
}

func GetUserFromContext(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(userContextKey).(*model.User)
	return user, ok
}
