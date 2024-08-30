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

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update
		edit_powers := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "s3iw8haw",
			"name": "powers",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 5,
				"values": [
					"sigs",
					"events",
					"addons",
					"testmakers",
					"deals",
					"super"
				]
			}
		}`), edit_powers); err != nil {
			return err
		}
		collection.Schema.AddField(edit_powers)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update
		edit_powers := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "s3iw8haw",
			"name": "powers",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 5,
				"values": [
					"sigs",
					"events",
					"addons",
					"testmakers",
					"super"
				]
			}
		}`), edit_powers); err != nil {
			return err
		}
		collection.Schema.AddField(edit_powers)

		return dao.SaveCollection(collection)
	})
}
