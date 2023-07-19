package books

import (
	"net/http"

	"github.com/Atoo35/basic-crud/models"
	"github.com/Atoo35/basic-crud/schema"
	"github.com/Atoo35/basic-crud/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBook(ctx *gin.Context) {
	body := schema.CreateBook{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			utils.BuildResponse_(
				utils.BadRequest.GetResponseStatus(),
				utils.BadRequest.GetResponseMessage(),
				err.Error(),
			),
		)
		return
	}

	var book models.Book

	book.Name = body.Name
	book.Author = body.Author
	book.Publication = body.Publication

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			utils.BuildResponse_(
				utils.DatabaseError.GetResponseStatus(),
				utils.DatabaseError.GetResponseMessage(),
				result.Error.Error(),
			),
		)
		return
	}
	FormatBookResponse(book, &body)
	ctx.JSON(http.StatusCreated, utils.BuildResponse(utils.Success, &body))
}
