package v1alpha1Controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ik8s-ir/beaveragent/pkg/ovsagent"
	"github.com/ik8s-ir/beaveragent/pkg/types"
)

func PostOvsBridge(ctx *gin.Context) {
	var body types.TovsBridge
	if err := ctx.BindJSON(&body); err != nil {
		return
	}
	r, err := ovsagent.CreateDistrubutedSwitch(body.Bridge)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": r,
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": r,
		})
	}
}
