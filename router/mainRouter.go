package router

import (
	"github.com/gin-gonic/gin"
	Admin "golang_blog/controllers/admin"
	"golang_blog/middlewares"
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

func MainRouterInit(r *gin.Engine){
	defaultRouter :=r.Group("/main",middlewares.InitMiddleware)
	{
		defaultRouter.GET("/user", defaults)
		defaultRouter.GET("/user/add", Admin.UserController{}.Add)
		defaultRouter.GET("/user/edit", Admin.UserController{}.Edit)
		defaultRouter.GET("/user/articleControllerIndex", Admin.ArticleController{}.Index)
		defaultRouter.GET("/user/articleControllerNews", Admin.ArticleController{}.News)

		defaultRouter.GET("/user/deleteCookie", Admin.ArticleController{}.DeleteCookie)
		defaultRouter.GET("/user/setCookie", Admin.ArticleController{}.Setcookie)
		defaultRouter.GET("/user/getCookie", Admin.ArticleController{}.Getcookie)

		defaultRouter.GET("/user/setsession", Admin.ArticleController{}.SetSession)
		defaultRouter.GET("/user/getsession", Admin.ArticleController{}.GetSession)
		defaultRouter.GET("/user/deletesession", Admin.ArticleController{}.DeleteSession)

		// 以下為 post
		defaultRouter.POST("/user/DoDo2", Admin.UserController{}.DoDo2)
		defaultRouter.POST("/user/doUpload", Admin.UserController{}.DoUpload)

	}
}




