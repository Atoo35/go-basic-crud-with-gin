package books

import (
	"net/http"

	"github.com/Atoo35/basic-crud/models"
	"github.com/Atoo35/basic-crud/schema"
	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetBook(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			utils.BuildResponse_(
				utils.DataNotFound.GetResponseStatus(),
				utils.DataNotFound.GetResponseMessage(),
				result.Error.Error(),
			),
		)
		return
	}
	var response schema.CreateBook
	FormatBookResponse(book, &response)

	ctx.JSON(http.StatusOK, utils.BuildResponse(utils.Success, response))
}
