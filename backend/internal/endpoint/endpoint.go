package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavatee/children_backend/internal/service"
)

type Endpoint struct {
	Services *service.Service
}

func NewEndpoint(services *service.Service) *Endpoint {
	return &Endpoint{
		Services: services,
	}
}

func (e *Endpoint) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		ctx.Writer.Header().Set("Access-Control-Allow-Creditionals", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
	})
	api := router.Group("/api")
	{
		api.POST("/codes", e.SendCode)
		api.GET("/children", e.GetChildren)
		api.PUT("/children/:id", e.TakeChild)
		api.POST("/admin/children", e.GetChildrenInfo)
		api.POST("/admin", e.NewAdmin)
	}
	return router
}
