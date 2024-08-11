package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("wawd580chzpiyam")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("user = @request.auth.id && @request.data.qt <= event_schedule.Max_external_guests")

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_MojgDSh` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `event_schedule` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_2e2dJ6L` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `user` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_eVBjVFC` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `paid` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_brps2pT` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `show` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_9d6BCOM` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_tkxAKzP` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `updated` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_byAIwvC` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (\n  ` + "`" + `event_schedule` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_cDLJO3I` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `qt` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("wawd580chzpiyam")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("user = @request.auth.id")

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_MojgDSh` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `event_schedule` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_2e2dJ6L` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `user` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_eVBjVFC` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `paid` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_brps2pT` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `show` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_9d6BCOM` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_tkxAKzP` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (` + "`" + `updated` + "`" + `)",
			"CREATE UNIQUE INDEX ` + "`" + `idx_byAIwvC` + "`" + ` ON ` + "`" + `events_schedule_subscribers` + "`" + ` (\n  ` + "`" + `event_schedule` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	})
}
