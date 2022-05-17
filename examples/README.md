# Arcaptcha Verification Example

this example is an HTTP server that serves a html form with a captcha and test the verification.

## [Guide](https://arcaptcha.ir/guide)
### Installation
```shell
go get -u github.com/arcaptcha/arcaptcha-go
cd $GOPATH/src/github.com/arcaptcha/arcaptcha-go
go build example.go
```
or
```shell
git clone https://github.com/arcaptcha/arcaptcha-go.git
cd arcaptcha-go/examples
go build example.go
```

### Run
```shell
./example <siteKey> <secretKey>
```

Now visit http://localhost:9876 on browser.

Note: you must add `localhost` to your website domains.