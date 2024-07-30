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

		collection.ListRule = types.Pointer("is_ready = true")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("30eqczpxcx155p1")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		return dao.SaveCollection(collection)
	})
}
