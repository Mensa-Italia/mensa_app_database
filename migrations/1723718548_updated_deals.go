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
			"CREATE INDEX `+"`"+`idx_7t5Me6C`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`updated`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_JqPonvH`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`owner`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_1jswY5L`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`is_active`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_is_active := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "mczqkdl6",
			"name": "is_active",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_is_active); err != nil {
			return err
		}
		collection.Schema.AddField(new_is_active)

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
			"CREATE INDEX `+"`"+`idx_otedWnF`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`details`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_E4wsq6f`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`who`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_iCh99qb`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`starting`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_1fWt8bL`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`ending`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_dLphsZC`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`created`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_7t5Me6C`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`updated`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_JqPonvH`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`owner`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("mczqkdl6")

		return dao.SaveCollection(collection)
	})
}
