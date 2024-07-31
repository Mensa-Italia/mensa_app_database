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

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_Rzn3Oss` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `powers` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_QoWiNmc` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `expire_membership` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_tw8YWeW` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `addons` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_qqhQVTf` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `favourite_addons` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_3TNbpsC` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `is_membership_active` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_favourite_addons := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "6krufnfk",
			"name": "favourite_addons",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "30eqczpxcx155p1",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_favourite_addons); err != nil {
			return err
		}
		collection.Schema.AddField(new_favourite_addons)

		// add
		new_is_membership_active := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "xyjpad5m",
			"name": "is_membership_active",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_is_membership_active); err != nil {
			return err
		}
		collection.Schema.AddField(new_is_membership_active)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_Rzn3Oss` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `powers` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_QoWiNmc` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `expire_membership` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_tw8YWeW` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `addons` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("6krufnfk")

		// remove
		collection.Schema.RemoveField("xyjpad5m")

		return dao.SaveCollection(collection)
	})
}
