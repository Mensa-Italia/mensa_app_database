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
			"id": "jd6v6egqp096rer",
			"created": "2024-08-15 19:00:58.694Z",
			"updated": "2024-08-15 19:00:58.694Z",
			"name": "users_devices",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "r3xs7tpz",
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
					"id": "sdiezkbb",
					"name": "device_name",
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
					"id": "rfgcvoln",
					"name": "firebase_id",
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
				"CREATE INDEX ` + "`" + `idx_F0gqJyw` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `device_name` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_n7ZmHEL` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `firebase_id` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_5GzkmFX` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_GjoEt1B` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `updated` + "`" + `)"
			],
			"listRule": "@request.auth.id = user.id",
			"viewRule": "@request.auth.id = user.id",
			"createRule": "@request.auth.id = user.id",
			"updateRule": "@request.auth.id = user.id && @request.data.firebase_id = firebase_id && @request.data.device_name = device_name && @request.data.created = created",
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

		collection, err := dao.FindCollectionByNameOrId("jd6v6egqp096rer")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
