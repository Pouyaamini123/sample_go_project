package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainpageHandler struct {
}

func NewMainpageHandler() *MainpageHandler {
	return &MainpageHandler{}
}

func (h *MainpageHandler) Mainpage_GET(c *gin.Context) {
	//c.JSON(http.StatusOK, "working")
	c.HTML(http.StatusOK, "mainpage.tmpl", gin.H{
		"title":       "خوش امدید",
		"description": "این یک سایت ساخت جدول زمان بندی هفتگی است که میتواند در موارد مختلفی مثل برنامه ریزی کاری،برنامه ریزی درسی،برنامه ریزی امتحانات و یا برنامه ریزی دروس طول ترم استفاده شود.",
	})
	return
}

func (h *MainpageHandler) Mainpage_POST(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, "Method Not Allowed")
	return
}
