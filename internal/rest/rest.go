package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Action func(c *gin.Context)

type ActionMap map[string]Action

type ResourceMap map[string]ActionMap

var resources ResourceMap

func notFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found"})
}

func createAction(c *gin.Context) {
	resource, exists := resources[c.Param("resource")]
	if exists == false {
		notFound(c)
		return
	}
	action, exists := resource["create"]
	if exists == false {
		notFound(c)
		return
	}
	action(c)
}

func updateAction(c *gin.Context) {
	resource, exists := resources[c.Param("resource")]
	if exists == false {
		notFound(c)
		return
	}
	action, exists := resource["update"]
	if exists == false {
		notFound(c)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		notFound(c)
		return
	}
	c.Set("ID", id)
	action(c)
}

func detailAction(c *gin.Context) {
	resource, exists := resources[c.Param("resource")]
	if exists == false {
		notFound(c)
		return
	}
	action, exists := resource["detail"]
	if exists == false {
		notFound(c)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		notFound(c)
		return
	}
	c.Set("ID", id)
	action(c)
}

func listAction(c *gin.Context) {
	resource, exists := resources[c.Param("resource")]
	if exists == false {
		notFound(c)
		return
	}
	action, exists := resource["list"]
	if exists == false {
		notFound(c)
		return
	}
	action(c)
}

func listChildAction(c *gin.Context) {
	resource, exists := resources[c.Param("resource")]
	if exists == false {
		notFound(c)
		return
	}
	child := c.Param("child")
	action, exists := resource["*/"+child]
	if exists == false {
		notFound(c)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		notFound(c)
		return
	}
	c.Set("ID", id)
	action(c)
}

func deleteAction(c *gin.Context) {
	resource, exists := resources[c.Param("resource")]
	if exists == false {
		notFound(c)
		return
	}
	action, exists := resource["delete"]
	if exists == false {
		notFound(c)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		notFound(c)
		return
	}
	c.Set("ID", id)
	action(c)
}

// See https://www.openmymind.net/RESTful-routing-in-Go/
func AttachEndpoints(resourceMap ResourceMap, r *gin.Engine) {
	resources = resourceMap
	r.POST("/:resource", createAction)
	r.GET("/:resource", listAction)
	r.GET("/:resource/:id", detailAction)
	r.GET("/:resource/:id/:child", listChildAction)
	r.DELETE("/:resource/:id", deleteAction)
	r.PUT("/:resource/:id", updateAction)
}
