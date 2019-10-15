package main

import (
	"server/config"
	"server/controllers"
	"server/models"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read in the configuration file.
	c := config.New(getDir())

	// Initialize router.
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
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

/*
//creates a new engine that allows the use of auth
func engine() *gin.Engine {
	r := gin.New()


	/* Leaving this out for testing purposes
	r.POST("/api/v1/post/login", login)
	r.GET("/api/v1/get/logout", logout)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}

	return r
}
*/
