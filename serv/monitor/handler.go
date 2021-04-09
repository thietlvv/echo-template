package monitor

import (
	"fmt"
	"net/http"

	main_db "billing/db/main_db"
	"billing/utils/response"

	"github.com/labstack/echo/v4"
)

// HealthCheck ... health check service
func HealthCheck(c echo.Context) error {
	resultDB := main_db.CheckDatabase()
	if resultDB == "" {
		return c.JSON(http.StatusServiceUnavailable, echo.Map{"code": 4000, "message": "Service Unavailable"})
	}
	return response.Success(c, echo.Map{"DB_TIME": resultDB})
}

// LoadTest ...
func LoadTest(c echo.Context) error {

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

	return response.Success(c, echo.Map{"data": "ok"})
}
