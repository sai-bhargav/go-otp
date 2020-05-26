package config

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/sai-bhargav/mfa-cli/totp"
)

type Apps struct {
	Apps []Credentials
}

type Credentials struct {
	App_name   string
	Secret_key   string
}

func Config(appName string, filePath string){

	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
			panic(err)
	}

	var apps Apps

	err = json.Unmarshal(jsonFile, &apps)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i := 0; i<len(apps.Apps); i++{
		if apps.Apps[i].App_name == appName {
			ans := authenticator.AuthenticatorTotp(apps.Apps[i].Secret_key)

			fmt.Print("Generated OTP \n")
			fmt.Print(ans)
		}
	}
}
