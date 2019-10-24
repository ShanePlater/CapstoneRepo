package main

import (
	"server/config"
	"server/controllers"
	"server/models"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read in the configuration file.
	c := config.New(getDir())

	// Initialize router.
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	if mode == "debug" {
		r.Use(func(c *gin.Context) {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "content-type")
			c.Header("Access-Control-Allow-Methods", "POST GET")
			c.Next()
		})
		r.OPTIONS("/api/v1/post/:name", func(c *gin.Context) {
			c.Next()
		})
		r.OPTIONS("/api/v1/get/:name", func(c *gin.Context) {
			c.Next()
		})
	}

	// Initialize controller.
	ctl := controllers.New(models.New(c))

	// Serve API.
	r.POST("/api/v1/post/:name", func(g *gin.Context) {
		controllers.Serve(ctl, g)
	})
	r.GET("/api/v1/get/:name", func(g *gin.Context) {
		controllers.Serve(ctl, g)
	})

	// Serve website page and other static files.
	r.StaticFile("/", c.Get("server.website")+"/index.html")
	r.Static("/static", c.Get("server.website")+"/static")
	r.Static("/pub", c.Get("server.website")+"/pub")

	// Run backend server.
	r.Run(":" + c.Get("server.port"))
}
