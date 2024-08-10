package main

import (
	ics "github.com/arran4/golang-ical"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

type IcalEvents struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Created        types.DateTime `json:"created"`
	Updated        types.DateTime `json:"updated"`
	WhenStart      types.DateTime `json:"when_start"`
	WhenEnd        types.DateTime `json:"when_end"`
	Location       string         `json:"location"`
	Lat            string         `json:"lat"`
	Lon            string         `json:"lon"`
	State          string         `json:"state"`
	Owner          string         `json:"owner"`
	OrganizerEmail string         `json:"organizer_email"`
	InfoLink       string         `json:"info_link"`
}

func RetrieveICAL(c echo.Context) error {
	hashCode := c.PathParam("hash")

	resultCalendarStates, _ := app.Dao().FindRecordsByExpr("calendar_link", dbx.NewExp("hash = {:user}", dbx.Params{
		"user": hashCode,
	}))

	if len(resultCalendarStates) == 0 {
		return c.String(404, "Calendar not found")
	}

	var calendarStates = []interface{}{}

	for _, data := range resultCalendarStates[0].GetStringSlice("state") {
		calendarStates = append(calendarStates, data)
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	query := app.Dao().DB().Select(
		"events.id as id",
		"events.name as name",
		"events.description as description",
		"events.created as created",
		"events.updated as updated",
		"events.when_start as when_start",
		"events.when_end as when_end",
		"positions.name as location",
		"events.owner as owner",
		"events.info_link as info_link",
		"positions.name as location",
		"positions.state as state",
		"positions.lat as lat",
		"positions.lon as lon",
		"users.email as organizer_email",
	).From("events").InnerJoin("positions", dbx.NewExp("events.position = positions.id")).InnerJoin("users", dbx.NewExp("events.owner = users.id")).Where(
		dbx.In("positions.state", calendarStates...))

	var records []IcalEvents
	if err := query.All(&records); err != nil {
		return err
	}

	for _, record := range records {
		event := cal.AddEvent(record.Id)
		event.SetCreatedTime(record.Created.Time())
		event.SetDtStampTime(record.Created.Time())
		event.SetModifiedAt(record.Updated.Time())
		event.SetStartAt(record.WhenStart.Time())
		event.SetEndAt(record.WhenEnd.Time())
		event.SetDescription(record.Description)
		event.SetSummary(record.Name)
		event.SetLocation(record.Location)
		event.SetGeo(record.Lat, record.Lon)
		event.SetDescription(record.Description)
		if record.InfoLink != "" {
			event.SetURL(record.InfoLink)
		}
		event.SetOrganizer(record.OrganizerEmail)

		c.Response().Header().Set("Content-Type", "text/calendar")
	}
	return c.String(200, cal.Serialize())

}
