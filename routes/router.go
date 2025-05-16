package routes

import (
	"net/http"

	Handlers "github.com/MaTb3aa/Project-Base-Training/handdlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handlers.DocumentHandler) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	documents := r.Group("/documents")
	{
		documents.POST("/", handler.CreateDocument)
		documents.GET("/", handler.GetAllDocuments)
		documents.GET("/:id", handler.GetDocumentByID)
		documents.PUT("/:id", handler.UpdateDocument)
		documents.DELETE("/:id", handler.DeleteDocument)
	}

	return r
}
