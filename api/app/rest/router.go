package rest

import (
	"franigen-example/config"
	"net/http"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {

	router := gin.New()
	router.MaxMultipartMemory = (2 << 20) * 8 // 16mb

	// Must be first
	router.GET("/", gcpHealthCheck)

	router.Use(gin.Logger())
	router.Use(nice.Recovery(recoveryHandler))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET,PUT,POST,DELETE,PATCH"},
		AllowHeaders:    []string{"accept,x-access-token,content-type,authorization"},
	}))

	return router
}

// GKE ingress health check
// More info: https://github.com/kubernetes/kubernetes/issues/20555
func gcpHealthCheck(c *gin.Context) { c.Status(http.StatusOK) }

func recoveryHandler(c *gin.Context, err interface{}) {
	detail := ""
	if e, ok := err.(error); ok {
		if config.Get().Env == "develop" {
			detail = e.Error()
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"description": detail})
}
