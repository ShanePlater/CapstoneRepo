package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"server/models"
	"server/types"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getMe(g *gin.Context, m *models.Context) {
	var data types.GetMeJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if keyword is empty.
	if data.Username == "" || data.Password == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss username or password."})
		return
	}

	if err := authCheck(data.Username, data.Password); err != nil {
		g.JSON(http.StatusForbidden, err)
		return
	}

	token := md5.New()
	token.Write([]byte(data.Username + strconv.FormatInt(time.Now().Unix(), 10) + data.Password))

	// Serve the result.
	g.JSON(http.StatusOK, gin.H{"Token": hex.EncodeToString(token.Sum(nil))})
}
