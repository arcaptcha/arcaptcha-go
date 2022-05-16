package arcaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// arcaptchaApi arcaptcha verify API for captcha V2
const arcaptchaApi = "https://arcaptcha.ir/2/siteverify"

type Website struct {
	SiteKey   string
	SecretKey string
}

type verifyReq struct {
	SiteKey     string `json:"site_key"`
	SecretKey   string `json:"secret_key"`
	ChallengeID string `json:"challenge_id"`
}

// VerifyResp represents verify API response
// error codes are available in https://docs.arcaptcha.ir/en/API/Verify
type VerifyResp struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

// NewWebsite creates a new Website
func NewWebsite(siteKey, secretKey string) *Website {
	return &Website{
		SiteKey:   siteKey,
		SecretKey: secretKey,
	}
}

// Verify calls arcaptcha verify API and returns result.
//
// if an error occurs while sending or receiving the request, returns error.
// server side errors are available in VerifyResp.ErrorCodes.
func (w *Website) Verify(token string) (VerifyResp, error) {
	data := &verifyReq{
		SiteKey:     w.SiteKey,
		SecretKey:   w.SecretKey,
		ChallengeID: token,
	}
	var resp VerifyResp
	err := sendRequest(http.MethodPost, arcaptchaApi, data, &resp)
	return resp, err
}

// sendRequest sends http request to 'url' and fill 'resp' by response body
func sendRequest(method, url string, data interface{}, resp interface{}) error {
	bin, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bin))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
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
	if err = json.Unmarshal(bin, resp); err != nil {
		return err
	}
	return nil
}
