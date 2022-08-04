package main

import (
	"fmt"
	"github.com/arcaptcha/arcaptcha-go"
	"html/template"
	"log"
	"net/http"
	"os"
)

var website *arcaptcha.Website

// main get two argument siteKey and secretKey and serve arcaptcha demo html page
func main() {
	// read siteKey and secretKey from args
	var siteKey, secretKey string
	if len(os.Args) != 3 {
		fmt.Println("missing args siteKey and secretKey")
		return
	} else {
		siteKey = os.Args[1]
		secretKey = os.Args[2]
	}

	// init website
	website = arcaptcha.NewWebsite(siteKey, secretKey)

	fmt.Println("serving on :9876")
	http.HandleFunc("/", handleDemo)
	if err := http.ListenAndServe(":9876", nil); err != nil {
		log.Fatal("failed to start server", err)
	}

}

func handleDemo(w http.ResponseWriter, r *http.Request) {
	var successMsg, errorMsg string
	// check form is submitted
	r.ParseForm()
	submitted := r.FormValue("submitted")
	if submitted != "" {
		// verify captcha
		response := r.FormValue("arcaptcha-response")
		result, err := website.Verify(response)
		if err != nil {
			log.Fatal(err)
		}
		if result.Success {
			successMsg = "captcha verified"
		} else {
			errorMsg = fmt.Sprintf("captcha not verified, error-codes: %v", result.ErrorCodes)
		}
	}

	// render html
	tmpl, err := template.ParseFiles("demo.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	data := struct {
		SiteKey    string
		SuccessMsg string
		ErrorMsg   string
	}{
		SiteKey:    website.SiteKey,
		SuccessMsg: successMsg,
		ErrorMsg:   errorMsg,
	}
	err = tmpl.ExecuteTemplate(w, "demo", &data)
}
