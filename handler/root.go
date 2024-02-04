package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mattcollie/mattcollecutt.com/view/root"
)

type RootHandler struct{}

func (h RootHandler) HandleRootShow(c echo.Context) error {
	return render(c, root.Show())
}
