package arcaptcha

type verifyCaptchaRequest struct {
	SiteKey     string `json:"site_key"`
	SecretKey   string `json:"secret_key"`
	ChallengeID string `json:"challenge_id"`
}

type VerifyCaptchaResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}
