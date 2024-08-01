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

		collection, err := dao.FindCollectionByNameOrId("yqljv1qcd98y99c")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_5hThibd` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `name` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_aaLanjA` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `name` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_rKOwHIk` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `description` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_0RxwrTy` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `info_link` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_eH7D9Nl` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `booking_link` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Q4ZH5fG` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `when` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_vhI76vh` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `contact` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_image := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "zdkzfcau",
			"name": "image",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), new_image); err != nil {
			return err
		}
		collection.Schema.AddField(new_image)

		// add
		new_description := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ndbpzusb",
			"name": "description",
			"type": "editor",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"convertUrls": false
			}
		}`), new_description); err != nil {
			return err
		}
		collection.Schema.AddField(new_description)

		// add
		new_info_link := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "4ll0pbry",
			"name": "info_link",
			"type": "url",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), new_info_link); err != nil {
			return err
		}
		collection.Schema.AddField(new_info_link)

		// add
		new_booking_link := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "814aczs4",
			"name": "booking_link",
			"type": "url",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), new_booking_link); err != nil {
			return err
		}
		collection.Schema.AddField(new_booking_link)

		// add
		new_when := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "m8ijlzap",
			"name": "when",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_when); err != nil {
			return err
		}
		collection.Schema.AddField(new_when)

		// add
		new_contact := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "sp0hrxip",
			"name": "contact",
			"type": "email",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"exceptDomains": null,
				"onlyDomains": null
			}
		}`), new_contact); err != nil {
			return err
		}
		collection.Schema.AddField(new_contact)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yqljv1qcd98y99c")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_5hThibd` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("zdkzfcau")

		// remove
		collection.Schema.RemoveField("ndbpzusb")

		// remove
		collection.Schema.RemoveField("4ll0pbry")

		// remove
		collection.Schema.RemoveField("814aczs4")

		// remove
		collection.Schema.RemoveField("m8ijlzap")

		// remove
		collection.Schema.RemoveField("sp0hrxip")

		return dao.SaveCollection(collection)
	})
}
