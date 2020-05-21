package main

import (
	"fmt"
	"github.com/sai-bhargav/mfa-cli/totp"
)

func main() {

	sample_access_code := "remxew64ockyb4ed7ocx2vzotpogc34q"

	ans := authenticator.AuthenticatorTotp(sample_access_code)

	fmt.Print("Generated OTP \n")
	fmt.Print(ans)

}

