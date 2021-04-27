package monitor

import "github.com/gin-gonic/gin"

func Routes(g *gin.RouterGroup) {
	gMonitor := g.Group("/m")

	gMonitor.GET("/health_check", HealthCheck)
	gMonitor.GET("/load_test", LoadTest)
}
