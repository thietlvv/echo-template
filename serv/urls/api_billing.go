package urls

import (
	"billing/serv/monitor"
	"billing/serv/users"

	"github.com/gin-gonic/gin"
)

func InitUrlsBilling(r *gin.Engine) {
	// Init group /billing
	gBilling := r.Group("/billing")

	// Init group /billing/m
	monitor.Routes(gBilling)

	// Init group /billing/users
	users.Routes(gBilling)
}
