package api

import (
    "net/http"

    "github.com/labstack/echo"
)

func healthcheck(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}
