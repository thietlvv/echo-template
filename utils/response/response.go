package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormatResultAPI(status int, errStr string, data interface{}) interface{} {
	var dataR gin.H
	if status != 200 {
		dataR = gin.H{
			"message": errStr,
			"error":   status,
			"data":    data,
		}
		return dataR
	}
	dataR = gin.H{"data": data}
	return dataR["data"]
}

func Error(c *gin.Context, status int, errStr string, data interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, FormatResultAPI(http.StatusInternalServerError, err.Error(), data))
	}
	c.JSON(status, FormatResultAPI(status, errStr, data))
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, FormatResultAPI(http.StatusOK, "", data))
}
