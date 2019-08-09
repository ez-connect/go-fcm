package fcm

import "testing"

func TestGetDeviceInfo(t *testing.T) {
	client := NewClient("sever-key")
	_, err := client.GetDeviceInfo("device-token")
	if err != nil {
		t.Error(err)
	}
}

func TestSubscribeToTopic(t *testing.T) {
	client := NewClient("sever-key")
	err := client.SubscribeToTopic([]string{"device-token"}, "test-topic-name")
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

	client := NewClient("server-key")
	for _, v := range messages {
		if _, err := client.Send(v); err != nil {
			t.Error(err)
		}
	}
}

func TestUnsubscribeToTopic(t *testing.T) {
	client := NewClient("server-key")
	err := client.UnsubscribeFromTopic([]string{"device-token"}, "test-topic")
	if err != nil {
		t.Error(err)
	}
}
