package arcaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// arcaptchaApi arcaptcha verify API for captcha V2
const (
	arcaptchaApi   = "https://arcaptcha.co/2/siteverify"
	defaultTimeout = 5 * time.Second
)

type Website struct {
	SiteKey   string
	SecretKey string

	client    *http.Client
	verifyUrl string
}

type verifyReq struct {
	SiteKey   string `json:"sitekey"`
	SecretKey string `json:"secret"`
	Response  string `json:"response"`
	RemoteIp  string `json:"remoteip"`
}

// VerifyResp represents verify API response
// error codes are available in https://docs.arcaptcha.ir/en/API/Verify
type VerifyResp struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts,omitempty"`
	Hostname    string   `json:"hostname,omitempty"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
}

// NewWebsite creates a new Website
func NewWebsite(siteKey, secretKey string) *Website {
	return &Website{
		SiteKey:   siteKey,
		SecretKey: secretKey,
		client:    &http.Client{Timeout: defaultTimeout},
		verifyUrl: arcaptchaApi,
	}
}

func (w *Website) SetVerifyUrl(url string) {
	w.verifyUrl = url
}

func (w *Website) SetTimeout(timeout time.Duration) {
	w.client.Timeout = timeout
}

// Verify calls arcaptcha verify API and returns result.
//
// if an error occurs while sending or receiving the request, returns error.
// server side errors are available in VerifyResp.ErrorCodes.
func (w *Website) Verify(response string) (VerifyResp, error) {
	data := &verifyReq{
		SiteKey:   w.SiteKey,
		SecretKey: w.SecretKey,
		Response:  response,
	}
	var resp VerifyResp
	err := w.sendRequest(http.MethodPost, w.verifyUrl, data, &resp)
	return resp, err
}

// sendRequest sends http request to 'url' and fill 'resp' by response body
func (w *Website) sendRequest(method, url string, data, resp interface{}) error {
	bin, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bin))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := w.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	bin, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%v: %v", res.Status, string(bin))
	}
	if err = json.Unmarshal(bin, resp); err != nil {
		return err
	}
	return nil
}
