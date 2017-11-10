package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func WebResultFail(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{"code":1001,"msg":"system error,please try later"})
}

func WebResultSuccess(data interface{},c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"data": data,"code":200,"msg":"success"})
}
