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

func (h *ProjectHandler) AttempAdd(c *gin.Context) {
    c.MultipartForm();
    c.JSON(http.StatusOK, c.Request.PostForm)
}
