package handler

import (
	"apallis/portfolio/renderer"
	"apallis/portfolio/view/pages"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {}

func (h *ImageHandler) Show(c *gin.Context) {
	render := renderer.New(c, http.StatusOK, pages.ShowImages())
	c.Render(http.StatusOK, render)
	return
}

func (h *ImageHandler) Upload(c *gin.Context) {
    file, err := c.FormFile("image")
    if err != nil {
        log.Println(err)
        c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
        return
    }
    log.Println(file.Filename)
    c.SaveUploadedFile(file, "writable/images/" + file.Filename)
    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
