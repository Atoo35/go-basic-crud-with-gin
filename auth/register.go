package auth

import (
	"fmt"
	"net/http"

	"github.com/Atoo35/basic-crud/models"
	"github.com/Atoo35/basic-crud/schema"
	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) Register(ctx *gin.Context) {
	body := schema.CreateUser{}

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

	hashedPassword, hashPwdError := utils.HashPassword(body.Password)
	if hashPwdError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			utils.BuildResponse_(
				utils.UnknownError.GetResponseStatus(),
				"Error while hashing password", hashPwdError.Error(),
			),
		)
		return
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Username = body.Username
	user.Password = hashedPassword

	if result := h.DB.Create(&user); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			utils.BuildResponse_(
				utils.DatabaseError.GetResponseStatus(),
				"Error while creating user",
				result.Error.Error(),
			),
		)
		return
	}

	userResponse := schema.CreateUserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
	ctx.JSON(http.StatusOK, utils.BuildResponse(utils.Success, userResponse))

}
