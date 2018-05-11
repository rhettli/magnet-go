package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"magnet/rabbitmq/topic/send"
	"fmt"
)


func HomePage(c *gin.Context)  {

	c.HTML(http.StatusOK,"index.htm",nil)
}

func About(c *gin.Context)  {
	c.HTML(http.StatusOK,"about.htm",nil)
}


func Donate(c *gin.Context)  {
	c.HTML(http.StatusOK,"donate.htm",nil)
}


func SearchPage(c *gin.Context)  {
	key:=c.Params.ByName("key")
	page:=c.Params.ByName("page")

	if _,err:=strconv.Atoi(page);err!= nil {//不是数字格式的，则默认为第一页
		page="1"
	}

	c.HTML(http.StatusOK,"search.htm",gin.H{"key":key,"page":page})
}


func DetailPage(c *gin.Context)  {
	id:=c.Params.ByName("id")

	if _,err:=strconv.Atoi(id);err!= nil {//检测数字格式
		c.JSON(400,nil)
		return
	}

	jsonstr,err:=GetDetailApi(id)

	//fmt.Println(id,jsonstr)

	if err!=nil{
		c.JSON(400,nil)
		return
	}

	if err=send.SendRbmqMsg(id);err!=nil{
		fmt.Println("send msg to rabbitmq error:",err)
	}

	c.HTML(http.StatusOK,"detail.htm",gin.H{
		"id":id,
		"jsonstr":jsonstr,
		"name":jsonstr.Name,

	})
}


