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
			"id": "pjbwpa2ue52awg8",
			"created": "2024-08-15 19:03:47.301Z",
			"updated": "2024-08-15 19:03:47.301Z",
			"name": "users_metadata",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "acwozftv",
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
					"id": "owfshcso",
					"name": "key",
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
					"id": "vicrebwq",
					"name": "value",
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
				"CREATE INDEX ` + "`" + `idx_V4lmL65` + "`" + ` ON ` + "`" + `users_metadata` + "`" + ` (` + "`" + `user` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_U2b90dZ` + "`" + ` ON ` + "`" + `users_metadata` + "`" + ` (` + "`" + `key` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_giF62fh` + "`" + ` ON ` + "`" + `users_metadata` + "`" + ` (` + "`" + `value` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_pjhcCSh` + "`" + ` ON ` + "`" + `users_metadata` + "`" + ` (` + "`" + `created` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_PVaY2DC` + "`" + ` ON ` + "`" + `users_metadata` + "`" + ` (` + "`" + `updated` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_WjVexnv` + "`" + ` ON ` + "`" + `users_metadata` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `key` + "`" + `\n)"
			],
			"listRule": "@request.auth.id = user.id",
			"viewRule": "@request.auth.id = user.id",
			"createRule": "@request.auth.id = user.id",
			"updateRule": "@request.auth.id = user.id",
			"deleteRule": "@request.auth.id = user.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("pjbwpa2ue52awg8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
