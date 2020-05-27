package authenticator

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
)

func Run() {

	var app string
	var token string
	
	flag.StringVar(&app, "app", "", "name of the app")
	flag.StringVar(&token, "token", "", "secret token for mfa authentication")

	set_config := flag.Bool("config", false , "to add new apps for mfa authentication")
	flag.Parse()

	resolveCommands(app, token, set_config)
}

func resolveCommands(app string, token string, set_config *bool, ) {

	if *set_config{
		configureAppTokens(app, token)
	} else {
		generateOtp(app)
	}
}

func configureAppTokens(app string, token string){

	appSecrets := make(map[string]string)

	jsonFile, err := ioutil.ReadFile("config/secrets.json")
	if err != nil {
			panic(err)
	}

	err = json.Unmarshal(jsonFile, &appSecrets)
	if err != nil {
		panic(err)
	}

	appSecrets[app] = token

	empData, err_marshal := json.Marshal(appSecrets)   
	if err_marshal != nil {
			panic(err_marshal)
			return
	}

	appSecretsBytes := []byte(string(empData))

	err_write := ioutil.WriteFile("config/secrets.json", appSecretsBytes, 0644)
	if err_write != nil {
			panic(err_write)
	}

	msg := "Configuration for " + app + " added successfully"
	fmt.Println(msg)
}
