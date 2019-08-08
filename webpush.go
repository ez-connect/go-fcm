package fcm

type WebpushConfig struct {
	Headers      map[string]string `json:"headers,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	Notification Notification      `json:"notification,omitempty"`
	FCMOptions   WebpushFcmOptions `json:"fcm_options,omitempty"`
}

type WebpushFcmOptions struct {
	Link string `json:"link,omitempty"`
}
