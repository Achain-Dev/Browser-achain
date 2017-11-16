package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func WebResultFail(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"code":1001,"msg":"system error,please try later"})
}

func WebResultMiss(c *gin.Context,code int,msg string)  {
	c.JSON(http.StatusOK,gin.H{"code":code,"msg":msg})
}

func WebResultSuccess(data interface{},c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"data": data,"code":200,"msg":"success"})
}

func WebResultSuccessWithMap(c *gin.Context,data map[string]interface{}) {
	c.JSON(http.StatusOK, data)
}

