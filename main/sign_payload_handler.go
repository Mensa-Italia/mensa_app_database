package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/tidwall/gjson"
	"mensadb/tools/signatures"
	"time"
)

func SignPayloadHandler(c echo.Context) error {

	isLogged, authUser := isLoggedIn(c)
	if !isLogged {
		return apis.NewUnauthorizedError("Unauthorized", errors.New("Unauthorized"))
	}

	payloadRaw := c.FormValue("payload")
	addonsId := c.FormValue("addon")

	if payloadRaw == "" || !gjson.Valid(payloadRaw) {
		return apis.NewBadRequestError("Payload is required", errors.New("Payload is required"))
	}

	if gjson.Parse(payloadRaw).Get("email").String() != authUser.Email {
		return apis.NewBadRequestError("Invalid payload", errors.New("Invalid payload"))
	}

	var payloadJSON map[string]interface{}

	_ = json.Unmarshal([]byte(payloadRaw), &payloadJSON)

	payloadJSON["id"] = authUser.Id
	payloadJSON["email"] = authUser.Email
	payloadJSON["signature_time"] = time.Now().Format(time.RFC3339)

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
