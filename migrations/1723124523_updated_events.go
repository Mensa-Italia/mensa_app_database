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
			"CREATE INDEX ` + "`" + `idx_Q4ZH5fG` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `when_start` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_vhI76vh` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `contact` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_MLOSQRj` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `position` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_G4xmq3V` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `is_national` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_when_end := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "oozwx9ur",
			"name": "when_end",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_when_end); err != nil {
			return err
		}
		collection.Schema.AddField(new_when_end)

		// update
		edit_when_start := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "m8ijlzap",
			"name": "when_start",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), edit_when_start); err != nil {
			return err
		}
		collection.Schema.AddField(edit_when_start)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
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
			"CREATE INDEX ` + "`" + `idx_vhI76vh` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `contact` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_MLOSQRj` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `position` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_G4xmq3V` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `is_national` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("oozwx9ur")

		// update
		edit_when_start := &schema.SchemaField{}
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
		}`), edit_when_start); err != nil {
			return err
		}
		collection.Schema.AddField(edit_when_start)

		return dao.SaveCollection(collection)
	})
}
