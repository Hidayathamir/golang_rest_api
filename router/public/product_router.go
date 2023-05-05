package public

import (
	"github.com/Hidayathamir/golang_rest_api/handler"
	"github.com/gin-gonic/gin"
)

func addProductRoute(r *gin.Engine, allHandler handler.AllHandler) {
	r.POST("/products", allHandler.ProductHandler.AddProduct)
	r.GET("/products", allHandler.ProductHandler.GetProducts)
}
