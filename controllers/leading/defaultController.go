package leading

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hunterhug/go_image"
	qrcode "github.com/skip2/go-qrcode"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(http.StatusOK, "首页")
}

func (con DefaultController) Thumbnail1(c *gin.Context) {
	//按宽度进行比例缩放，输入输出都是文件
	//filename string, savepath string, width int
	fileName := "static/upload/01.jpg"
	savepath := "static/upload/01_200.jpg"
	err := go_image.ScaleF2F(fileName, savepath, 200)
	if err != nil {
		c.String(http.StatusOK, "生成图片失败")
		fmt.Println(err)
		return
	}
	c.String(http.StatusOK, "生成图片成功")
}

func (con DefaultController) Thumbnail2(c *gin.Context) {
	fileName := "static/upload/01.jpg"
	savepath := "static/upload/01_400.png"
	//按宽度和高度进行比例缩放，输入和输出都是文件
	err := go_image.ThumbnailF2F(fileName, savepath, 400, 400)
	if err != nil {
		c.String(http.StatusOK, "失败")
		return
	}
	c.String(http.StatusOK, "Thumbnail2成功")
}

// 生成二维码
func (con DefaultController) Qrcode1(c *gin.Context) {
	var png []byte
	png, err := qrcode.Encode("http://www.baidu.com", qrcode.Medium, 300)
	if err != nil {
		c.String(http.StatusOK, "生成二维码失败")
		return
	}
	c.String(http.StatusOK, string(png))
}
func (con DefaultController) Qrcode2(c *gin.Context) {
	// 二维码图片 保存到本地
	savepath := "static/upload/二维码.png"
	err := qrcode.WriteFile("http://www.baidu.com", qrcode.Medium, 400, savepath)
	if err != nil {
		c.String(http.StatusOK, "生成二维码图片失败")
		return
	}
	file, _ := ioutil.ReadFile(savepath)
	c.String(http.StatusOK, string(file))
}

// func (con DefaultController) Index(c *gin.Context) {
// 设置sessions
// session := sessions.Default(c)
// 配置session的过期时间
// session.Options(sessions.Options{
// MaxAge: 3600 * 6, // 6hrs   MaxAge单位是秒
// })
// session.Set("username", "张三 111")
// session.Save() //设置session的时候必须调用
//
// c.HTML(http.StatusOK, "default/index.html", gin.H{
// "msg": "我是一个msg",
// "t":   1629788418,
// })
// }
// func (con DefaultController) News(c *gin.Context) {
// 获取sessions
// session := sessions.Default(c)
// username := session.Get("username")
// c.String(200, "username=%v", username)
// }
//
