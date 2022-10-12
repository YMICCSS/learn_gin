package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang_blog/models"
	"golang_blog/router"
	"html/template"
	"net/http"
)

type IndexData struct {
	Title   string
	Content string
}

func defaults(c *gin.Context) {
	mapdata := map[string]string{}
	mapdata["Title"]="預設主題"
	mapdata["Content"]="預設內容"
	data := new(IndexData)
	data.Title = "預設主題"
	data.Content = "預設內容"
	c.HTML(http.StatusOK, "admin/index.html",gin.H{
		"mapdata" : mapdata,
		"data" : data,
	})
}

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

type UserInfo struct{
	Username string `json:"username" form:"username"`
	Password int `json:"password" form:"password"`
}


func main() {
	// 創建默認的路由引擎
	server := gin.Default()

	// 自定義模板函數，注意要把這個函數放在加載模板前
	server.SetFuncMap(template.FuncMap{
		"UnixToTime":models.UnixToTime,
	})

	// LoadHTMLGlob 方法不有二級目錄，有的話會panic
	server.LoadHTMLGlob("template/**/*")

	// 配置靜態服務
	server.Static("/dwz","./asset")

	// 配置 session 中間件
	// 創建基於 cookie 的存儲引擎，secrect11111 參數是用於加密的密鑰
	store := cookie.NewStore([]byte("secret11111"))
	// 配置 session 的中間件 , store 是前面的存儲引擎，我們可以換成其他存儲引擎
	server.Use(sessions.Sessions("mysession", store))

	// 前台路由
	router.MainRouterInit(server)

	// 後台路由
	router.AdminRouterInit(server)

	// 啟動一個web服務
	server.Run(":8888")
}

//func Response(c *gin.Context){
//
//	//one := c.PostForm("result")
//	//two := c.DefaultQuery("result", "22")
//	//ome, ok := c.GetPostForm("result")
//	//fmt.Println(one)
//	//fmt.Println(two)
//	//fmt.Println(ome)
//	//fmt.Println(ok)
//	postData := c.PostForm("result")
//	fmt.Println(len(postData))
//	fmt.Println(postData)
//	for _,num := range postData{
//		fmt.Println(num)
//	}
//	c.JSON(200,gin.H{"aa":"aa"})
//
//}



