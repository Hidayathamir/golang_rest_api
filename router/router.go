package router

import (
	"github.com/Hidayathamir/golang_rest_api/handler"
	"github.com/Hidayathamir/golang_rest_api/router/public"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func GetRouter(allHandler handler.AllHandler) *gin.Engine {
	if r == nil {
		r = gin.Default()
	}
	public.AddPublicRouter(r, allHandler)
	return r
}
