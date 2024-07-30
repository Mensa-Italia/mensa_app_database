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
			"id": "yqljv1qcd98y99c",
			"created": "2024-07-30 10:50:51.397Z",
			"updated": "2024-07-30 10:50:51.397Z",
			"name": "events",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "fdegr8qg",
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
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_5hThibd` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `name` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("yqljv1qcd98y99c")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
