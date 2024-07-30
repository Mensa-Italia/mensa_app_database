package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "30eqczpxcx155p1",
			"created": "2024-07-30 09:06:48.362Z",
			"updated": "2024-07-30 09:06:48.362Z",
			"name": "addons",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "sjllgym3",
					"name": "name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "9nfgylol",
					"name": "description",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
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
				},
				{
					"system": false,
					"id": "e4dihetm",
					"name": "icon",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [],
						"thumbs": [],
						"maxSelect": 1,
						"maxSize": 5242880,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "mfbpmcsm",
					"name": "public_key",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_mcYdJRX` + "`" + ` ON ` + "`" + `addons` + "`" + ` (` + "`" + `name` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_KbKUrae` + "`" + ` ON ` + "`" + `addons` + "`" + ` (` + "`" + `description` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_MI7KtPH` + "`" + ` ON ` + "`" + `addons` + "`" + ` (` + "`" + `metadata_url` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_XxQ0YCD` + "`" + ` ON ` + "`" + `addons` + "`" + ` (` + "`" + `public_key` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_thQVhad` + "`" + ` ON ` + "`" + `addons` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_YlTarP6` + "`" + ` ON ` + "`" + `addons` + "`" + ` (` + "`" + `updated` + "`" + `)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
