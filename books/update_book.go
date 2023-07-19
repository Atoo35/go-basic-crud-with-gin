package books

import (
	"fmt"
	"net/http"

	"github.com/Atoo35/basic-crud/models"
	"github.com/Atoo35/basic-crud/schema"
	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Params.ByName("id")

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			utils.BuildResponse_(
				utils.BadRequest.GetResponseStatus(),
				utils.BadRequest.GetResponseMessage(),
				err.Error(),
			),
		)
		return
	}
	result := h.DB.Model(&book).Where("id = ?", id).Updates(&book)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			utils.BuildResponse_(
				utils.DatabaseError.GetResponseStatus(),
				utils.DatabaseError.GetResponseMessage(),
				result.Error.Error(),
			),
		)
		return
	}

	if result.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			utils.BuildResponse_(
				utils.DataNotFound.GetResponseStatus(),
				utils.DataNotFound.GetResponseMessage(),
				fmt.Sprintf("Book with id %s not found", id),
			),
		)
		return
	}

	response := schema.CreateBook{}
	FormatBookResponse(book, &response)
	ctx.JSON(http.StatusOK, utils.BuildResponse(utils.Success, &response))

}
