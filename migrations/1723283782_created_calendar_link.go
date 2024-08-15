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
			"id": "nq7uqdrz46xryck",
			"created": "2024-08-10 09:56:22.496Z",
			"updated": "2024-08-10 09:56:22.496Z",
			"name": "calendar_link",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "riaimzz2",
					"name": "user",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "xczyzuz9",
					"name": "hash",
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
				"CREATE UNIQUE INDEX ` + "`" + `idx_guJTLvu` + "`" + ` ON ` + "`" + `calendar_link` + "`" + ` (` + "`" + `user` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_5PY6PEw` + "`" + ` ON ` + "`" + `calendar_link` + "`" + ` (` + "`" + `hash` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("nq7uqdrz46xryck")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
