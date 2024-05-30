package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ik8s-ir/beaveragent/pkg/controllers/health"
	v1alpha1Controller "github.com/ik8s-ir/beaveragent/pkg/controllers/v1alpha1"
)

func main() {
	router := gin.Default()
	router.GET("/health", health.GetHealth)
	v1alpha1 := router.Group("/v1alpha1")

	v1alpha1.POST("/ovs", v1alpha1Controller.PostOvsBridge)
	v1alpha1.DELETE("/ovs/:bridge", v1alpha1Controller.DeleteOvsBridge)
	router.Run("0.0.0.0:8000")
}
