package public

import (
	"github.com/gin-gonic/gin"
)

func AddPublicRouter(r *gin.Engine) {
	addProductRoute(r)
}
