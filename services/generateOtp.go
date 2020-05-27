package authenticator

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"github.com/sai-bhargav/mfa-cli/totp"
)

func generateOtp(app string) {
	
	token := fetchAppToken(app)
	otp := authenticator.AuthenticatorTotp(token)

	fmt.Println("Generated OTP")
	fmt.Println(otp)

}

func fetchAppToken(appName string) (string) {

	jsonFile, err := ioutil.ReadFile("config/secrets.json")
	if err != nil {
			panic(err)
	}

	var apps map[string]string

	err = json.Unmarshal(jsonFile, &apps)
	if err != nil {
		panic(err)
	}

	if apps[appName] == "" {
		err := "App " + appName + " is not configured.\nPlease ensure to configure the app before generating otp."
		fmt.Println(err)
		os.Exit(0)
	}

	return apps[appName]
}
