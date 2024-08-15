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
			"id": "k4vigrrbibw2hu8",
			"created": "2024-08-15 10:33:44.721Z",
			"updated": "2024-08-15 10:33:44.721Z",
			"name": "deals",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "ir4m6jhk",
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
					"id": "nkppvdkd",
					"name": "commercial_sector",
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
					"id": "hpea5go1",
					"name": "position",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "48avh9pvrlw8w7n",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "l2vj1w0j",
					"name": "is_local",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "30om98io",
					"name": "state",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"piemonte"
						]
					}
				},
				{
					"system": false,
					"id": "n5qpzqba",
					"name": "details",
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
					"id": "3yirfjos",
					"name": "who",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"active_members",
							"active_members and relatives"
						]
					}
				},
				{
					"system": false,
					"id": "hlwrfjoi",
					"name": "starting",
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
					"id": "zhjyd6dk",
					"name": "ending",
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
					"id": "bi9e4paj",
					"name": "how_to_get",
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
					"id": "idtnh1le",
					"name": "link",
					"type": "url",
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
				"CREATE INDEX ` + "`" + `idx_BjbROyY` + "`" + ` ON ` + "`" + `deals` + "`" + ` (` + "`" + `name` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_BAqtECx` + "`" + ` ON ` + "`" + `deals` + "`" + ` (` + "`" + `commercial_sector` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_p1cWNyw` + "`" + ` ON ` + "`" + `deals` + "`" + ` (` + "`" + `position` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_iGJuxMB` + "`" + ` ON ` + "`" + `deals` + "`" + ` (` + "`" + `is_local` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_xdlmhxV` + "`" + ` ON ` + "`" + `deals` + "`" + ` (` + "`" + `state` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_otedWnF` + "`" + ` ON ` + "`" + `deals` + "`" + ` (` + "`" + `details` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("k4vigrrbibw2hu8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
