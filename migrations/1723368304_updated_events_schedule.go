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

		collection, err := dao.FindCollectionByNameOrId("876rlfre9wpzm8i")
		if err != nil {
			return err
		}

		// add
		new_event := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rzguld92",
			"name": "event",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "yqljv1qcd98y99c",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_event); err != nil {
			return err
		}
		collection.Schema.AddField(new_event)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("876rlfre9wpzm8i")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("rzguld92")

		return dao.SaveCollection(collection)
	})
}
