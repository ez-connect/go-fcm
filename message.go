package fcm

type FCMOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
}

type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Image string `json:"image,omitempty"`
}

type Message struct {
	Name         string            `json:"name,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
	Notification *Notification     `json:"notification,omitempty"`
	Android      *AndroidConfig    `json:"android,omitempty"`
	Webpush      *WebpushConfig    `json:"webpush,omitempty"`
	APNS         *APNSConfig       `json:"apns,omitempty"`
	FCMOptions   *FCMOptions       `json:"fcm_options,omitempty"`
	Token        string            `json:"token,omitempty"`
	Topic        string            `json:"-"`
	Condition    string            `json:"condition,omitempty"`
}

// The request body contains data.
// https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages/send#request-body
type MessageBody struct {
	ValidateOnly bool    `json:"validate_only,omitempty"`
	Message      Message `json:"message,omitempty"`
}
