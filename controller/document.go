package controller

import (
	"net/http"
	"github.com/MaTb3aa/Project-Base-Training/config"
	"github.com/MaTb3aa/Project-Base-Training/models"
	"github.com/gin-gonic/gin"
)

func GetDocuments(c *gin.Context) {
	documents := []models.Document{}
	config.DB.Find(&documents)
	SuccessResponse(http.StatusOK, c, documents)
}

func CreateDocument(c *gin.Context) {
	document := models.Document{}
	if err := c.ShouldBindJSON(&document); err != nil {
		RespondWithStatusBadRequest(c, err)
		return
	}
	config.DB.Create(&document)
	SuccessResponse(http.StatusCreated, c, document)
}

func GetDocument(c *gin.Context) {
	id := c.Param("id")
	document := models.Document{}
	if err := config.DB.First(&document, id).Error; err != nil {
		RespondWithStatusNotFound(c, err)
		return
	}
	SuccessResponse(http.StatusOK, c, document)
}

func UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	document := models.Document{}
	if err := config.DB.First(&document, id).Error; err != nil {
		RespondWithStatusNotFound(c, err)
		return
	}
	if err := c.ShouldBindJSON(&document); err != nil {
		RespondWithStatusBadRequest(c, err)
		return
	}
	config.DB.Save(&document)
	SuccessResponse(http.StatusOK, c, document)
}

func DeleteDocument(c *gin.Context) {
	id := c.Param("id")
	document := models.Document{}
	if err := config.DB.First(&document, id).Error; err != nil {
		RespondWithStatusNotFound(c, err)
		return
	}
	config.DB.Delete(&document)
	SuccessResponse(http.StatusNoContent, c, nil)
}

func RespondWithStatusNotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{"message": "failed", "error": err.Error()})
}

func RespondWithStatusBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "failed", "error": err.Error()})
}

func SuccessResponse(status int, c *gin.Context, data interface{}) {
	c.JSON(status, gin.H{"message": "success", "data": data})
}
