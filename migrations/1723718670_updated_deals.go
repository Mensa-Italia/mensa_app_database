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
		new_vat_number := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "8sflozq2",
			"name": "vat_number",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_vat_number); err != nil {
			return err
		}
		collection.Schema.AddField(new_vat_number)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("k4vigrrbibw2hu8")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("8sflozq2")

		return dao.SaveCollection(collection)
	})
}
