package controllers

import (
	"fmt"
	"net/http"
	"os"
	"server/models"

	"github.com/gin-gonic/gin"
)

func uploadResource(g *gin.Context, m *models.Context) {

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
	if err := g.SaveUploadedFile(header, path+header.Filename); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	/*
		fmt.Println("Attempting uploadResource in controllers")

		type uploadForm struct {
			Files []*multipart.FileHeader `form:"files" binding:"required"`
			Data  string                  `form:"data" binding:"required"`
		}

		var form uploadForm


		if err := g.ShouldBindWith(&form, binding.Form); err != nil {
			fmt.Println("binding form")
			g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		for _, file := range form.Files {
			// Save file to filesystem.
			if err := g.SaveUploadedFile(file, path); err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
		}

		/* need to test to see if the actual db entry gets created and entered at some point
		if err := g.BindJSON(&res); err != nil {
			fmt.Println("binding JSON")
			g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		fmt.Println(res)

		// IT CURRENTLY SHOULD UPL;OAD THE FILE JUST NEED TO FIGURE OUT HOW TO GET THE PATH
		//categoryMap := m.GetCompanyDocumentResourceCategoriesTable().Load(res.CategoryID)
		//res.URL := "/static/" + reflect.ValueOf(categoryMap).FieldByName("CategoryFolderPath").String() + "/"


		// Multipart form.

		g.Request.ParseMultipartForm()
		for key, value := range g.Request.PostForm {
			fmt.Println(key, value)

		}
		formFiles := form.File["files"]
		formFile := form.File["file"]

		for _, file := range formFile {
			// Save file to filesystem.
			if err := g.SaveUploadedFile(file, path); err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}
		}
		/*
			file, fileErr := g.FormFile("file")
			if fileErr != nil {
				g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}

			if err := g.SaveUploadedFile(file, path+file.Filename); err != nil {
				fmt.Println("Error saving file to system")
				g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}

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
	g.JSON(http.StatusOK, nil)
}
