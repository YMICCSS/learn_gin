package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang_blog/models"
)

type ArticleController struct{

}

func (con ArticleController) Index(c *gin.Context){
	// 設置cookie
	// 第六個secure參數為true時，cookie在HTTP中是無效，在HTTPS中才有效
	// 第七個參數httponly,設置為true代表只有後端才可以操作它，設置為false，代表可以用JS、applet操作我們的cookie
	fmt.Println(models.UnixToTime(1629788418))
	c.SetCookie("username","張三",3600,"/","localhost",false,true)
	c.HTML(200,"default/index.html",gin.H{
		"msg" : "我是一個msg",
		"t" : 1629788418,
	})
}

func (con ArticleController) News(c *gin.Context){
	// 獲取 cookie
	username,_ := c.Cookie("username")
	c.String(200,"cookie = "+username)
}

func (con ArticleController) Setcookie(c *gin.Context){
	c.SetCookie("username","張三",3600,"/","127.0.0.1",false,true)
	_,err :=c.Cookie("username")
	if err != nil {
		fmt.Println(err)
	}
	c.String(200,"設置成功")
}

func (con ArticleController) Getcookie(c *gin.Context){
	value,err :=c.Cookie("username")
	if err != nil {
		fmt.Println(err)
	}
	c.String(200,"取得成功 :"+value)
}

func (con ArticleController) DeleteCookie(c *gin.Context){
	c.SetCookie("username","張三",-1,"/","127.0.0.1",false,true)
	c.String(200,"刪除成功")
}

func (con ArticleController) SetSession(c *gin.Context){
	// 設置 session
	session := sessions.Default(c)
	session.Set("username","張三 222")
	session.Save() // 設置session的時候必須調用
	c.String(200,"Set session成功")
}

func (con ArticleController) GetSession(c *gin.Context){
	// 獲取 session
	session := sessions.Default(c)
	username := session.Get("username")
	c.String(200,"username = %v",username)
}

func (con ArticleController) DeleteSession(c *gin.Context){
	// 獲取 session
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
	c.String(200,"已刪除session")
}



