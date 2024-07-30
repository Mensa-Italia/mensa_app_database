package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"log"
	_ "mensadb/migrations"
	"os"
)

var app = pocketbase.New()

func main() {

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         "./migrations",
		Automigrate: true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/auth-with-area", AuthWithAreaHandler)
		e.Router.POST("/sign-payload", SignPayloadHandler)
		e.Router.GET("/keys/:addon", GetAddonPublicKeysHandler)
		e.Router.GET("/static/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})
	app.OnRecordAfterCreateRequest("addons").Add(GeneratePublicPrivateKeys)
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
