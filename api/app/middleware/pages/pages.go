package pages

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Values struct {
	Offset uint
	Limit  uint
}

// TODO configure min/max values
func Paginate(base Values) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
				c.JSON(http.StatusBadRequest, gin.H{})
				c.Abort()
			}
		}()

		values := base

		if param, ok := c.GetQuery("limit"); ok {
			v, err := strconv.ParseUint(param, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"description": "Limit invalid"})
				c.Abort()
				return
			}
			values.Limit = uint(v)
		}

		if param, ok := c.GetQuery("offset"); ok {
			v, err := strconv.ParseUint(param, 10, 32)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"description": "Offset invalid"})
				c.Abort()
				return
			}
			values.Offset = uint(v)
		}

		if values.Limit == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"description": "Limit invalid"})
			c.Abort()
			return
		}

		c.Set("pages", values)
	}
}
