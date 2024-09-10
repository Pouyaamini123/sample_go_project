package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type data struct {
	Name [6][25]string
}

var datas=data{}
var datas_chek=data{}

type StartPage struct {
}

func NewStartPage() *StartPage {
	return &StartPage{}
}

func (h *StartPage) StartPage_GET(c *gin.Context) {
	c.HTML(http.StatusOK, "startpage.tmpl", gin.H{
		"title": "نام و روز های هفته و تایم دروس خود را وارد کنید و سپس ثبت کنید",
		"datas": datas,
	})
	return
}

func (h *StartPage) StartPage_POST(c *gin.Context) {

	dname := c.PostForm("dname")
	days := c.PostFormArray("day")
	dtimes := c.PostForm("starttime")
	dtimee := c.PostForm("endtime")
	if dname == "" || len(days) == 0 {
		c.JSON(400, datas)
	} else {
		result := 1
		i1,_:=strconv.Atoi(dtimee)
		i2,_:=strconv.Atoi(dtimes)
		sub :=  i1-i2
		for day := range days {
			for i := 1; i <= sub; i++ {
				if datas_chek.Name[day][i] == "" {
					datas_chek.Name[day][i] =dname
				}else{
					result=0
					c.JSON(422, datas)
					return
				}
			}
		}
		if result==1{
			for day := range days {
				for i := 1; i <= sub; i++ {
						datas.Name[day][i] =dname
				}
			}	
		}
		//response := data{Name: dname}
		//datas = append(datas, response)
		c.JSON(http.StatusOK, datas)
		return
	}
}
