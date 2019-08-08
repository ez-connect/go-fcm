package fcm

type APNSConfig struct {
	Headers    map[string]string `json:"headers,omitempty"`
	Payload    *APNSPayload      `json:"payload,omitempty"`
	FCMOptions *APNSFCMOptions   `json:"fcm_options,omitempty"`
}

type APNSPayload struct {
	Aps        *Aps                   `json:"aps,omitempty"`
	CustomData map[string]interface{} `json:"-"`
}

type APNSFCMOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
}

type Aps struct {
	AlertString string `json:"-"`
	// Alert            *ApsAlert              `json:"-"`
	Badge *int   `json:"badge,omitempty"`
	Sound string `json:"-"`
	// CriticalSound    *CriticalSound         `json:"-"`
	ContentAvailable bool                   `json:"-"`
	MutableContent   bool                   `json:"-"`
	Category         string                 `json:"category,omitempty"`
	ThreadID         string                 `json:"thread-id,omitempty"`
	CustomData       map[string]interface{} `json:"-"`
}
