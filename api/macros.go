package api

import (
    "net/http"
    "net/url"
    "io/ioutil"

    "github.com/labstack/echo"
    "github.com/cristianossd/unbearable-eater/config"
)

func getCalories(c echo.Context) error {
    var err error

    resp, err := http.PostForm(config.CaloriesDataEndpoint,
        url.Values{"keyword": {"cassava"}, "page": {"0"}, "ppage": {"10"}})

    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }


    return c.String(http.StatusOK, string(body))
}
