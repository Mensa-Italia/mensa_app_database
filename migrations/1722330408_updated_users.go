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

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"allowEmailAuth": false,
			"allowOAuth2Auth": false,
			"allowUsernameAuth": false,
			"exceptEmailDomains": null,
			"manageRule": null,
			"minPasswordLength": 5,
			"onlyEmailDomains": null,
			"onlyVerified": true,
			"requireEmail": false
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_Rzn3Oss`+"`"+` ON `+"`"+`users`+"`"+` (`+"`"+`powers`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_QoWiNmc`+"`"+` ON `+"`"+`users`+"`"+` (`+"`"+`expire_membership`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_powers := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "s3iw8haw",
			"name": "powers",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 3,
				"values": [
					"sigs",
					"events",
					"addons"
				]
			}
		}`), new_powers); err != nil {
			return err
		}
		collection.Schema.AddField(new_powers)

		// add
		new_expire_membership := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "srjzwutc",
			"name": "expire_membership",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_expire_membership); err != nil {
			return err
		}
		collection.Schema.AddField(new_expire_membership)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": true,
			"exceptEmailDomains": null,
			"manageRule": null,
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"onlyVerified": false,
			"requireEmail": false
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		if err := json.Unmarshal([]byte(`[]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("s3iw8haw")

		// remove
		collection.Schema.RemoveField("srjzwutc")

		return dao.SaveCollection(collection)
	})
}
