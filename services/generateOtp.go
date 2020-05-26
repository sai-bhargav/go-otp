package generateOtp

import (
	"fmt"
	"github.com/sai-bhargav/mfa-cli/config"
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func Otp() {

	var app string
	var path string
	
	flag.StringVar(&app, "app", "", "help message for flagname")
	flag.StringVar(&path, "path", "", "help message for flagname")

	set_config := flag.Bool("config", false , "set path to your secrets.json file")
	flag.Parse()

	if *set_config {

		d1 := []byte(path)

		err := ioutil.WriteFile("config/secrets", d1, 0644)
    if err != nil {
        panic(err)
		}
	}

	fileIO, err := os.OpenFile("config/secrets", os.O_RDWR, 0600)

		if err != nil {
			panic(err)
		}
		defer fileIO.Close()

		rawBytes, err := ioutil.ReadAll(fileIO)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(rawBytes), "\n")
		file_req := lines[0]
		fmt.Print("\n")

	config.Config(app, file_req)
}
