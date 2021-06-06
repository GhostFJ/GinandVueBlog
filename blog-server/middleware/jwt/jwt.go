package jwt

import (
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		appG := app.Gin{C: c}
		code = e.SUCCESS
		tokenString := c.GetHeader("Authorization") //在header中存放token

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			code = e.INVALID_PARAMS
		} else {
			Authorization := tokenString[7:]
			claims, err := util.ParseToken(Authorization)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
			c.Set("claims", claims)
		}
		if code != e.SUCCESS {
			appG.Response(http.StatusUnauthorized, code, map[string]interface{}{
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
