package v1alpha1Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ik8s-ir/beaveragent/pkg/types"
)

func PostOvsBridge(ctx *gin.Context) {
	var body types.TovsBridge
	if err := ctx.BindJSON(&body); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, body)
}
