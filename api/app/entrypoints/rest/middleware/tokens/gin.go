package tokens

import (
	"github.com/gin-gonic/gin"
)

func ValidateToken(t Toker, env string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("id", 1)
		/*defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
				c.JSON(http.StatusUnauthorized, gin.H{})
				c.Abort()
			}
		}()

		tokenString := c.Request.Header.Get("Authorization")

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		token, err := t.Get(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		if token.Env != env {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		c.Set("id", token.Id)
		c.Set("name", token.Name)
		c.Set("email", token.Email)
		*/
	}
}
