package config

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/sai-bhargav/mfa-cli/totp"
)

func Config(appName string, filePath string){

	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
			panic(err)
	}

	var apps map[string]string

	err = json.Unmarshal(jsonFile, &apps)
	if err != nil {
		panic(err)
	}

	otp := authenticator.AuthenticatorTotp(apps[appName])

	fmt.Println("Generated OTP")
	fmt.Println(otp)
}
