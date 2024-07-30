package main

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/tidwall/gjson"
	"sync"
	"time"
)

var lock sync.Mutex

func updateAddonsData() {
	successLock := lock.TryLock()
	if !successLock { // not able to lock so is already running, abort this run
		return
	}
	defer lock.Unlock()
	time.Sleep(1 * time.Minute)
	query := app.Dao().RecordQuery("addons")
	records := []*models.Record{}
	if err := query.All(&records); err != nil {
		return
	}

	for _, record := range records {
		urlToCheck := record.Get("url").(string) + "/mensadata.json"
		if urlToCheck == "" {
			setInvalid(record)
			continue
		}
		get, err := resty.New().R().Get(urlToCheck)
		if err != nil {
			setInvalid(record)
			return
		}
		if get.StatusCode() != 200 {
			setInvalid(record)
			return
		}
		if get.Body() == nil {
			setInvalid(record)
			return
		}
		dataToUse := gjson.ParseBytes(get.Body())
		if dataToUse.Get("id").String() != record.GetId() {
			setInvalid(record)
			return
		}
		record.Set("name", dataToUse.Get("name").String())
		record.Set("description", dataToUse.Get("description").String())
		record.Set("version", dataToUse.Get("version").String())

		err = app.Dao().Save(record)
		if err != nil {
			setInvalid(record)
			return
		}

		fileImage, err := filesystem.NewFileFromUrl(context.Background(), dataToUse.Get("icon").String())
		if err == nil {
			form := forms.NewRecordUpsert(app, record)
			form.AddFiles("icon", fileImage)
			_ = form.Submit()
		}

		record2, err := app.Dao().FindRecordById("addons", record.GetId())
		if err != nil {
			setInvalid(record)
			return
		}

		if record2.GetString("name") != "" && record2.GetString("description") != "" && record2.GetString("version") != "" && record2.GetString("icon") != "" {
			record2.Set("is_ready", true)
			err = app.Dao().Save(record2)
			if err != nil {
				return
			}
		} else {
			record2.Set("is_ready", false)
			err = app.Dao().Save(record2)
			if err != nil {
				return
			}
		}
	}

}

func setInvalid(record *models.Record) {
	record.Set("is_ready", false)
	err := app.Dao().Save(record)
	if err != nil {
		return
	}
}
