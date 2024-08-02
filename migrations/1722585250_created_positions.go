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
			"id": "48avh9pvrlw8w7n",
			"created": "2024-08-02 07:54:10.971Z",
			"updated": "2024-08-02 07:54:10.971Z",
			"name": "positions",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "cjrpakf1",
					"name": "lat",
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
					"id": "t1yk87hd",
					"name": "lon",
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
					"id": "qtnj4pn2",
					"name": "norm",
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
					"id": "0qw0kv7w",
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
				"CREATE INDEX ` + "`" + `idx_kkOwDvS` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lat` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_AFNaJmf` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lon` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_QKfPe6I` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `norm` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_xUJRaTe` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `name` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
