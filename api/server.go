package api

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func Server() *echo.Echo {
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  return e
}

func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
