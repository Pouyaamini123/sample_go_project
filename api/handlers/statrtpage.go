package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type data_time struct {
	Time1  string
	Time2  string
	Time3  string
	Time4  string
	Time5  string
	Time6  string
	Time7  string
	Time8  string
	Time9  string
	Time10 string
	Time11 string
	Time12 string
	Time13 string
	Time14 string
	Time15 string
	Time16 string
	Time17 string
	Time18 string
	Time19 string
	Time20 string
	Time21 string
	Time22 string
	Time23 string
	Time24 string
}

type data_day struct {
	D1 data_time
	D2 data_time
	D3 data_time
	D4 data_time
	D5 data_time
}

var datem = data_day{}

type data struct {
	Name [6][25]string
}

var datas = data{}
var datas_chek = data{}

type StartPage struct {
}

func NewStartPage() *StartPage {
	return &StartPage{}
}

func (h *StartPage) StartPage_GET(c *gin.Context) {
	c.HTML(http.StatusOK, "startpage.tmpl", gin.H{
		"title": "نام و روز های هفته و تایم دروس خود را وارد کنید و سپس ثبت کنید",
		"datem": datem,
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
		i1, _ := strconv.Atoi(dtimee)
		i2, _ := strconv.Atoi(dtimes)
		if i2-i1>=0{
			c.JSON(432, datas)
					return
		}
		for _, day1 := range days {
			day, _ := strconv.Atoi(day1)
			for i := i2; i < i1; i++ {
				if datas_chek.Name[day][i] == "" {
					datas_chek.Name[day][i] = dname
				} else {
					result = 0
					c.JSON(422, datas)
					return
				}
			}
		}

		if result == 1 {
			for _, day1 := range days {
				day, _ := strconv.Atoi(day1)
				for i := i2; i < i1; i++ {
					datas.Name[day][i] = dname
					switch {
					case day == 1 && i == 1:
						datem.D1.Time1 = dname
					case day == 1 && i == 2:
						datem.D1.Time2 = dname
					case day == 1 && i == 3:
						datem.D1.Time3 = dname
					case day == 1 && i == 4:
						datem.D1.Time4 = dname
					case day == 1 && i == 5:
						datem.D1.Time5 = dname
					case day == 1 && i == 6:
						datem.D1.Time6 = dname
					case day == 1 && i == 7:
						datem.D1.Time7 = dname
					case day == 1 && i == 8:
						datem.D1.Time8 = dname
					case day == 1 && i == 9:
						datem.D1.Time9 = dname
					case day == 1 && i == 10:
						datem.D1.Time10 = dname
					case day == 1 && i == 11:
						datem.D1.Time11 = dname
					case day == 1 && i == 12:
						datem.D1.Time12 = dname
					case day == 1 && i == 13:
						datem.D1.Time13 = dname
					case day == 1 && i == 14:
						datem.D1.Time14 = dname
					case day == 1 && i == 15:
						datem.D1.Time15 = dname
					case day == 1 && i == 16:
						datem.D1.Time16 = dname
					case day == 1 && i == 17:
						datem.D1.Time17 = dname
					case day == 1 && i == 18:
						datem.D1.Time18 = dname
					case day == 1 && i == 19:
						datem.D1.Time19 = dname
					case day == 1 && i == 20:
						datem.D1.Time20 = dname
					case day == 1 && i == 21:
						datem.D1.Time21 = dname
					case day == 1 && i == 22:
						datem.D1.Time22 = dname
					case day == 1 && i == 23:
						datem.D1.Time23 = dname
					case day == 1 && i == 24:
						datem.D1.Time24 = dname

					case day == 2 && i == 1:
						datem.D2.Time1 = dname
					case day == 2 && i == 2:
						datem.D2.Time2 = dname
					case day == 2 && i == 3:
						datem.D2.Time3 = dname
					case day == 2 && i == 4:
						datem.D2.Time4 = dname
					case day == 2 && i == 5:
						datem.D2.Time5 = dname
					case day == 2 && i == 6:
						datem.D2.Time6 = dname
					case day == 2 && i == 7:
						datem.D2.Time7 = dname
					case day == 2 && i == 8:
						datem.D2.Time8 = dname
					case day == 2 && i == 9:
						datem.D2.Time9 = dname
					case day == 2 && i == 10:
						datem.D2.Time10 = dname
					case day == 2 && i == 11:
						datem.D2.Time11 = dname
					case day == 2 && i == 12:
						datem.D2.Time12 = dname
					case day == 2 && i == 13:
						datem.D2.Time13 = dname
					case day == 2 && i == 14:
						datem.D2.Time14 = dname
					case day == 2 && i == 15:
						datem.D2.Time15 = dname
					case day == 2 && i == 16:
						datem.D2.Time16 = dname
					case day == 2 && i == 17:
						datem.D2.Time17 = dname
					case day == 2 && i == 18:
						datem.D2.Time18 = dname
					case day == 2 && i == 19:
						datem.D2.Time19 = dname
					case day == 2 && i == 20:
						datem.D2.Time20 = dname
					case day == 2 && i == 21:
						datem.D2.Time21 = dname
					case day == 2 && i == 22:
						datem.D2.Time22 = dname
					case day == 2 && i == 23:
						datem.D2.Time23 = dname
					case day == 2 && i == 24:
						datem.D2.Time24 = dname

					case day == 3 && i == 1:
						datem.D3.Time1 = dname
					case day == 3 && i == 2:
						datem.D3.Time2 = dname
					case day == 3 && i == 3:
						datem.D3.Time3 = dname
					case day == 3 && i == 4:
						datem.D3.Time4 = dname
					case day == 3 && i == 5:
						datem.D3.Time5 = dname
					case day == 3 && i == 6:
						datem.D3.Time6 = dname
					case day == 3 && i == 7:
						datem.D3.Time7 = dname
					case day == 3 && i == 8:
						datem.D3.Time8 = dname
					case day == 3 && i == 9:
						datem.D3.Time9 = dname
					case day == 3 && i == 10:
						datem.D3.Time10 = dname
					case day == 3 && i == 11:
						datem.D3.Time11 = dname
					case day == 3 && i == 12:
						datem.D3.Time12 = dname
					case day == 3 && i == 13:
						datem.D3.Time13 = dname
					case day == 3 && i == 14:
						datem.D3.Time14 = dname
					case day == 3 && i == 15:
						datem.D3.Time15 = dname
					case day == 3 && i == 16:
						datem.D3.Time16 = dname
					case day == 3 && i == 17:
						datem.D3.Time17 = dname
					case day == 3 && i == 18:
						datem.D3.Time18 = dname
					case day == 3 && i == 19:
						datem.D3.Time19 = dname
					case day == 3 && i == 20:
						datem.D3.Time20 = dname
					case day == 3 && i == 21:
						datem.D3.Time21 = dname
					case day == 3 && i == 22:
						datem.D3.Time22 = dname
					case day == 3 && i == 23:
						datem.D3.Time23 = dname
					case day == 3 && i == 24:
						datem.D3.Time24 = dname

					case day == 4 && i == 1:
						datem.D4.Time1 = dname
					case day == 4 && i == 2:
						datem.D4.Time2 = dname
					case day == 4 && i == 3:
						datem.D4.Time3 = dname
					case day == 4 && i == 4:
						datem.D4.Time4 = dname
					case day == 4 && i == 5:
						datem.D4.Time5 = dname
					case day == 4 && i == 6:
						datem.D4.Time6 = dname
					case day == 4 && i == 7:
						datem.D4.Time7 = dname
					case day == 4 && i == 8:
						datem.D4.Time8 = dname
					case day == 4 && i == 9:
						datem.D4.Time9 = dname
					case day == 4 && i == 10:
						datem.D4.Time10 = dname
					case day == 4 && i == 11:
						datem.D4.Time11 = dname
					case day == 4 && i == 12:
						datem.D4.Time12 = dname
					case day == 4 && i == 13:
						datem.D4.Time13 = dname
					case day == 4 && i == 14:
						datem.D4.Time14 = dname
					case day == 4 && i == 15:
						datem.D4.Time15 = dname
					case day == 4 && i == 16:
						datem.D4.Time16 = dname
					case day == 4 && i == 17:
						datem.D4.Time17 = dname
					case day == 4 && i == 18:
						datem.D4.Time18 = dname
					case day == 4 && i == 19:
						datem.D4.Time19 = dname
					case day == 4 && i == 20:
						datem.D4.Time20 = dname
					case day == 4 && i == 21:
						datem.D4.Time21 = dname
					case day == 4 && i == 22:
						datem.D4.Time22 = dname
					case day == 4 && i == 23:
						datem.D4.Time23 = dname
					case day == 4 && i == 24:
						datem.D4.Time24 = dname

					case day == 5 && i == 1:
						datem.D5.Time1 = dname
					case day == 5 && i == 2:
						datem.D5.Time2 = dname
					case day == 5 && i == 3:
						datem.D5.Time3 = dname
					case day == 5 && i == 4:
						datem.D5.Time4 = dname
					case day == 5 && i == 5:
						datem.D5.Time5 = dname
					case day == 5 && i == 6:
						datem.D5.Time6 = dname
					case day == 5 && i == 7:
						datem.D5.Time7 = dname
					case day == 5 && i == 8:
						datem.D5.Time8 = dname
					case day == 5 && i == 9:
						datem.D5.Time9 = dname
					case day == 5 && i == 10:
						datem.D5.Time10 = dname
					case day == 5 && i == 11:
						datem.D5.Time11 = dname
					case day == 5 && i == 12:
						datem.D5.Time12 = dname
					case day == 5 && i == 13:
						datem.D5.Time13 = dname
					case day == 5 && i == 14:
						datem.D5.Time14 = dname
					case day == 5 && i == 15:
						datem.D5.Time15 = dname
					case day == 5 && i == 16:
						datem.D5.Time16 = dname
					case day == 5 && i == 17:
						datem.D5.Time17 = dname
					case day == 5 && i == 18:
						datem.D5.Time18 = dname
					case day == 5 && i == 19:
						datem.D5.Time19 = dname
					case day == 5 && i == 20:
						datem.D5.Time20 = dname
					case day == 5 && i == 21:
						datem.D5.Time21 = dname
					case day == 5 && i == 22:
						datem.D5.Time22 = dname
					case day == 5 && i == 23:
						datem.D5.Time23 = dname
					case day == 5 && i == 24:
						datem.D5.Time24 = dname
					}
				}
			}
		}
		//response := data{Name: dname}
		//datas = append(datas, response)
		c.JSON(http.StatusOK, datas)
		return
	}
}
