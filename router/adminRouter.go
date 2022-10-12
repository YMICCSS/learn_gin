package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func admin(c *gin.Context) {
	mapdata := map[string]string{}
	mapdata["Title"]="首頁map"
	mapdata["Content"]="首頁內容map"
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "首頁內容"
	c.HTML(http.StatusOK, "admin/index.html",gin.H{
		"mapdata" : mapdata,
		"data" : data,
		"hobby": []string{"吃飯","睡覺","寫代碼"},
		"newlist":[]interface{}{"吃飯","睡覺","寫代碼"},
		"date": 1645154003,
	})
}

func AdminRouterInit(c *gin.Engine){
	apiRouter :=c.Group("/api")
	{
		apiRouter.GET("/admin", admin)
		apiRouter.GET("/admin/name", admin)
	}
}





