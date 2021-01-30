package user

import (
	"errors"
	"franigen-example/api/app/middleware/pages"
	contracts "franigen-example/api/core/contracts/user"
	ucerrors "franigen-example/api/core/usecases/errors"
	usecases "franigen-example/api/core/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Select struct {
	UseCase usecases.Select
}

func (h *Select) Handle(c *gin.Context) {

	pagination := c.MustGet("pages").(pages.Values)

	req, err := h.bindRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "No se pudo leer la solicitud: " + err.Error()})
		return
	}

	req.Limit, req.Offset = pagination.Limit, pagination.Offset

	res, err := h.UseCase.Execute(c, req)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, ucerrors.ErrGettingUsers):
			status, message = http.StatusInternalServerError, "No se pudieron obtener los usuarios"
		default:
			status, message = http.StatusInternalServerError, "Algo salió mal"
		}

		c.JSON(status, gin.H{"description": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": res})
}

func (Select) bindRequest(c *gin.Context) (contracts.SelectUsersRequest, error) {
	var req contracts.SelectUsersRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		return req, err
	}
	return req, nil
}
