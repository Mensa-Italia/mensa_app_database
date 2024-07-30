package main

import "github.com/labstack/echo/v5"

func ForceUpdateAddonsHandler(c echo.Context) error {
	updateAddonsData()
	return c.String(200, "OK")
}
