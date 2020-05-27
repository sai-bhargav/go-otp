package generateOtp

import (
	"github.com/sai-bhargav/mfa-cli/config"
	"flag"
	"io/ioutil"
	"encoding/json"
)

func Otp() {

	var app string
	var secret string
	
	flag.StringVar(&app, "app", "", "help message for flagname")
	flag.StringVar(&secret, "secret", "", "help message for flagname")

	set_config := flag.Bool("config", false , "set path to your secrets.json file")
	flag.Parse()

	if *set_config {

		tes := make(map[string]string) 

		jsonFile, errz := ioutil.ReadFile("config/secrets.json")
		if errz != nil {
				panic(errz)
		}

		errz = json.Unmarshal(jsonFile, &tes)
		if errz != nil {
			panic(errz)
		}

		tes[app] = secret

		empData, errp := json.Marshal(tes)   
    if errp != nil {
        panic(errp)
        return
    }

		d1 := []byte(string(empData))

		err := ioutil.WriteFile("config/secrets.json", d1, 0644)
    if err != nil {
        panic(err)
		}
	}

	config.Config(app, "config/secrets.json")
}
