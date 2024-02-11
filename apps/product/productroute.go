package product

import "github.com/gin-gonic/gin"

func RegisterRoute(r *gin.RouterGroup, url string) {
	r.GET(url, Index)
	r.GET(url+"/:id", Show)
	r.POST(url, Create)
	r.PUT(url+"/:id", Update)
	r.DELETE(url+"/:id", Delete)
}
