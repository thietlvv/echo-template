package monitor

import (
	"fmt"
	"net/http"

	main_db "billing/db/main_db"
	"billing/utils/response"

	"github.com/gin-gonic/gin"
)

// HealthCheck ... health check service
func HealthCheck(c *gin.Context) {
	resultDB := main_db.CheckDatabase()
	if resultDB == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 4000, "message": "Service Unavailable"})
	}
	response.Success(c, gin.H{"DB_TIME": resultDB})
}

// LoadTest ...
func LoadTest(c *gin.Context) {

	// #1
	for i := 0; i < 10; i++ {
		tmp := i
		fmt.Println(tmp)
	}

	// #2
	resultDB := main_db.CheckDatabase()
	fmt.Println(resultDB)

	// #3
	tmp := 0
	for i := 0; i < 10000; i++ {
		tmp += i
	}
	fmt.Println(tmp)

	// #4
	resultDB = main_db.CheckDatabase()
	fmt.Println(resultDB)

	c.JSON(http.StatusOK, "ok")
}
