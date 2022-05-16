package arcaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

type VerifyResp struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

func NewWebsite(siteKey, secretKey string) *Website {
	return &Website{
		SiteKey:   siteKey,
		SecretKey: secretKey,
	}
}

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
