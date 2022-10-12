package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_blog/models"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct {

}

func (con UserController) Add(c *gin.Context){
	c.HTML(http.StatusOK,"admin/useradd.html",gin.H{})
}

func (con UserController) DoUpload(c *gin.Context){
	username :=c.PostForm("username")
	file,err :=c.FormFile("face")
	// file.Filename 獲取文件名稱
	dst := path.Join("./asset/upload",file.Filename)
	if err ==nil{
		c.SaveUploadedFile(file,dst)
	}
	//c.String(200,"執行上傳")
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"username":username,
		"dst":dst,
	})
}

func (con UserController) Edit(c *gin.Context){
	c.HTML(http.StatusOK,"admin/useredit.html",gin.H{})
}

func (con UserController) DoEdit(c *gin.Context){
	username :=c.PostForm("username")
	face1,err1 :=c.FormFile("face1")
	// file.Filename 獲取文件名稱
	dst1 := path.Join("./asset/upload",face1.Filename)
	if err1 ==nil{
		c.SaveUploadedFile(face1,dst1)
	}
	face2,err2 :=c.FormFile("face2")
	// file.Filename 獲取文件名稱
	dst2 := path.Join("./asset/upload",face2.Filename)
	if err2 ==nil{
		c.SaveUploadedFile(face2,dst2)
	}
	//c.String(200,"執行上傳")
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"username":username,
		"dst1":dst1,
		"dst2":dst2,
	})
}

func (con UserController) DoDo2(c *gin.Context){
	username :=c.PostForm("username")
	// 獲取上傳的文件
	file,err :=c.FormFile("face")
	// file.Filename 獲取文件名稱
	dst := path.Join("./asset/upload",file.Filename)
	if err ==nil{
		// 2.獲取後綴名，判斷類型是否正確，(jpg.png.gif.jpeg)
		extName := path.Ext(file.Filename)
		allowExtMap :=map[string]bool{
			".jpg":true,
			".png":true,
			".gif":true,
			".jpeg":true,
		}
		if _,ok := allowExtMap[extName] ; !ok{
			c.String(200,"上傳數據不合法")
			return
		}
		// 3.創建圖片保存目錄 asset/upload/20220601

		day := models.GetDay()
		dir := "./asset/upload/"+day
		err := os.MkdirAll(dir,0666)
		if (err!=nil){
			fmt.Println(err)
			c.String(200,"MKdirAll失敗")
			return
		}
		// 4.生成文件名稱和文件保存的目錄
		fileName :=strconv.FormatInt(models.GetUnix(),10) + extName
		// 5.執行上傳
		dst := path.Join(dir,fileName)
		c.SaveUploadedFile(file,dst)
	}
	//c.String(200,"執行上傳")
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"username":username,
		"dst":dst,
	})
}

func (con UserController) Index(c *gin.Context){
	greet01()
	byteLearn()
	c.String(200,"用戶列表")
}

func greet01() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("GO GO GO") // 有可能因為主程式已經執行完而看不到
		// 使用 Buffered Channel 的話，不會等到 channel 中的值讀完才結束主程式
		fmt.Printf("Receive value from channel %v\n", <-c)
	}()
	fmt.Println("Before Receive")
	// STEP 1：寫入 channel
	c <- true
	c <- false // block here
	fmt.Println("After Receive")
}

func greet02(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func greet03(){
	var firstlist []int= []int{1,2,3}
	fmt.Println(firstlist)
	for _,num :=range(firstlist){
		fmt.Println(num)
	}
}

func byteLearn(){
	name :="andy"
	for _,n := range name{
		fmt.Println(n)
	}
}

