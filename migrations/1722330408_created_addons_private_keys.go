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
			"id": "2yn1ow9epkq83ku",
			"created": "2024-07-30 09:06:48.362Z",
			"updated": "2024-07-30 09:06:48.362Z",
			"name": "addons_private_keys",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "hiflm6v6",
					"name": "addon",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "30eqczpxcx155p1",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "aytiqimv",
					"name": "private_key",
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
				"CREATE INDEX ` + "`" + `idx_F1OP45M` + "`" + ` ON ` + "`" + `addons_private_keys` + "`" + ` (` + "`" + `addon` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("2yn1ow9epkq83ku")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
