package user

import (
	"errors"
	"franigen-example/api/app/entrypoints/rest/middleware/paraval"
	ucerrors "franigen-example/api/core/usecases/errors"
	usecases "franigen-example/api/core/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Single struct {
	UseCase usecases.Single
}

func (h *Single) Handle(c *gin.Context) {

	params := c.MustGet("validated").(paraval.Validated)
	userID := params["userID"]

	res, err := h.UseCase.Execute(c, userID)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, ucerrors.ErrGettingUser):
			status, message = http.StatusInternalServerError, "No se pudo obtener al usuario"
		default:
			status, message = http.StatusInternalServerError, "Algo salió mal"
		}

		c.JSON(status, gin.H{"description": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": map[string]interface{}{"user": res}})
}
