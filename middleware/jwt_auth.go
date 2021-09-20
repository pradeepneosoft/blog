package middleware

import (
	"blog/api/service"
	"blog/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			util.ErrorJSON(c, http.StatusBadRequest, "NO token found")
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims
			log.Println(claims)

		} else {
			log.Println(err)
			util.ErrorJSON(c, http.StatusBadRequest, "token is not valid")
			return
		}
	}
}
