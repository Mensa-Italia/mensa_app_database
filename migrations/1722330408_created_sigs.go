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
			"id": "ejk4za0j9k23dyf",
			"created": "2024-07-30 09:06:48.362Z",
			"updated": "2024-07-30 09:06:48.362Z",
			"name": "sigs",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "sflaodu3",
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
					"id": "fvhtnl3r",
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
					"id": "wrgi19wz",
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
					"id": "e6kexrr6",
					"name": "link",
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
					"id": "9asnsiz8",
					"name": "field",
					"type": "email",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": null,
						"onlyDomains": null
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_Gv4aZDk` + "`" + ` ON ` + "`" + `sigs` + "`" + ` (` + "`" + `name` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_iRXlnIz` + "`" + ` ON ` + "`" + `sigs` + "`" + ` (` + "`" + `description` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_jou21SC` + "`" + ` ON ` + "`" + `sigs` + "`" + ` (` + "`" + `image` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_6ulAfUT` + "`" + ` ON ` + "`" + `sigs` + "`" + ` (` + "`" + `link` + "`" + `)"
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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("ejk4za0j9k23dyf")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
