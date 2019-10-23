package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"server/models"
	"server/types"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func uploadResource(g *gin.Context, m *models.Context) {
	//var data types.Resource

	data := new(types.Resource)

	data.FileFriendlyName = g.Request.FormValue("friendlyFileName")
	data.FileRevision = g.Request.FormValue("fileRevision")
	data.AuthorizedBy = g.Request.FormValue("authorizedBy")
	data.AuthorizedDate = g.Request.FormValue("authorizedDate")
	data.CategoryID = utils.Atoi(g.Request.FormValue("categoryID"))
	data.URL = g.Request.FormValue("url")

	fmt.Println(data)
	//var saveFileName = g.Request.FormValue("fileNameNoExt") // if saving file with filename appended to include version

	table, ok := m.GetCompanyDocumentResourceCategoriesTable().Load(utils.Itoa(data.CategoryID))
	if !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": errors.New("file path not found")})
		return
	}

	path := m.GetConf().Get("server.website") + "/static/" + reflect.ValueOf(table).FieldByName("CategoryFolderPath").String() + "/"

	// Check if folder exists.

	_, err := g.MultipartForm()
	if err != nil {
		fmt.Println("Error in Multipart form")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	_, header, err := g.Request.FormFile("file")
	data.FileName = header.Filename
	if err != nil {
		fmt.Println("Error in osFileSave")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	fmt.Println("Saving Entry  to DB")
	if err := m.CreateOrUpdateResource(data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		fmt.Println("controllers/UploadResource.go  there was an error creating the Resource")
		return
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Making New DIr")
		os.MkdirAll(path, os.ModeDir)
	}

	//if err := g.SaveUploadedFile(header, path+saveFileName); err != nil { // if saving file with filename appended to include version
	if err := g.SaveUploadedFile(header, path+header.Filename); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	m.ReloadCache()
	fmt.Println("Reached the end")
	g.JSON(http.StatusOK, data)
}
