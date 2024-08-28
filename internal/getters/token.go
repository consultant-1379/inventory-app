package getters

import (
	"os"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/helpers"
)

func GetHydraToken(tokenPath string) string {

	token, err := os.ReadFile(tokenPath)
	helpers.CheckError(err)
	return string(token)
}
