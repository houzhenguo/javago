package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 匿名函数测试2
	r.GET("/hou", func(c *gin.Context) {
		// 获取参数
		username := c.Query("username")
		sex := c.DefaultQuery("sex", "男")
		c.JSON(200, gin.H{
			"name": username,
			"age":  18,
			"sex":  sex,
		})
	})
	// 使用结构体
	r.GET("json", func(c *gin.Context) {
		var msg struct {
			Name string
			Age  int `json:"age"` // 必须要大写
		}
		msg.Name = "张三"
		msg.Age = 23
		c.JSON(http.StatusOK, msg)
	})

	// 获取参数 POST form
	r.POST("/post", func(c *gin.Context) {
		username := c.PostForm("username") // 注意参数格式为 form表单格式
		fmt.Println("username is", username)
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
		})
	})
	// 获取json 参数
	r.POST("/post/json", func(c *gin.Context) {
		a, b := c.GetRawData()       // 在request 中读取 c.Request.Body 请求数据
		fmt.Println("b is", b)       // b is nil
		var m map[string]interface{} // 定义map或者结构体
		_ = json.Unmarshal(a, &m)    // 反序列化
		c.JSON(http.StatusOK, m)
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
