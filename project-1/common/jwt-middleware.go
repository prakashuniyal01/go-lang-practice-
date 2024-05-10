package common

import (
	"fmt"
	"net/http"
	"trademarkia/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(user models.UserData, c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usermail": user.Email,
	})
	tokenString, err := token.SignedString([]byte("qwertyacid12345acidqwerty"))

	if err == nil {
		fmt.Println("token created")
	}
	c.Set("userEmail", user.Email)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorise", tokenString, 3600, "", "", false, true)
}

func ValidateCookie(c *gin.Context) {
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		fmt.Println("cookie not found")
		c.JSON(http.StatusBadGateway, gin.H{"error": "user not logged in"})
		c.Abort()
	} else {
		fmt.Println("cookie", cookie)
	}
	c.Next()
}

func RetriveJwtToken(c *gin.Context) {
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		fmt.Println("cookie not found")
		c.JSON(http.StatusBadGateway, gin.H{"error": "user not logged in"})
		c.Abort()
	} else {
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("qwertyacid12345acidqwerty"), nil
		})

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "user authentication failed"})
			c.Abort()
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userEmail := claims["usermail"].(string)
			c.Set("userEmail", userEmail)
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusBadGateway, gin.H{"error": "user authentication failed"})
		}
	}
}

func DeleteCookie(c *gin.Context, cookieName string) {
	c.SetCookie(cookieName, "", -1, "/", "", false, true)
}
