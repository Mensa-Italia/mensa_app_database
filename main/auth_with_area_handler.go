package main

import (
	"context"
	"crypto"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"mensadb/area32"
	"mensadb/tools/env"
	"slices"
	"strings"
)

func AuthWithAreaHandler(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	scraperApi := area32.NewAPI()
	areaUser, err := scraperApi.DoLoginAndRetrieveMain(email, password)
	if err != nil {
		return apis.NewBadRequestError("Invalid credentials", err)
	}

	byUser, err := app.Dao().FindRecordById("users", areaUser.Id)

	if byUser == nil || err != nil {
		// Create a new user
		collection, _ := app.Dao().FindCollectionByNameOrId("users")
		newUser := models.NewRecord(collection)
		newUser.SetId(areaUser.Id)
		newUser.SetEmail(email)
		newUser.SetUsername(app.Dao().SuggestUniqueAuthRecordUsername("users", strings.Split(email, "@")[0]))
		newUser.SetPassword(generatePassword(areaUser.Id))
		newUser.SetVerified(true)
		newUser.Set("name", areaUser.Fullname)
		newUser.Set("expire_membership", areaUser.ExpireDate)
		newUser.Set("is_membership_active", areaUser.IsMembershipActive)
		if areaUser.IsATestMaker {
			newUser.Set("powers", []string{"testmakers"})
		}

		if err := app.Dao().SaveRecord(newUser); err != nil {
			return apis.NewBadRequestError("Invalid credentials", err)
		}
		fileImage, err := filesystem.NewFileFromUrl(context.Background(), areaUser.ImageUrl)
		if err == nil {
			form := forms.NewRecordUpsert(app, newUser)
			form.AddFiles("avatar", fileImage)
			_ = form.Submit()
		}
		return apis.RecordAuthResponse(app, c, newUser, nil)
	} else {
		byUser.SetEmail(email)
		byUser.SetVerified(true)
		byUser.Set("name", areaUser.Fullname)
		byUser.Set("expire_membership", areaUser.ExpireDate)
		byUser.Set("is_membership_active", areaUser.IsMembershipActive)

		powers := byUser.GetStringSlice("powers")
		if areaUser.IsATestMaker && !slices.Contains(powers, "testmakers") {
			powers = append(powers, "testmakers")
			byUser.Set("powers", powers)
		} else if !areaUser.IsATestMaker && slices.Contains(powers, "testmakers") {
			powers = removeFromSlice(powers, "testmakers")
			byUser.Set("powers", powers)
		}

		if err := app.Dao().Save(byUser); err != nil {
			return apis.NewBadRequestError("Invalid credentials", err)
		}

		byUser, _ = app.Dao().FindRecordById("users", areaUser.Id)

		if err != nil || !byUser.ValidatePassword(generatePassword(areaUser.Id)) {
			return apis.NewBadRequestError("Invalid credentials", err)
		}
		return apis.RecordAuthResponse(app, c, byUser, nil)
	}

}

func generatePassword(id string) string {
	pass := crypto.SHA256.New()
	pass.Write([]byte(id + uuid.NewMD5(uuid.MustParse(env.GetPasswordUUID()), []byte(id)).String() + env.GetPasswordSalt()))
	return hex.EncodeToString(pass.Sum(nil))
}

func removeFromSlice(slice []string, element string) []string {
	i := slices.Index(slice, element)
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
