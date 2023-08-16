package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24

var Secret = []byte("secret")

func GenToken(id int64, username string) (string, error) {
	c := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "clipboard",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		claims, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		c.Set("username", claims.ID)
		c.Next()
	}
}
