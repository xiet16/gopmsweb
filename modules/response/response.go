package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.xiet16.com/gopmsweb/modules/lang"
)

func ShowError(c *gin.Context, msg string) {
	msg = lang.Get(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}

func ShowData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": data,
	})
}
