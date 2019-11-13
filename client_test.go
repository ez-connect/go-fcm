package fcm

import "testing"

// Settings > General > Web API Key
const API_KEY = "API Key"

// Settings > Cloud Messaging > Server key
const SERVER_KEY = "Server Key"

func TestGetDeviceInfo(t *testing.T) {
	client := NewClient(API_KEY, SERVER_KEY)
	_, err := client.GetDeviceInfo("device-token")
	if err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	client := NewClient(API_KEY, SERVER_KEY)
	_, err := client.GetUser("eyJhbGciOiJSUzI1NiIsImtpZCI6IjI1MDgxMWNkYzYwOWQ5MGY5ODE1MTE5MWIyYmM5YmQwY2ViOWMwMDQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vZm1zdHlsZS1hcHAiLCJhdWQiOiJmbXN0eWxlLWFwcCIsImF1dGhfdGltZSI6MTU3MzYyNjM1MywidXNlcl9pZCI6IjNjZnc5RmR2eUtjUjNFSG00VnVNUEdFSlRMajEiLCJzdWIiOiIzY2Z3OUZkdnlLY1IzRUhtNFZ1TVBHRUpUTGoxIiwiaWF0IjoxNTczNjI2MzUzLCJleHAiOjE1NzM2Mjk5NTMsInBob25lX251bWJlciI6Iis4NDkwNjUxNjU3OCIsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsicGhvbmUiOlsiKzg0OTA2NTE2NTc4Il19LCJzaWduX2luX3Byb3ZpZGVyIjoicGhvbmUifX0.I2Jiu8_Iato2WQLa0F5AcNdOgDnHBUonwJkhC2oIvR97JQGt2uougOIIODi6ur3h4VHpUiYvfaBXscmxui9H6ayCa-Hoqv_syXQc948uO9Ktk1AQpcPm5IfkHWKrRAlOZMYWXQw7pOqQHZpkzqqJrNhLvsyzA243fHeBQ3zMy6xs-ls0drLTy0ZTYZv7vVzfCecAQbAV3I9L8yquQqk0uRBAPO1gnE5Z1nmwiDc4c801B9_6f2v_0O_6S8FNCNiBpNBcsxNRjkkCm5Xr99lyDytmUahyl45TGrMJNOYRw0EMDdnvIUmSaUdyaOspNyJiPqJn8R7sciLXHfELQFBFeQ")
	if err != nil {
		t.Error(err)
	}
}

func TestSubscribeToTopic(t *testing.T) {
	client := NewClient(API_KEY, SERVER_KEY)
	err := client.SubscribeToTopic([]string{"device-token"}, "/topics/test-topic-name")
	if err != nil {
		t.Error(err)
	}
}

func TestSendMessage(t *testing.T) {
	messages := []Message{
		{
			To:     "device-token",
			DryRun: true,
			Notification: Notification{
				Title: "Test a device",
				Body:  "Test body",
			},
		},
		{
			To:     "wrong-value",
			DryRun: true,
			Notification: Notification{
				Title: "Test a device",
				Body:  "Test body",
			},
		},
		{
			To:     "/topics/a-topic",
			DryRun: true,
			Notification: Notification{
				Title: "Test",
				Body:  "Test body",
			},
		},
	}

	client := NewClient(API_KEY, SERVER_KEY)
	for _, v := range messages {
		if _, err := client.Send(v); err != nil {
			t.Error(err)
		}
	}
}

func TestUnsubscribeToTopic(t *testing.T) {
	client := NewClient(API_KEY, SERVER_KEY)
	err := client.UnsubscribeFromTopic([]string{"device-token"}, "/topics/test-topic")
	if err != nil {
		t.Error(err)
	}
}
