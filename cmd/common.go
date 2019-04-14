package cmd

import (
	"log"

	playfab "github.com/dgkanatsios/playfabsdk-go/sdk"
	"github.com/dgkanatsios/playfabsdk-go/sdk/authentication"
	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"
)

func getLoginSettings() (*playfab.Settings, string) {
	settings := playfab.NewSettingsWithDefaultOptions("titleID")
	loginData := &authentication.GetEntityTokenRequestModel{}

	res, err := authentication.GetEntityToken(settings, loginData, "", "", "entityToken")
	if err != nil {
		log.Fatal(err)
	}

	return settings, res.EntityToken
}

func lala() error {
	settings, entityToken := getLoginSettings()
	listMultiplayerServerDetails := &multiplayer.ListMultiplayerServersRequestModel{}
	_, err := multiplayer.ListMultiplayerServers(settings, listMultiplayerServerDetails, entityToken)
	if err != nil {
		return err
	}
	return nil
}
