package user

import (
	"errors"
	"fmt"
	"franigen-example/api/app/entrypoints/rest/middleware/paraval"
	contracts "franigen-example/api/core/contracts/user"
	ucerrors "franigen-example/api/core/usecases/errors"
	usecases "franigen-example/api/core/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Update struct {
	UseCase usecases.Update
}

func (h *Update) Handle(c *gin.Context) {

	params := c.MustGet("validated").(paraval.Validated)
	userID := params["userID"]

	req, err := h.bindRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "No se pudo leer la solicitud: " + err.Error()})
		return
	}

	req.ID = userID

	err = req.Check()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": fmt.Sprintf("Información errónea o faltante: %s", err)})
		return
	}

	res, err := h.UseCase.Execute(c, req)
	if err != nil {
		var status int
		var message string

		switch {
		case errors.Is(err, ucerrors.ErrUpdatingUser):
			status, message = http.StatusInternalServerError, "No se pudo actualizar al usuario"
		default:
			status, message = http.StatusInternalServerError, "Algo salió mal"
		}

		c.JSON(status, gin.H{"description": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": res})
}

func (Update) bindRequest(c *gin.Context) (contracts.UpdateUserRequest, error) {
	var req contracts.UpdateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}
	return req, nil
}
