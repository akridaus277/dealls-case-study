package utils

import (
	"authentication-service/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string, roles []models.Role) (string, error) {
	rolesStr := RolesToStrings(roles)
	log.Printf("generatejwt roles : %s", rolesStr)
	claims := jwt.MapClaims{
		"user":  username,
		"roles": rolesStr,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RolesToStrings(roles []models.Role) []string {
	result := make([]string, len(roles))
	for i, r := range roles {
		result[i] = r.Name
	}
	return result
}
