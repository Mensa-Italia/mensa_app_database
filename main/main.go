package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/tidwall/gjson"
	"log"
	_ "mensadb/migrations"
	"mensadb/tools/signatures"
	"os"
	"time"
)

var app = pocketbase.New()

func main() {

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		scheduler := cron.New()

		// Update addons data every day at 3:01
		scheduler.MustAdd("updateAddonsData", "1 3 * * *", func() {
			go updateAddonsData()
		})

		scheduler.Start()

		return nil
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         "./migrations",
		Automigrate: true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.POST("/api/cs/auth-with-area", AuthWithAreaHandler)
		e.Router.GET("/api/cs/sign-payload/:addon", SignPayloadHandler)
		e.Router.GET("/api/cs/keys/:addon", GetAddonPublicKeysHandler)
		e.Router.POST("/api/cs/verify-signature/:addon", VerifySignatureHandler)
		e.Router.GET("/api/cs/force-update-addons", ForceUpdateAddonsHandler)
		e.Router.GET("/ical/:hash", RetrieveICAL)
		e.Router.GET("/static/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		return nil
	})
	app.OnRecordAfterCreateRequest("addons").Add(GeneratePublicPrivateKeys)
	app.OnRecordAfterCreateRequest("positions").Add(PositionSetState)
	app.OnRecordAfterCreateRequest("calendar_link").Add(CalendarSetHash)

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {

		expr, err := app.Dao().FindRecordsByExpr("users")
		if err != nil {
			return nil
		}

		if len(expr) == 0 {
			return nil
		}
		for _, record := range expr {
			if record == nil {
				continue
			}

			calendarLinkCollection, _ := app.Dao().FindCollectionByNameOrId("calendar_link")
			newCalendar := models.NewRecord(calendarLinkCollection)
			newCalendar.Set("user", record.Id)
			newCalendar.Set("hash", randomHash())
			app.Dao().SaveRecord(newCalendar)
		}
		return nil

	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func GetAddonPublicKeysHandler(c echo.Context) error {
	addon := c.PathParam("addon")
	record, err := app.Dao().FindRecordById("addons", addon)
	if err != nil {
		return apis.NewBadRequestError("Invalid addon", err)
	}

	return c.String(200, record.Get("public_key").(string))
}

func VerifySignatureHandler(c echo.Context) error {
	addonId := c.PathParam("addon")
	signature := c.FormValue("signature")
	payload := c.FormValue("payload")

	record, err := app.Dao().FindRecordById("addons", addonId)
	if err != nil {
		return apis.NewBadRequestError("Invalid addon", err)
	}

	isValid := signatures.ValidateSignature(payload, signature, record.Get("public_key").(string))

	payloadPure := payloadFromBase64(payload)

	if !gjson.ValidBytes([]byte(payloadPure)) {
		return apis.NewBadRequestError("Invalid payload", nil)
	}

	dataToUse := gjson.ParseBytes([]byte(payloadPure))

	if dataToUse.Get("expires_at").Time().After(time.Now()) &&
		dataToUse.Get("addon_id").String() == addonId &&
		isValid {
		return c.String(200, "OK")
	}
	return apis.NewBadRequestError("Invalid signature", nil)

}
