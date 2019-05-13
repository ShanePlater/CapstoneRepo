package main

import (
	"github.com/gin-gonic/gin"
)

var mode string

func init() {
	gin.SetMode(mode)
}
