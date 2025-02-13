package controllers

import (
	"net/http"
	"os"
	"server/models"
	"server/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func uploadFile(g *gin.Context, m *models.Context) {
	var res []types.File
	path := m.GetConf().Get("server.website") + "/pub/HelpDesk/"

	// Check if folder exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeDir)
	}

	// Multipart form.
	form, err := g.MultipartForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	files := form.File["files"]

	for _, file := range files {
		data := types.File{FileName: file.Filename}

		// Create record.
		if err := m.CreateAttachedFile(&data); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		res = append(res, data)

		// Save file to filesystem.
		if err := g.SaveUploadedFile(file, path+strconv.Itoa(data.FileID)); err != nil {
			m.DeleteAttachedFile(data.FileID)
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
	}

	g.JSON(http.StatusOK, res)
}
