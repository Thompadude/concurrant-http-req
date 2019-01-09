package main

import (
	"os"

	"github.com/Thompadude/concurrant-http-req/handler"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(bind())
	log.EnableColor()
	log.SetHeader(`${level} ::: ${time_rfc3339} ::: line: ${line} ::: message:`)

	e.GET("/concurrent", handler.Concurrent)

	e.Logger.Fatal(e.Start(":1323"))
}

func bind() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("url_count", len(os.Args)-1)
			return next(c)
		}
	}
}
