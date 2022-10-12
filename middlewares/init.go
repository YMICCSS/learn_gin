package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)


func TT(s string,c chan string){
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(s)
	c<-"OOO"
	c<-"111"
	c<-"111"
}


//判斷用戶是否登入
func InitMiddleware(c *gin.Context){

	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username","張三")
	//定義一個goriutine日誌
	// 用copy方式才可以在gorutine中使用
	cCp :=c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done ! in path " +cCp.Request.URL.Path)
	}()

}



