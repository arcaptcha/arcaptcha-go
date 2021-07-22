package arcaptcha

type verifyCaptchaRequest struct {
	SiteKey     string `json:"site_key"`
	SecretKey   string `json:"secret_key"`
	ChallengeID string `json:"challenge_id"`
}

type Error struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}
