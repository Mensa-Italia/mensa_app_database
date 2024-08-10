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

		collection, err := dao.FindCollectionByNameOrId("nq7uqdrz46xryck")
		if err != nil {
			return err
		}

		// add
		new_state := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "s3pydfb2",
			"name": "state",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 21,
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
					"Sardegna",
					"Online"
				]
			}
		}`), new_state); err != nil {
			return err
		}
		collection.Schema.AddField(new_state)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nq7uqdrz46xryck")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("s3pydfb2")

		return dao.SaveCollection(collection)
	})
}
