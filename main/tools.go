package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
)

type AuthData struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func isLoggedIn(c echo.Context) (bool, *AuthData) {
	info := apis.RequestInfo(c)
	admin := info.Admin
	record := info.AuthRecord

	if record != nil {
		return true, &AuthData{Email: record.Email(), Id: record.Id, IsAdmin: false}
	}

	if admin != nil {
		return false, nil
	}

	return false, nil
}
