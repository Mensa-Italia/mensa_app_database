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

		collection, err := dao.FindCollectionByNameOrId("ejk4za0j9k23dyf")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("9asnsiz8")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("ejk4za0j9k23dyf")
		if err != nil {
			return err
		}

		// add
		del_field := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "9asnsiz8",
			"name": "field",
			"type": "email",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), del_field); err != nil {
			return err
		}
		collection.Schema.AddField(del_field)

		return dao.SaveCollection(collection)
	})
}
