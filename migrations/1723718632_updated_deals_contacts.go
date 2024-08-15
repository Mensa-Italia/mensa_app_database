package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("cl4zw57ht0zo7xk")
		if err != nil {
			return err
		}

		// add
		new_note := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "gwnnrjw3",
			"name": "note",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_note); err != nil {
			return err
		}
		collection.Schema.AddField(new_note)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("cl4zw57ht0zo7xk")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("gwnnrjw3")

		return dao.SaveCollection(collection)
	})
}
