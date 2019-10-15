package controllers

import (
	"fmt"
	"net/http"
	"os"
	"server/models"
	"server/types"
	"github.com/gin-gonic/gin"
)

func uploadResource(g *gin.Context, m *models.Context) {
	//var data types.Resource

	path := m.GetConf().Get("server.website") + "/static/testfolder/"

	// Check if folder exists.
	fmt.Println("Checking Dir")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Making New DIr")
		os.MkdirAll(path, os.ModeDir)
	}

	fmt.Println("Checking Multipart form")
	_, err := g.MultipartForm()
	if err != nil {
		fmt.Println("Error in Multipart form")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	_, header, err := g.Request.FormFile("file")

	if err != nil {
		fmt.Println("Error in osFileSave")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	fmt.Println(header)

	if err := g.SaveUploadedFile(header, path+header.Filename); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}	



	fmt.Println("Reached the end")
	g.JSON(http.StatusOK, header.Filename)
}
