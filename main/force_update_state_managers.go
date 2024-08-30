package main

import "github.com/labstack/echo/v5"

func ForceUpdateStateManagersHandler(c echo.Context) error {
	go updateStateManagers()
	return c.String(200, "OK")
}
