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

		collection, err := dao.FindCollectionByNameOrId("k4vigrrbibw2hu8")
		if err != nil {
			return err
		}

		// add
		new_attachment := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "d9c5s2e7",
			"name": "attachment",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"application/pdf"
				],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": true
			}
		}`), new_attachment); err != nil {
			return err
		}
		collection.Schema.AddField(new_attachment)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("k4vigrrbibw2hu8")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("d9c5s2e7")

		return dao.SaveCollection(collection)
	})
}
