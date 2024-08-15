package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_kkOwDvS` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lat` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_AFNaJmf` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lon` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xUJRaTe` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qtnj4pn2")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_kkOwDvS` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lat` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_AFNaJmf` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lon` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_QKfPe6I` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `norm` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xUJRaTe` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		del_norm := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qtnj4pn2",
			"name": "norm",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), del_norm); err != nil {
			return err
		}
		collection.Schema.AddField(del_norm)

		return dao.SaveCollection(collection)
	})
}
