package arcaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const arcaptchaApi = "https://arcaptcha.ir/2/siteverify"

type Website struct {
	SiteKey   string
	SecretKey string
}

func NewWebsite(siteKey, secretKey string) *Website {
	return &Website{
		SiteKey:   siteKey,
		SecretKey: secretKey,
	}
}

func (w *Website) ValidateCaptcha(challengeID string) (success bool, err error) {

	data := &verifyCaptchaRequest{
		SiteKey:     w.SiteKey,
		SecretKey:   w.SecretKey,
		ChallengeID: challengeID,
	}
	bin, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		http.MethodPost,
		arcaptchaApi,
		bytes.NewBuffer(bin),
	)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(string(bin))))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	bin, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var response VerifyCaptchaResponse
	if err = json.Unmarshal(bin, &response); err != nil {
		return
	}

	success = response.Success
	return
}
