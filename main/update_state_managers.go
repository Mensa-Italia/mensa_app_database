package main

import (
	"mensadb/importers"
	"slices"
	"sync"
	"time"
)

var lockStateManagers sync.Mutex

func updateStateManagers() {
	successLock := lockStateManagers.TryLock()
	if !successLock { // not able to lock so is already running, abort this run
		return
	}
	defer lockStateManagers.Unlock()
	app.Logger().Info("Updating states managers permissions, this may take a while. Waiting 1 minute before starting for security reasons.")
	time.Sleep(1 * time.Minute)

	records, err := app.Dao().FindRecordsByFilter("users", "powers:length > -1", "-created", -1, 0)
	if err != nil {
		return
	}
	segretari := importers.RetrieveForwardedMail("segretari")
	for _, record := range records {

		print(record.GetString("email"))
		powers := record.GetStringSlice("powers")
		newPowers := []string{}
		for _, power := range powers {
			if power == "events" {
				continue
			}
			newPowers = append(newPowers, power)
		}
		if slices.Contains(segretari, record.GetString("email")) {
			newPowers = append(newPowers, "events")
		}
		record.Set("powers", newPowers)
		_ = app.Dao().Save(record)
	}
}
