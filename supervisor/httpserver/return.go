package httpserver

import "github.com/gin-gonic/gin"

func ReturnError(httpcode int, err error, c *gin.Context) {
	c.JSON(httpcode, err.Error())
}

func ReturnSuccess(httpcode int, content interface{}, c *gin.Context) {
	c.JSON(httpcode, content)
}
