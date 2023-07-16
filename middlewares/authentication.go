package middlewares

import (
	"net/http"
	"strings"

	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		authorizationHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.BuildResponse_(
					utils.Unauthorized.GetResponseStatus(),
					utils.Unauthorized.GetResponseMessage(),
					"Invalid token"),
			)
			return
		}
		user, err := utils.VerifyJWT(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				utils.BuildResponse_(
					utils.Unauthorized.GetResponseStatus(),
					utils.Unauthorized.GetResponseMessage(),
					"Token expired"),
			)
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
