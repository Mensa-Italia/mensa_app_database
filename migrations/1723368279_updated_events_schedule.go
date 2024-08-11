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

		collection, err := dao.FindCollectionByNameOrId("876rlfre9wpzm8i")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_TZLPzUO` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `title` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_T1MZGUE` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `description` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Aoa7EFF` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `image` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_LgXuS1r` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `price` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_UZELKrp` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `info_link` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_uc5QLck` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `is_subscriptable` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_8JAJ8tS` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_8mzulbq` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `updated` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_LaatiDd` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `when_start` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_t7GOSIX` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `when_end` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("876rlfre9wpzm8i")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_TZLPzUO` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `title` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_T1MZGUE` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `description` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_Aoa7EFF` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `image` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_LgXuS1r` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `price` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_UZELKrp` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `info_link` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_uc5QLck` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `is_subscriptable` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_8JAJ8tS` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_8mzulbq` + "`" + ` ON ` + "`" + `events_schedule` + "`" + ` (` + "`" + `updated` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	})
}
