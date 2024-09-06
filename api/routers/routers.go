package routers

import (
	"sample_go_project/api/handlers"

	"github.com/gin-gonic/gin"
)

func Rutermainpage(r *gin.RouterGroup) {
	handler := handlers.NewMainpageHandler()
	r.GET("", handler.Mainpage_GET)
	r.POST("", handler.Mainpage_POST)
}

func Ruterstartpage(r *gin.RouterGroup) {
	handler := handlers.NewStartPage()
	r.GET("", handler.StartPage_GET)
}
