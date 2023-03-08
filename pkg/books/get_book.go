package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricky7xc-dev/read-env-var/pkg/common/models"
)

func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
