package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Allifiando/go-gin-boilerplate/helper"
	"github.com/Allifiando/go-gin-boilerplate/model"
	responses "github.com/Allifiando/go-gin-boilerplate/model/response"
	Error "github.com/Allifiando/go-gin-boilerplate/pkg/error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Credential ...
type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Auth ...
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorParams := map[string]interface{}{}
		statusCode := 401
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			errorParams["meta"] = map[string]interface{}{
				"status":  statusCode,
				"message": "Unauthorized",
			}
			errorParams["code"] = statusCode
			c.JSON(http.StatusUnauthorized, helper.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}
		runes := []rune(header)
		tokenString := string(runes[7:])
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		fmt.Println(token)
		fmt.Println(claims)
		if err != nil {
			c.Abort()
			return
		}
		if !token.Valid {
			c.Abort()
			return
		}

		uid := claims["uuid"].(string)
		userModel := model.UserModel{}
		userData, err := userModel.FindByUUID(uid)

		if userData.UUID != "" {
			errorParams["meta"] = map[string]interface{}{
				"status":  statusCode,
				"message": "Unauthorized",
			}
			errorParams["code"] = statusCode
			c.JSON(http.StatusUnauthorized, helper.OutputAPIResponseWithPayload(errorParams))
			c.Abort()
			return
		}

		claims["user_uuid"] = userData.UUID
		c.Set("User", claims)
		c.Next()
	}
}

// CreateToken ...
func CreateToken(data responses.UserModel) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uid"] = data.UUID
	atClaims["data"] = data
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		Error.Error(err)
		return "", err
	}
	return token, nil
}
