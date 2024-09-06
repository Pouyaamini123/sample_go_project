package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StartPage struct {
}

func NewStartPage() *StartPage {
	return &StartPage{}
}

func (h *StartPage) StartPage_GET(c *gin.Context) {
	c.HTML(http.StatusOK, "startpage.tmpl", gin.H{
		"title": "زمان و روز های هفته دروس خود را وارد کنید و سپس دکمه نمایش برنامه را وارد کنید",
	})
	return
}
