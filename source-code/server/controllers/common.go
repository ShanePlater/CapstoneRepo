package controllers

import (
	"server/models"
	"sync"
)

// Context is the data structure utilized for controllers.
type Context struct {
	model   *models.Context
	handler *sync.Map
}
