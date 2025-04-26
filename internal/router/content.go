package router

import (
	"promotion/internal/controller"
	"promotion/internal/middleware"

	"github.com/gin-gonic/gin"
)

type contentRouter struct {
	group        *gin.RouterGroup
	controller   *controller.ContentController
	internalAuth *middleware.InternalAuthMiddleware
}

func initContentRouter(
	group *gin.RouterGroup,
	c *controller.ContentController,
	internalAuth *middleware.InternalAuthMiddleware,
) {
	router := newContentRouter(group, c, internalAuth)
	router.handle()
}

func newContentRouter(
	group *gin.RouterGroup,
	c *controller.ContentController,
	internalAuth *middleware.InternalAuthMiddleware,
) *contentRouter {
	return &contentRouter{group, c, internalAuth}
}

func (r contentRouter) handle() {
	root := r.group.Group("/content")
	root.POST("", r.controller.GenerateLessonFromCode)
}
