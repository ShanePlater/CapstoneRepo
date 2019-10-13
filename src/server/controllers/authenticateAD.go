package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
	auth "gopkg.in/korylprince/go-ad-auth.v2"
)

//ADAuth - Connects to the LAR domain when an authenticator.
func authenticateAD(g *gin.Context, m *models.Context) {

	var data types.Authenticator
	fmt.Println("controllers/adauth.go at controllers")
	if err := g.BindJSON(&data); err != nil {
		fmt.Println("controllers/adauth.go error when marshelling")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	config := &auth.Config{
		Server:   "LR-BNE-ERP-01",       /* Box that contains domain controller */
		Port:     389,                   /* Default Port for AD connection */
		BaseDN:   "OU=lar,DC=net,DC=au", /* Domain */
		Security: auth.SecurityNone,
	}

	username := data.Username
	password := data.Password

	status, err := auth.Authenticate(config, username, password)

	if err != nil {
		//handle err
		fmt.Println("controllers/adauth.go Handle err")
		fmt.Println(err)
		return
	}

	if !status {
		//handle failed authentication
		fmt.Println("controllers/adauth.go Failed Auth")
		return
	}

	fmt.Println("controllers/adauth.go End of func")
}
