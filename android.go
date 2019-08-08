package fcm

import "time"

// AndroidConfig contains messaging options specific to the Android platform.
type AndroidConfig struct {
	CollapseKey           string               `json:"collapse_key,omitempty"`
	Priority              string               `json:"priority,omitempty"` // one of "normal" or "high"
	TTL                   *time.Duration       `json:"-"`
	RestrictedPackageName string               `json:"restricted_package_name,omitempty"`
	Data                  map[string]string    `json:"data,omitempty"` // if specified, overrides the Data field on Message type
	Notification          *AndroidNotification `json:"notification,omitempty"`
	FCMOptions            *AndroidFCMOptions   `json:"fcm_options,omitempty"`
}

// AndroidFCMOptions contains additional options for features provided by the FCM Android SDK.
type AndroidFCMOptions struct {
	AnalyticsLabel string `json:"analytics_label,omitempty"`
}

// AndroidNotification is a notification to send to Android devices.
type AndroidNotification struct {
	Title        string   `json:"title,omitempty"` // if specified, overrides the Title field of the Notification type
	Body         string   `json:"body,omitempty"`  // if specified, overrides the Body field of the Notification type
	Icon         string   `json:"icon,omitempty"`
	Color        string   `json:"color,omitempty"` // notification color in #RRGGBB format
	Sound        string   `json:"sound,omitempty"`
	Tag          string   `json:"tag,omitempty"`
	ClickAction  string   `json:"click_action,omitempty"`
	BodyLocKey   string   `json:"body_loc_key,omitempty"`
	BodyLocArgs  []string `json:"body_loc_args,omitempty"`
	TitleLocKey  string   `json:"title_loc_key,omitempty"`
	TitleLocArgs []string `json:"title_loc_args,omitempty"`
	ChannelID    string   `json:"channel_id,omitempty"`
}
