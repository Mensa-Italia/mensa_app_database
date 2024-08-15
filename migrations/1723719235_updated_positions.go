package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.powers:length > 0")

		collection.UpdateRule = types.Pointer("@request.auth.powers:length > 0")

		collection.DeleteRule = types.Pointer("@request.auth.powers:length > 0")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("48avh9pvrlw8w7n")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("(@request.auth.powers:each ?= \"events\") || (@request.auth.powers:each ?= \"super\")")

		collection.UpdateRule = types.Pointer("(@request.auth.powers:each ?= \"events\") || (@request.auth.powers:each ?= \"super\")")

		collection.DeleteRule = types.Pointer("(@request.auth.powers:each ?= \"events\") || (@request.auth.powers:each ?= \"super\")")

		return dao.SaveCollection(collection)
	})
}
