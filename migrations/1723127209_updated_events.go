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

		// remove
		collection.Schema.RemoveField("sqm6vtw7")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yqljv1qcd98y99c")
		if err != nil {
			return err
		}

		// add
		del_italian_state := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "sqm6vtw7",
			"name": "italian_state",
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
		}`), del_italian_state); err != nil {
			return err
		}
		collection.Schema.AddField(del_italian_state)

		return dao.SaveCollection(collection)
	})
}
