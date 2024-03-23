package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salehrashid/portfolio-saleh/internal/renderer"
	util "github.com/salehrashid/portfolio-saleh/internal/util"
	"net/http"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewTemplateRenderer(true)
	e.Static("/assets", util.GetString("WEB_PORTFOLIO_FILE_STATICS_PATH", "../../../template/statics"))

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.Logger.Fatal(e.Start(util.GetPort("DOCKER_WEB_PORTFOLIO_HOST_PORT")))
}
