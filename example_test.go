package fcm

import "fmt"

func Example() {
	// Create a client
	client := NewClient("your-sever-key")

	// New message
	message := Message{
		To: "device-token-or-topic", // /topics/topic-name
		Notification: Notification{
			Title: "Title",
			Body:  "Body",
		},
	}

	// Send the message
	resp, err := client.Send(message)
	if err != nil {
		fmt.Println(resp)
	}

	// Subscribe
	err = client.SubscribeToTopic([]string{"device-token"}, "/topics/test-topic-name")

	// Unsubscribe
	err = client.UnsubscribeFromTopic([]string{"device-token"}, "/topics/test-topic-name")

	// Get device info
	info, err := client.GetDeviceInfo("device-token")
	fmt.Println(info)
}
