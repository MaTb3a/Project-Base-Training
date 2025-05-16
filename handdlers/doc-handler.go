package Handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MaTb3aa/Project-Base-Training/models"
	Services "github.com/MaTb3aa/Project-Base-Training/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentHandler struct {
	service *Services.DocumentService
}

func NewDocumentHandler(service *Services.DocumentService) *DocumentHandler {
	return &DocumentHandler{service: service}
}

// CreateDocument godoc
// @Summary Create a new document
// @Description Create a new document with the input payload
// @Tags documents
// @Accept json
// @Produce json
// @Param document body models.Document true "Create document"
// @Success 201 {object} models.Document
// @Failure 400 {object} models.SwaggerErrorResponse
// @Failure 500 {object} models.SwaggerErrorResponse
// @Router /documents/ [post]
func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var doc models.Document

	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDoc(&doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create document"})
		return
	}

	c.JSON(http.StatusCreated, doc)
}

// GetAllDocuments godoc
// @Summary Get all documents
// @Description Get all documents
// @Tags documents
// @Produce json
// @Success 200 {array} models.Document
// @Failure 500 {object} models.SwaggerErrorResponse
// @Router /documents/ [get]
func (h *DocumentHandler) GetAllDocuments(c *gin.Context) {
	docs, err := h.service.GetAllDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch documents"})
		return
	}
	c.JSON(http.StatusOK, docs)
}

// GetDocumentByID godoc
// @Summary Get document by ID
// @Description Get document by ID
// @Tags documents
// @Produce json
// @Param id path int true "Document ID"
// @Success 200 {object} models.Document
// @Failure 400 {object} models.SwaggerErrorResponse
// @Failure 404 {object} models.SwaggerErrorResponse
// @Router /documents/{id} [get]
func (h *DocumentHandler) GetDocumentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	doc, err := h.service.GetDocumentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

// UpdateDocument godoc
// @Summary Update a document
// @Description Update a document with the input payload
// @Tags documents
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Param document body models.Document true "Update document"
// @Success 200 {object} models.SwaggerSuccessResponse
// @Failure 400 {object} models.SwaggerErrorResponse
// @Failure 500 {object} models.SwaggerErrorResponse
// @Router /documents/{id} [put]
func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	var doc models.Document
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	if err = c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err = h.service.UpdateDocument(&doc, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update the document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

// DeleteDocument godoc
// @Summary Delete a document
// @Description Delete a document by ID
// @Tags documents
// @Produce json
// @Param id path int true "Document ID"
// @Success 200 {object} models.SwaggerSuccessResponse
// @Failure 400 {object} models.SwaggerErrorResponse
// @Failure 404 {object} models.SwaggerErrorResponse
// @Failure 500 {object} models.SwaggerErrorResponse
// @Router /documents/{id} [delete]
func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = h.service.DeleteDocument(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the document"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
