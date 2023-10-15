package pages

import (
	"github.com/labstack/echo/v4"
)

type IndexPage struct {
}

func Index(c echo.Context) error {
	return c.Render(200, "index.html", IndexPage{})
}
