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
	log.Println(body.Bridge)
	ovsagent.CreateDistrubutedSwitch(body.Bridge)
	ctx.IndentedJSON(http.StatusOK, body)
}
