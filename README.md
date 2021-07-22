<p align="center"><img src="https://arcaptcha.ir/logo.png" height="150px"></p>

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/arcaptcha/arcaptcha-go)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/arcaptcha/arcaptcha-go/master/LICENSE)
[![Twitter](https://img.shields.io/badge/twitter-@arcaptcha-55acee.svg?style=flat-square)](https://twitter.com/arcaptcha)


Arcaptcha
=====================
Arcaptcha library implementation in Golang for validate captcha.

## [Guide](https://arcaptcha.ir/guide)
### Installation

```
go get -u github.com/arcaptcha/arcaptcha-go
```

### Usage
Register on [Arcaptcha](https://arcaptcha.ir/) and get your own SiteKey and SecretKey
```go
website := arcaptcha.NewWebsite(SiteKey, SecretKey)
err := website.ValidateCaptcha(challengeID)
```