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

		collection, err := dao.FindCollectionByNameOrId("cl4zw57ht0zo7xk")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("(@request.auth.powers:each ?= \"deals\" || @request.auth.id = deal.owner) || (@request.auth.powers:each ?= \"super\")")

		collection.ViewRule = types.Pointer("(@request.auth.powers:each ?= \"deals\" || @request.auth.id = deal.owner) || (@request.auth.powers:each ?= \"super\")")

		collection.CreateRule = types.Pointer("(@request.auth.powers:each ?= \"deals\" || @request.auth.id = deal.owner) || (@request.auth.powers:each ?= \"super\")")

		collection.UpdateRule = types.Pointer("(@request.auth.powers:each ?= \"deals\" || @request.auth.id = deal.owner) || (@request.auth.powers:each ?= \"super\")")

		collection.DeleteRule = types.Pointer("(@request.auth.powers:each ?= \"deals\" || @request.auth.id = deal.owner) || (@request.auth.powers:each ?= \"super\")")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("cl4zw57ht0zo7xk")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = nil

		collection.CreateRule = nil

		collection.UpdateRule = nil

		collection.DeleteRule = nil

		return dao.SaveCollection(collection)
	})
}
