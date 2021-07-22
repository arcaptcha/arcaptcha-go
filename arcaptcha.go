package arcaptcha

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const arcaptchaApi = "http://localhost:80/arcaptcha/api/verify"

type Website struct {
	SiteKey   string `json:"site_key"`
	SecretKey string `json:"secret_key"`
}

func NewWebsite(siteKey, secretKey string) *Website {
	return &Website{
		SiteKey:   siteKey,
		SecretKey: secretKey,
	}
}
func (w *Website) ValidateCaptcha(challengeID string) error {

	data := &verifyCaptchaRequest{
		SiteKey:     w.SiteKey,
		SecretKey:   w.SecretKey,
		ChallengeID: challengeID,
	}
	bin, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		arcaptchaApi,
		bytes.NewBuffer(bin),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(string(bin))))

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
	var errRes Error
	if err = json.Unmarshal(bin, &errRes); err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errRes
	}
	return nil

}
