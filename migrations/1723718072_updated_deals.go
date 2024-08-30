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
			"CREATE INDEX `+"`"+`idx_otedWnF`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`details`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("30om98io")

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
			"CREATE INDEX `+"`"+`idx_xdlmhxV`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`state`+"`"+`)",
			"CREATE INDEX `+"`"+`idx_otedWnF`+"`"+` ON `+"`"+`deals`+"`"+` (`+"`"+`details`+"`"+`)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		del_state := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "30om98io",
			"name": "state",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"piemonte"
				]
			}
		}`), del_state); err != nil {
			return err
		}
		collection.Schema.AddField(del_state)

		return dao.SaveCollection(collection)
	})
}
