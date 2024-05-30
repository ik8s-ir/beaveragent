package v1alpha1Controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ik8s-ir/beaveragent/pkg/ovsagent"
	"github.com/ik8s-ir/beaveragent/pkg/types"
)

func PostOvsBridge(ctx *gin.Context) {
	var body types.VswitchPostBody
	if err := ctx.BindJSON(&body); err != nil {
		return
	}
	r, err := ovsagent.CreateDistrubutedSwitch(body.Bridge,body.Topology)
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

func DeleteOvsBridge(ctx *gin.Context) {
	bridge:=ctx.Param("bridge")
	r, err := ovsagent.DeleteDistrubutedSwitch(bridge)
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
