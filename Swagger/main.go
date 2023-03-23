package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/ini8labs/Go-Learning/Swagger/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func SampleResponse(g *gin.Context) {
	g.JSON(http.StatusOK, "this is sample response")
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
			eg.GET("sample", SampleResponse)

		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
