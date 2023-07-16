package auth

import (
	"fmt"
	"net/http"

	"github.com/Atoo35/basic-crud/models"
	"github.com/Atoo35/basic-crud/schema"
	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) Login(ctx *gin.Context) {
	body := schema.Login{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println("error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			utils.BuildResponse_(
				utils.BadRequest.GetResponseStatus(),
				utils.BadRequest.GetResponseMessage(),
				err.Error(),
			),
		)
		return
	}

	var user models.User
	if result := h.DB.Where("username = ?", body.Username).First(&user); result.Error != nil {
		utils.UnauthorisedResponse(ctx)
		return
	}

	if err := utils.VerifyPassword(user.Password, body.Password); err != nil {
		utils.UnauthorisedResponse(ctx)
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			utils.BuildResponse_(
				utils.UnknownError.GetResponseStatus(),
				"Error while generating token",
				err.Error(),
			),
		)
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponse(utils.Success, gin.H{
		"token": token,
	}))
}
