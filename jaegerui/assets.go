package jaegerui

import (
	"github.com/labstack/echo/v4"
	"github.com/c3sr/web/jaegerui/static"
)

func Handle() echo.HandlerFunc {
	return echo.WrapHandler(static.Handler)
}
