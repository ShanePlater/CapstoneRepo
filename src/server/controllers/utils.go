package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"
	"time"

	"github.com/gin-gonic/gin"
)

// parseTime parse and complete start time and end time.
// It will check whether start and end string are in RFC3339 format.
// If start is empty string, it will parse start as "0001-01-01T00:00:00Z".
// If end is empty string, it will parse start as current local time in RFC3339 format.
func parseTime(start, end *string) error {
	// Check start time.
	t, err := time.Parse(time.RFC3339, *start)
	if *start != "" && err != nil {
		return err
	}

	// Format start time.
	*start = t.Format(time.RFC3339)

	// Check end time.
	t, err = time.Parse(time.RFC3339, *end)
	if *end != "" && err != nil {
		return err
	}

	// Parse end as current local time in RFC3339 format if it is empty string.
	if *end == "" {
		*end = time.Now().Format(time.RFC3339)
		return nil
	}

	// Format end time.
	*end = t.Format(time.RFC3339)

	return nil
}

func swapCodes(g *gin.Context, m *models.Context) {
	var data types.NameForCodes

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Situation that CategoryID is not set.
	if data.ProjectLocationCode == "" {
		fmt.Println("controllers/getProject.go  missing Proj Loc code")
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss CategoryID"})
		return
	}
	/* 	//if getting the data was good
	   	if res, ok := m.SwapProjLocCode(&data); ok {
	   		// Serve the result.
	   		g.JSON(http.StatusOK, res)
	   		return
	   	} */

	// Serve the result.
	// -1 is meant to ignore number of document records (get as more as possible).
	g.JSON(http.StatusOK, data)
}
