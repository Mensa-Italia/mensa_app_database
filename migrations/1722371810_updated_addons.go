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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_mcYdJRX`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`name`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_KbKUrae`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`description`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_thQVhad`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`created`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_YlTarP6`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`updated`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_tyTwDlK`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`version`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_f3WSdbm`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`url`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_version := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "eyk7anno",
			"name": "version",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": "^(\\d+\\.)?(\\d+\\.)?(\\*|\\d+)$"
			}
		}`), new_version); err != nil {
			return err
		}
		collection.Schema.AddField(new_version)

		// update
		edit_url := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ybyraffl",
			"name": "url",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_url); err != nil {
			return err
		}
		collection.Schema.AddField(edit_url)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_mcYdJRX`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`name`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_KbKUrae`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`description`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_MI7KtPH`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`metadata_url`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_XxQ0YCD`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`public_key`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_thQVhad`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`created`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_YlTarP6`+"`"+` ON `+"`"+`addons`+"`"+` (`+"`"+`updated`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("eyk7anno")

		// update
		edit_url := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ybyraffl",
			"name": "metadata_url",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_url); err != nil {
			return err
		}
		collection.Schema.AddField(edit_url)

		return dao.SaveCollection(collection)
	})
}
