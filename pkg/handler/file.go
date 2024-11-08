package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary UploadFile
// @Security ApiKeyAuth
// @Tags file
// @Description Upload File
// @ID upload-file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file to upload"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/file/upload [post]
func (h *Handler) uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	uuidFilename := uuid.New().String()

	fileExtension := filepath.Ext(file.Filename)

	fmt.Println("fileExtension: ", fileExtension)

	savePath := filepath.Join("./static/uploads", uuidFilename+fileExtension)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": filepath.Join("/static", uuidFilename+fileExtension)})
}
