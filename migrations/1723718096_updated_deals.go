package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("k4vigrrbibw2hu8")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_BjbROyY`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`name`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_BAqtECx`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`commercial_sector`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_p1cWNyw`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`position`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_iGJuxMB`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`is_local`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_otedWnF`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`details`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_E4wsq6f`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`who`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_iCh99qb`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`starting`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_1fWt8bL`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`ending`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_dLphsZC`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`created`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_7t5Me6C`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`updated`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("k4vigrrbibw2hu8")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX `+"`"+`idx_BjbROyY`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`name`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_BAqtECx`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`commercial_sector`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_p1cWNyw`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`position`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_iGJuxMB`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`is_local`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_otedWnF`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`details`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	})
}
