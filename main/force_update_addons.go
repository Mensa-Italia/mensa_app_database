package main

import "github.com/labstack/echo/v5"

func ForceUpdateAddonsHandler(c echo.Context) error {
	go updateAddonsData()
	return c.String(200, "OK")
}
