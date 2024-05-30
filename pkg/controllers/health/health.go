package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {

	ctx.IndentedJSON(http.StatusOK, `ok`)
}
