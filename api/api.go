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


// <html lang="fa">
// <head>
// <meta charset="UTF-8">
// <title>فرم ارسال اطلاعات</title>
// </head>
// <body>
//     <h2 dir="rtl">
//     {{.title}}
//     </h2>
// <form dir="rtl" action="http://localhost:5000/startpage" method="POST">
//      &nbsp;&nbsp;&nbsp;&nbsp;<label for="dname" style="font-size: 30px;">نام درس:</label><br>
//      &nbsp;&nbsp;<input style="font-size: 17px;" type="text" id="dname" name="dname"><br>
//      &nbsp;&nbsp;<input style="width:17px; height:17px;" type="checkbox" id="d1" name="day" value="شنبه">
//      &nbsp;&nbsp;<label for="d1" style="font-size: 30px;">شنبه</label><br>
//      &nbsp;&nbsp;<input style="width:17px; height:17px;" type="checkbox" id="d2" name="day" value="یک شنبه">
//      &nbsp;&nbsp;<label for="d2" style="font-size: 30px;">یک شنبه</label><br>
//      &nbsp;&nbsp;<input style="width:17px; height:17px;" type="checkbox" id="d3" name="day" value="دوشنبه">
//      &nbsp;&nbsp;<label for="html" style="font-size: 30px;">دوشنبه</label><br>
//      &nbsp;&nbsp;<input style="width:17px; height:17px;" type="checkbox" id="d4" name="day" value="سه شنبه">
//      &nbsp;&nbsp;<label for="html" style="font-size: 30px;">سه شنبه </label><br>
//      &nbsp;&nbsp;<input style="width:17px; height:17px;" type="checkbox" id="d5" name="day" value="چهارشنه">
//      &nbsp;&nbsp;<label for="html" style="font-size: 30px;">چهارشنبه</label><br><br>
//      &nbsp;&nbsp;<input type="submit"style="background-color:black; color:white; font-size: 30px;" value="ثبت درس" />
// </form>
// </body>
// </html>