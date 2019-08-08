package fcm

import "errors"

// POST https://fcm.googleapis.com/v1/{parent=projects/*}/messages:send
const API_MESSAGE = "https://fcm.googleapis.com/v1/projects/%s/messages:send"

type Client struct {
	ServerKey string
}

func (c *Client) New(serverKey string) error {
	if serverKey != "" {
		c.ServerKey = serverKey
		return nil
	}

	return errors.New("Missing `serverKey`")
}

// Send sends a Message to Firebase Cloud Messaging.
//
// The Message must specify exactly one of Token, Topic and Condition fields.
// FCM will customize the message for each target platform based on the arguments
// specified in the Message.
func (c *Client) Send(message Message) (*Response, error) {
	return nil, nil
}

// SubscribeToTopic subscribes a list of registration tokens to a topic.
// The tokens list must not be empty, and have at most 1000 tokens.
func (c *Client) SubscribeToTopic(tokens []string, topic string) (*TopicManagementResponse, error) {
	return nil, nil
}

// UnsubscribeFromTopic unsubscribes a list of registration tokens from a topic.
// The tokens list must not be empty, and have at most 1000 tokens.
func (c *Client) UnsubscribeFromTopic(tokens []string, topic string) (*TopicManagementResponse, error) {
	return nil, nil
}
