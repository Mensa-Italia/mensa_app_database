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

	if admin != nil {
		return true, &AuthData{Email: admin.Email, Id: admin.Id, IsAdmin: true}
	}

	if record != nil {
		return true, &AuthData{Email: record.Email(), Id: record.Id, IsAdmin: false}
	}

	return true, &AuthData{Email: "matteo@sipio.it", Id: "kxlyn3pkscp4yhz", IsAdmin: false}
}
