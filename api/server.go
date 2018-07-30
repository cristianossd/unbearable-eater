package api

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func Server() *echo.Echo {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/health", healthcheck)

    // Crawler api
    crawlerApiGroup := e.Group("")
    crawlerApiGroup.GET("/macros", getMacros)

    return e
}
