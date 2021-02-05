package user

import (
	"fmt"
	"net/http"

	contracts "franigen-example/api/core/contracts/user"
	"franigen-example/api/core/usecases/errors"
	usecases "franigen-example/api/core/usecases/user"

	"github.com/gin-gonic/gin"
)

type Create struct {
	UseCase usecases.Create
}

func (h *Create) Handle(c *gin.Context) {

	req, err := h.bindRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "No se pudo leer la solicitud: " + err.Error()})
		return
	}

	err = req.Check()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": fmt.Sprintf("Información errónea o faltante: %s", err)})
		return
	}

	res, err := h.UseCase.Execute(c, req)
	if err != nil {
		var status int
		var message string

		switch err {
		case errors.ErrCreatingUser:
			status, message = http.StatusInternalServerError, "No se pudo crear al usuario"
		default:
			status, message = http.StatusInternalServerError, "Algo salió mal"
		}

		c.JSON(status, gin.H{"description": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": map[string]interface{}{"user": res}})
}

func (Create) bindRequest(c *gin.Context) (contracts.CreateUserRequest, error) {
	var req contracts.CreateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}
	return req, nil
}
