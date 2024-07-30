package main

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"mensadb/tools/signatures"
)

func GeneratePublicPrivateKeys(e *core.RecordCreateEvent) error {
	keyPub, keyPriv := signatures.GenerateKeyPairs()
	e.Record.Set("public_key", keyPub)
	if err := app.Dao().Save(e.Record); err != nil {
		return err
	}

	collection, err := app.Dao().FindCollectionByNameOrId("addons_private_keys")
	if err != nil {
		return err
	}
	record := models.NewRecord(collection)
	record.Set("private_key", keyPriv)
	record.Set("addon", e.Record.Id)
	err = app.Dao().Save(record)
	return err
}
