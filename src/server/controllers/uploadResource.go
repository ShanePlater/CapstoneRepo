package controllers

import (
	"net/http"
	"os"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func uploadResource(g *gin.Context, m *models.Context) {
	var res types.Resource
	// need to test to see if the actual db entry gets created and entered at some point
	if err := g.BindJSON(&res); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// IT CURRENTLY SHOULD UPL;OAD THE FILE JUST NEED TO FIGURE OUT HOW TO GET THE PATH
	//categoryMap := m.GetCompanyDocumentResourceCategoriesTable().Load(res.CategoryID)
	//res.URL := "/static/" + reflect.ValueOf(categoryMap).FieldByName("CategoryFolderPath").String() + "/"
	path := ""
	// Check if folder exists.
	if _, err := os.Stat(res.URL); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeDir)
	}

	// Multipart form.
	form, err := g.MultipartForm()
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	files := form.File["files"]

	if err := g.SaveUploadedFile(files[0], path+files[0].Filename); err != nil {
		//m.DeleteAttachedFile(data.FileID)
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	// ok right now im looking into how multipart will send the form data and the form file
	//
	g.JSON(http.StatusOK, res)
}
