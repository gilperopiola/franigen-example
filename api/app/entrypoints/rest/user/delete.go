package user

import (
	"errors"
	"franigen-example/api/app/entrypoints/rest/middleware/paraval"
	ucerrors "franigen-example/api/core/usecases/errors"
	usecases "franigen-example/api/core/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Delete struct {
	UseCase usecases.Delete
}

func (h *Delete) Handle(c *gin.Context) {

	params := c.MustGet("validated").(paraval.Validated)
	userID := params["userID"]

	err := h.UseCase.Execute(c, userID)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, ucerrors.ErrDeletingUser):
			status, message = http.StatusInternalServerError, "No se pudo eliminar al usuario"
		default:
			status, message = http.StatusInternalServerError, "Algo salió mal"
		}

		c.JSON(status, gin.H{"description": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": ""})
}
