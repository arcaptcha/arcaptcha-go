## Arcaptcha Verification Example

this example is an HTTP server that serves a html form with a captcha and test the verification.

### Install
```shell
go get -u github.com/arcaptcha/arcaptcha-go
cd $GOPATH/src/github.com/arcaptcha/arcaptcha-go
go build example.go
```

### Run
```shell
./example <siteKey> <secretKey>
```

Now visit http://localhost:9876 on browser.

Note: you must add `localhost` to your website domains.