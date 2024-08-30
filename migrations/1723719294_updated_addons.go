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

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.powers ?= \"addons\" || @request.auth.powers ?= \"super\"")

		collection.UpdateRule = types.Pointer("@request.auth.powers ?= \"addons\" || @request.auth.powers ?= \"super\"")

		collection.DeleteRule = types.Pointer("@request.auth.powers ?= \"addons\" || @request.auth.powers ?= \"super\"")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.powers ?= \"addons\"")

		collection.UpdateRule = types.Pointer("@request.auth.powers ?= \"addons\"")

		collection.DeleteRule = types.Pointer("@request.auth.powers ?= \"addons\"")

		return dao.SaveCollection(collection)
	})
}
