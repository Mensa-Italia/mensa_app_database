package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"mensadb/tools/signatures"
	"slices"
	"time"
)

func SignPayloadHandler(c echo.Context) error {

	isLogged, authUser := isLoggedIn(c)
	if !isLogged {
		return apis.NewUnauthorizedError("Unauthorized", errors.New("Unauthorized"))
	}

	addonsId := c.FormValue("addon")

	user, err := app.Dao().FindRecordById("users", authUser.Id)
	if err != nil {
		return err
	}

	payloadJSON := map[string]interface{}{
		"id":                user.Get("id"),
		"email":             user.Get("email"),
		"name":              user.Get("name"),
		"avatar":            "https://svc.mensa.it/api/files/_pb_users_auth_/" + user.Get("id").(string) + "/" + user.Get("avatar").(string),
		"powers":            user.Get("powers"),
		"expire_membership": user.Get("expire_membership"),
		"signed_at":         time.Now().Format(time.RFC3339),
	}

	if !slices.Contains(user.GetStringSlice("addons"), addonsId) {
		user.Set("addons", append(user.GetStringSlice("addons"), addonsId))
		err = app.Dao().Save(user)
		if err != nil {
			return err
		}
	}

	payload, _ := json.Marshal(payloadJSON)

	record, err := app.Dao().FindFirstRecordByData("addons_private_keys", "addon", addonsId)
	if err != nil {
		return apis.NewBadRequestError("Invalid addon", err)
	}

	payloadBase64 := payloadToBase64(string(payload))
	signature, err := signatures.SignData([]byte(payloadBase64), record.Get("private_key").(string))

	if err != nil {
		return apis.NewBadRequestError("Failed to sign payload", err)
	}

	return c.JSON(200, map[string]interface{}{
		"payload":   payloadBase64,
		"signature": signature,
	})
}

func payloadToBase64(payload string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(payload))
}
