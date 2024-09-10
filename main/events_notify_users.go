package main

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"google.golang.org/api/option"
	"mensadb/tools/env"
)

func EventsNotifyUsers(e *core.RecordCreateEvent) error {
	if e.Record.Get("is_national") != true {
		return nil
	}
	decodedKey, err := getDecodedFireBaseKey()
	if err != nil {
		return err
	}

	opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}
	appFirebase, err := firebase.NewApp(context.Background(), nil, opts...)
	if err != nil {
		return err
	}
	fcmClient, err := appFirebase.Messaging(context.Background())
	if err != nil {
		return err
	}

	expr, err := app.Dao().FindRecordsByExpr("users_devices", dbx.NewExp(`firebase_id != {:id}`, dbx.Params{
		"id": "NOTOKEN",
	}))
	if err != nil {
		return err
	}
	var tokens []string
	for _, record := range expr {
		tokens = append(tokens, record.GetString("firebase_id"))
	}

	// each 400 tokens will be sent in a separate request
	for i := 0; i < len(tokens); i += 400 {
		maxcount := i + 400
		if maxcount > len(tokens) {
			maxcount = len(tokens)
		}
		go func(newTokenList []string) {
			fcmClient.SendEachForMulticast(context.Background(), &messaging.MulticastMessage{
				Notification: &messaging.Notification{
					Title: "New National Event!",
					Body:  "Check out the new national event and get ready to participate!",
				},
				Tokens: newTokenList,
			})
		}(tokens[i:maxcount])
	}
	return nil
}

func getDecodedFireBaseKey() ([]byte, error) {

	fireBaseAuthKey := env.GetFireBaseAuthKey()

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}
