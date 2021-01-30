package paraval

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{})
				c.Abort()
			}
		}()

		params := make([]Param, 0, len(c.Params))
		for _, p := range c.Params {
			params = append(params, Param{
				Key:   p.Key,
				Value: p.Value,
			})
		}

		validated, err := Validate(params)
		if err != nil {
			var status int
			var msg string
			switch err {
			case ErrMissingParam:
				status, msg = http.StatusBadRequest, "missing url param"
			case ErrInvalidParam:
				status, msg = http.StatusBadRequest, "invalid url param"
			default:
				status, msg = http.StatusInternalServerError, "Unknown error"
			}
			c.JSON(status, gin.H{"description": msg})
			c.Abort()
			return
		}

		c.Set("validated", validated)
	}
}
