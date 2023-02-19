<p align="center"><img src="https://arcaptcha.ir/_nuxt/023053fecdcdf20e40bdc993c754d487.svg" height="100px"></p>

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/arcaptcha/arcaptcha-go)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/arcaptcha/arcaptcha-go/master/LICENSE)
[![Twitter](https://img.shields.io/badge/twitter-@arcaptcha-55acee.svg?style=flat-square)](https://twitter.com/arcaptcha)


Arcaptcha
=====================
Arcaptcha library implementation in Golang to verify captcha.

## [Guide](https://arcaptcha.ir/guide)

### Installation

```
go get -u github.com/arcaptcha/arcaptcha-go
```

### Usage

Register on [Arcaptcha](https://arcaptcha.ir/), create website and get your own SiteKey and SecretKey

```go
website := arcaptcha.NewWebsite("YOUR_SITE_KEY", "YOUR_SECRET_KEY")
//'arcaptcha-response' is created for each captcha
//After you put captcha widget in your website, you can get 'arcaptcha-response' from form
result, err := website.Verify("arcaptcha-response")
if err != nil {
// error in sending or receiving API request
// handle error
}
if !result.Success {
// captcha not verified
// can see result.ErrorCodes to find what's wrong
// throw specific error
}
// it's ok
```

#### Use in middleware:

```go
package main

import (
	"net/http"

	"github.com/arcaptcha/arcaptcha-go"
)

var website *arcaptcha.Website

func main() {
	website = arcaptcha.NewWebsite("YOUR_SITE_KEY", "YOUR_SECRET_KEY")

	myHandler := http.HandlerFunc(handler)
	http.Handle("/", verifyCaptcha(myHandler))
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// handle request
}

func verifyCaptcha(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		response := r.FormValue("arcaptcha-response")
		result, err := website.Verify(response)
		if err != nil {
			// error in sending or receiving API request
			// handle error
		}
		if !result.Success {
			// captcha not verified
			// can see result.ErrorCodes to find what's wrong
			// throw specific error
			return
		}
		// it's ok
		next.ServeHTTP(w, r)
	})
}
```

you can set `timeout` for verify request:

```go
website := arcaptcha.NewWebsite("YOUR_SITE_KEY", "YOUR_SECRET_KEY")
website.SetTimeout(1*time.Second)
result, err := website.Verify("arcaptcha-response") // returns "context deadline" error if it takes more than 1 second
```