package login

import (
	"github.com/fighthorse/redisAdmin/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterHttp(r *gin.Engine) {
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/login")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/out", loginOutEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/check", checkEndpoint)
	}
}
