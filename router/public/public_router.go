package public

import (
	"github.com/Hidayathamir/golang_rest_api/handler"
	"github.com/gin-gonic/gin"
)

func AddPublicRouter(r *gin.Engine, allHandler handler.AllHandler) {
	addProductRoute(r, allHandler)
}
