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
	fmt.Println("Attempting uploadResource in controllers")
	var res types.Resource

	/* need to test to see if the actual db entry gets created and entered at some point
	if err := g.BindJSON(&res); err != nil {
		fmt.Println("binding JSON")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	fmt.Println(res)
	*/
	// IT CURRENTLY SHOULD UPL;OAD THE FILE JUST NEED TO FIGURE OUT HOW TO GET THE PATH
	//categoryMap := m.GetCompanyDocumentResourceCategoriesTable().Load(res.CategoryID)
	//res.URL := "/static/" + reflect.ValueOf(categoryMap).FieldByName("CategoryFolderPath").String() + "/"
	path := m.GetConf().Get("server.website") + "/static/testfolder/"
	// Check if folder exists.
	fmt.Println("Checking Dir")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Making New DIr")
		os.MkdirAll(path, os.ModeDir)
	}

	// Multipart form.

	fmt.Println("Checking Multipart form")
	form, err := g.MultipartForm()
	if err != nil {
		fmt.Println("Error in Multipart form")

		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	for k := range form.File {
		fmt.Println(k)
	}

	files := form.File["fileList"]

	for _, file := range files {
		fmt.Println(file.Filename)
		// Save file to filesystem.
		if err := g.SaveUploadedFile(file, path+file.Filename); err != nil {
			fmt.Println("Error saving file to system")
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
	}
	/*
		file, err := g.FormFile("file")
		if err != nil {
			fmt.Println("Error in Multipart form")

			g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		if err := g.SaveUploadedFile(file, path+file.Filename); err != nil {
			//m.DeleteAttachedFile(data.FileID)
			fmt.Println("Error saving file to system")
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		fmt.Println("Retrieving formFiles")
		//files := form.File["file"]
	*/
	fmt.Println("Reached the end")
	// ok right now im looking into how multipart will send the form data and the form file
	//
	g.JSON(http.StatusOK, res)
}
