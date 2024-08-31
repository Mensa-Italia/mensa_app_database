package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("jd6v6egqp096rer")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_F0gqJyw` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `device_name` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_n7ZmHEL` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `firebase_id` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_5GzkmFX` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_GjoEt1B` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `updated` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_FvsAfc8` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `firebase_id` + "`" + `\n)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("jd6v6egqp096rer")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_F0gqJyw` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `device_name` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_n7ZmHEL` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `firebase_id` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_5GzkmFX` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_GjoEt1B` + "`" + ` ON ` + "`" + `users_devices` + "`" + ` (` + "`" + `updated` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	})
}
