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
			"id": "cl4zw57ht0zo7xk",
			"created": "2024-08-15 10:40:07.580Z",
			"updated": "2024-08-15 10:40:07.580Z",
			"name": "deals_contacts",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jipq5asj",
					"name": "deal",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "k4vigrrbibw2hu8",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "nwjkhgpz",
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
					"id": "j41jdvn3",
					"name": "email",
					"type": "email",
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
					"id": "zsh0lek5",
					"name": "phone_number",
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
				"CREATE INDEX ` + "`" + `idx_qD5AucX` + "`" + ` ON ` + "`" + `deals_contacts` + "`" + ` (` + "`" + `name` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_mOHO1dA` + "`" + ` ON ` + "`" + `deals_contacts` + "`" + ` (` + "`" + `deal` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_EVN6hrg` + "`" + ` ON ` + "`" + `deals_contacts` + "`" + ` (` + "`" + `email` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_ABp4FIk` + "`" + ` ON ` + "`" + `deals_contacts` + "`" + ` (` + "`" + `phone_number` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_7uNPaod` + "`" + ` ON ` + "`" + `deals_contacts` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_qBa8G5K` + "`" + ` ON ` + "`" + `deals_contacts` + "`" + ` (` + "`" + `updated` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("cl4zw57ht0zo7xk")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
