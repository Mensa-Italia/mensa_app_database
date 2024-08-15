package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nq7uqdrz46xryck")
		if err != nil {
			return err
		}

		collection.CreateRule = nil

		collection.UpdateRule = types.Pointer("@request.auth.id = user && @request.data.hash:isset = false")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("nq7uqdrz46xryck")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id = user")

		collection.UpdateRule = types.Pointer("")

		return dao.SaveCollection(collection)
	})
}
