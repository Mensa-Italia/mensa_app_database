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
			"id": "876rlfre9wpzm8i",
			"created": "2024-08-11 09:24:26.894Z",
			"updated": "2024-08-11 09:24:26.894Z",
			"name": "events_schedule",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ruvt0t0s",
					"name": "title",
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
					"id": "rkgfi8ka",
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
					"id": "kdo2qi6k",
					"name": "image",
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
					"id": "gwglkxpm",
					"name": "when_start",
					"type": "date",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "7molze7d",
					"name": "when_end",
					"type": "date",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "qm0hmdco",
					"name": "price",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "ojnqyka8",
					"name": "info_link",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": null,
						"onlyDomains": null
					}
				},
				{
					"system": false,
					"id": "elkde680",
					"name": "is_subscriptable",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_TZLPzUO` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `title` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_T1MZGUE` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `description` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_Aoa7EFF` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `image` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_LgXuS1r` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `price` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_UZELKrp` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `info_link` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_uc5QLck` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `is_subscriptable` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_8JAJ8tS` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_8mzulbq` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `updated` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("876rlfre9wpzm8i")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
