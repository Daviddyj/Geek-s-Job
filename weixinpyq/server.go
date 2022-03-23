package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn.go/pkg/apis"
	"strconv"

	"log"
	"net/http"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:1165207594dyj@tcp(127.0.0.1:3306)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var rankServer ServerInterface = NewDbRank(conn, NewFatRateRank())
	if initFank, ok := rankServer.(RankInitInterface); ok {
		if err := initFank.Init(); err != nil {
			log.Fatal("初始化失败:", err)
		}
	}
	r := gin.Default()

	pprof.Register(r)

	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "无法解析注册信息"})
			return
		}
		//Todo 注册到排行榜

		if err := rankServer.RegisterPersonInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "注册失败"})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.PUT("/personalinfo", func(c *gin.Context) {
		pi := &apis.PersonInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "无法解析注册信息"})
			return
		}
		if resp, err := rankServer.UpdatePersonInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "服务器更新用户失败"})
			return
		} else {
			c.JSON(200, resp)
		}

	})
	r.GET("/rank/:name", func(c *gin.Context) {
		name := c.Param("name")
		if rank, err := rankServer.GetFateRate(name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "获取用户排名失败"})
			return
		} else {
			c.JSON(200, rank)
		}
	})
	r.GET("/ranktop", func(c *gin.Context) {
		if frTop, err := rankServer.GetTop(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "获取用户前十排名失败"})
			return
		} else {
			c.JSON(200, frTop)
		}
	})
	//删除pyq状态,标记要删除的用户数据Visiable为1   表示数据不可见
	r.DELETE("/deleteStatus/:id", func(c *gin.Context) {
		id := c.Param("id")
		newid, err := strconv.Atoi(id)
		if err != nil {
			log.Println("删除状态时转换id出错:\n", err)
		}
		fmt.Println("我在这里:", newid)
		if err := rankServer.DeletePersonPyq(newid); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "删除用户状态失败"})
			return
		} else {
			c.JSON(200, gin.H{
				"message": "delete success",
			})
		}
	})
	//查询所有用户状体,选择Visiable为0的数据，表示可见的朋友圈状态用户
	r.GET("/showPersonPyq", func(c *gin.Context) {
		if result, err := rankServer.ShowPersonPyq(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "获取微信朋友圈失败"})
			return
		} else {
			c.JSON(200, result)
		}
	})

	if err := r.Run("127.0.0.1:8081"); err != nil {
		log.Fatal(err)
	}

}
