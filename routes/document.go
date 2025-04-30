package routes

import (
	"github.com/MaTb3aa/Project-Base-Training/controller"
	"github.com/gin-gonic/gin"
)

func DocumentRoute(router *gin.Engine) {
	router.GET("/documents", controller.GetDocuments)
	router.POST("/document", controller.CreateDocument)
	router.GET("/document/:id", controller.GetDocument)
	router.PUT("/document/:id", controller.UpdateDocument)
	router.DELETE("/document/:id", controller.DeleteDocument)
}

