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

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_kkOwDvS` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lat` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_AFNaJmf` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lon` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xUJRaTe` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `name` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_6rvfiYt` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `state` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// add
		new_state := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "t5awr8ik",
			"name": "state",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"Piemonte",
					"Valle d'Aosta",
					"Lombardia",
					"Trentino-Alto Adige",
					"Veneto",
					"Friuli-Venezia Giulia",
					"Liguria",
					"Emilia-Romagna",
					"Toscana",
					"Umbria",
					"Marche",
					"Lazio",
					"Abruzzo",
					"Molise",
					"Campania",
					"Puglia",
					"Basilicata",
					"Calabria",
					"Sicilia",
					"Sardegna"
				]
			}
		}`), new_state); err != nil {
			return err
		}
		collection.Schema.AddField(new_state)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `idx_kkOwDvS` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lat` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_AFNaJmf` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `lon` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_xUJRaTe` + "`" + ` ON ` + "`" + `positions` + "`" + ` (` + "`" + `name` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("t5awr8ik")

		return dao.SaveCollection(collection)
	})
}
