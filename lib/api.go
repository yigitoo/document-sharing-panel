package lib

import "github.com/gin-gonic/gin"

var ApiConfig Config = NewConfig()

func SetupApi() *gin.Engine {
	r := gin.Default()

	return r
}
