package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mattcollie/mattcollecutt.com/view/root"
	"net/http"
)

type RootHandler struct{}

func (h RootHandler) HandleRootShow(c *gin.Context) {
	c.HTML(http.StatusOK, "", root.Show())
}
