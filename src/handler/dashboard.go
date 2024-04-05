package handler

import (
	"apallis/portfolio/renderer"
	"apallis/portfolio/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct{}

func (h *DashboardHandler) Show(c *gin.Context) {
	r := renderer.New(c, http.StatusOK, view.Dashboard())
	c.Render(http.StatusOK, r)
}
