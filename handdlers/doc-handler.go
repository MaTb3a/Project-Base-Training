package Handlers

import (
	"net/http"
	"strconv"

	"github.com/MaTb3aa/Project-Base-Training/models"
	Services "github.com/MaTb3aa/Project-Base-Training/services"
	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	service *Services.DocumentService
}

func NewDocumentHandler(service *Services.DocumentService) *DocumentHandler {
	return &DocumentHandler{service: service}
}

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
func (h *DocumentHandler) GetAllDocuments(c *gin.Context) {
	docs, err := h.service.GetAllDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch documents"})
		return
	}
	c.JSON(http.StatusOK, docs)
}

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

func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	var doc models.Document
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateDocument(&doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update document"})
		return
	}
	c.JSON(http.StatusOK, doc)
}
