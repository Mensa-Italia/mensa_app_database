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

		collection, err := dao.FindCollectionByNameOrId("wawd580chzpiyam")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("user = @request.auth.id && @request.data.qt <= event_schedule.max_external_guests")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("wawd580chzpiyam")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("user = @request.auth.id && @request.data.qt <= event_schedule.Max_external_guests")

		return dao.SaveCollection(collection)
	})
}
