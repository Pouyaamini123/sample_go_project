package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Days []string
	Name string
}

var datas = []data{}

type StartPage struct {
}

func NewStartPage() *StartPage {
	return &StartPage{}
}

func (h *StartPage) StartPage_GET(c *gin.Context) {
	c.HTML(http.StatusOK, "startpage.tmpl", gin.H{
		"title": "نام درس و روز های هفته دروس خود را وارد کنید و سپس ثبت کنید و در نهایت دکمه نمایش برنامه را وارد کنید",
	})
	return
}

func (h *StartPage) StartPage_POST(c *gin.Context) {

	dname := c.PostForm("dname")
	days := c.PostFormArray("day")
	response := data{Name: dname, Days: days}
	datas = append(datas, response)
	c.JSON(http.StatusOK, datas)
}
