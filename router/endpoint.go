// Routes/Routes.go
package Routes

import (
	us "login-user/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func EndPoint(usecase *us.Usecase) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "https://localhost:8080", "http://localhost:8080", "https://localhost"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Headers, Accept-Encoding, X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	grp1 := r.Group("/api")
	{
		grp1.POST("user", usecase.CreateUser)
		grp1.POST("user/login", usecase.Login)
		grp1.POST("user/update", usecase.UpdateLocation)
	}
	return r
}
