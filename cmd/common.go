package cmd

import (
	"errors"
	"log"

	playfab "github.com/dgkanatsios/playfabsdk-go/sdk"
	"github.com/dgkanatsios/playfabsdk-go/sdk/authentication"
	"github.com/dgkanatsios/playfabsdk-go/sdk/multiplayer"
	"github.com/spf13/viper"
)

const titleIDConfig = "titleID"
const entityTokenConfig = "token"

func getLoginSettings() (*playfab.Settings, string) {
	settings := playfab.NewSettingsWithDefaultOptions("titleID")
	loginData := &authentication.GetEntityTokenRequestModel{}

	res, err := authentication.GetEntityToken(settings, loginData, "", "", "entityToken")
	if err != nil {
		log.Fatal(err)
	}

	return settings, res.EntityToken
}

func login(titleID, entityToken string) (string, error) {
	settings := playfab.NewSettingsWithDefaultOptions(titleID)
	loginData := &authentication.GetEntityTokenRequestModel{}

	res, err := authentication.GetEntityToken(settings, loginData, "", "", entityToken)
	if err != nil {
		return "", err
	}

	if res.EntityToken == "" {
		return "", errors.New("Incorrect login details")
	}

	return res.EntityToken, nil
}

func getSettings() *playfab.Settings {
	titleID := viper.Get(titleIDConfig)

	if titleID == nil {
		log.Fatal("Cannot retrieve titleID, you need to login first")
	}
	if titleID.(string) == "" {
		log.Fatal("Empty titleID, maybe you need to login first")
	}

	return playfab.NewSettingsWithDefaultOptions(titleID.(string))
}

func getEntityToken() string {
	token := viper.Get(entityTokenConfig)

	if token == nil {
		log.Fatal("Cannot retrieve entityToken, you need to login first")
	}

	if token.(string) == "" {
		log.Fatal("Empty entityToken, maybe you need to login first")
	}
	return token.(string)
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
