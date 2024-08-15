package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/pocketbase/pocketbase/core"
)

func CalendarSetHash(e *core.RecordCreateEvent) error {
	e.Record.Set("hash", randomHash())
	app.Dao().Save(e.Record)
	return nil
}

func randomHash() string {
	h := sha256.New()
	h.Write([]byte(uuid.New().String() + "  " + uuid.New().String()))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
