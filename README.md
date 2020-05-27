# go-otp
A CLI tool used for Multifactor Authentication by generating Time Based OTP

#### Getting Started

```sh
# Install the package
$ go get github.com/sai-bhargav/go-otp

# Go to the directory where the pkg is installed and build
$ cd $GOPATH/src/github.com/sai-bhargav/go-otp
$ go build

# You should see the new executable created 'go-otp' after the build
```

#### Usage

```sh
# Configure the app and its access token 
$ go-otp -config -app=appName -token=YOURACCESSTOKENFORAPP

# To generate the OTP
$ go-otp -app=appName
```

#### Additional Info
Alternatively copy the executable to `/usr/local/bin` to make it available globally 
```sh
cp go-otp /usr/local/bin
```
