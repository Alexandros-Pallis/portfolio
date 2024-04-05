package handler

import (
	"apallis/portfolio/renderer"
	"apallis/portfolio/view/pages"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {}

func (h *ProjectHandler) Add(c *gin.Context) {
    render := renderer.New(c, http.StatusOK, pages.AddProject())
    c.Render(http.StatusOK, render)
}
