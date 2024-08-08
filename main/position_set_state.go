package main

import (
	"github.com/pocketbase/pocketbase/core"
	"mensadb/tools/spatial"
)

func PositionSetState(e *core.RecordCreateEvent) error {
	lat := e.Record.Get("lat").(float64)
	lon := e.Record.Get("lon").(float64)
	state := spatial.LoadState(lat, lon)
	e.Record.Set("state", state)
	app.Dao().Save(e.Record)
	return nil
}
