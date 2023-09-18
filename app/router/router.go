package router

import (
	"be-map-test/app/controller/track"
	"be-map-test/app/db"

	docs "be-map-test/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) // for release mod, uncomment if need it
	r := gin.Default()
	db.ConnectDatabase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	docs.SwaggerInfo.Title = "BE Maps API "
	docs.SwaggerInfo.Description = "This is API for FE Test."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	apiUri := r.Group("/v1")

	trackRoute := apiUri.Group("")
	{
		trackRoute.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// trackRoute.GET("/track", track.GetAll)
		trackRoute.GET("/track/:name", track.GetAllByName)
		trackRoute.POST("/track", track.Create)
		trackRoute.PUT("/track", track.Update)
		trackRoute.DELETE("/track", track.Delete)
	}

	return r
}
