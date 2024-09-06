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
	r.POST("", handler.StartPage_POST)
}
// <!DOCTYPE html>
// <html lang="fa">
// <head>
// <meta charset="UTF-8">
// <title>فرم ارسال اطلاعات</title>
// </head>
// <body>
// <form dir="rtl" action="http://localhost:5000/startpage" method="POST">
// <label for="dname" style="font-size: 30px;">نام درس:</label><br>
// <input style="font-size: 17px;" type="text" id="dname" name="dname"><br>
// <input style="width:17px; height:17px;" type="checkbox" id="d1" name="day" value="شنبه">
// <label for="d1" style="font-size: 30px;">شنبه</label><br>
// <input style="width:17px; height:17px;" type="checkbox" id="d2" name="day" value="یک شنبه">
// <label for="d2" style="font-size: 30px;">یک شنبه</label><br>
// <input style="width:17px; height:17px;" type="checkbox" id="d3" name="day" value="دوشنبه">
// <label for="d3" style="font-size: 30px;">دوشنبه</label><br>
// <input style="width:17px; height:17px;" type="checkbox" id="d4" name="day" value="سه شنبه">
// <label for="d4" style="font-size: 30px;">سه شنبه</label><br>
// <input style="width:17px; height:17px;" type="checkbox" id="d5" name="day" value="چهارشنبه">
// <label for="d5" style="font-size: 30px;">چهارشنبه</label><br><br>
// <input type="submit" onclick="location.href='startpage';" style="background-color:black; color:white; font-size: 30px;" value="ثبت درس">
// </form>
// </body>
// </html>
