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

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		// add
		new_is_ready := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "teb13vf6",
			"name": "is_ready",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_is_ready); err != nil {
			return err
		}
		collection.Schema.AddField(new_is_ready)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("teb13vf6")

		return dao.SaveCollection(collection)
	})
}
