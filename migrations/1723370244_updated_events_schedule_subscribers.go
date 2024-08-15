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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("wawd580chzpiyam")
		if err != nil {
			return err
		}

		// add
		new_qt := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "hudblc9k",
			"name": "qt",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": 0,
				"max": 1000,
				"noDecimal": true
			}
		}`), new_qt); err != nil {
			return err
		}
		collection.Schema.AddField(new_qt)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("wawd580chzpiyam")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("hudblc9k")

		return dao.SaveCollection(collection)
	})
}
