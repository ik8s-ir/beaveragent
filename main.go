package main

import (
	"github.com/gin-gonic/gin"
	v1alpha1Controller "github.com/ik8s-ir/beaveragent/pkg/controllers/v1alpha1"
)

func main() {
	router := gin.Default()
	v1alpha1 := router.Group("/apis/networking.ik8s.ir/v1alpha1")
	v1alpha1.POST("/ovs", v1alpha1Controller.PostOvsBridge)
	router.Run("localhost:8000")
}
