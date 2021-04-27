package response

import (
	"net/http"
	//"search-engine/models"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error) {
	//mega1Error, ok := err.(*models.Mega1Error)
	//if ok == true {
	//	return c.JSON(http.StatusOK, echo.Map{"code": mega1Error.Code, "message": mega1Error.Message})
	//}
	c.JSON(http.StatusOK, gin.H{"code": 4000, "message": err.Error()})
}

func Success(c *gin.Context, msg map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 2000, "data": msg})
}
