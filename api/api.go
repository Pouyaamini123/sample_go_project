package api

import (
	"sample_go_project/api/routers"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.LoadHTMLGlob("templates/*")

	v1 := r.Group("/mainpage")
	{
		routers.Rutermainpage(v1)
	}

	v2 := r.Group("/startpage")
	{
		routers.Ruterstartpage(v2)
	}

	r.Run(":5000")
}
