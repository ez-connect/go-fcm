// Golang client library for Firebase Cloud Messaging.
// It uses Firebase Cloud Messaging HTTP protocol: https://firebase.google.com/docs/cloud-messaging/http-server-ref.
package fcm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func NewClient(serverKey string) *Client {
	client := &Client{
		ServerKey: serverKey,
	}

	return client
}

type Client struct {
	// Server key
	// Firebase console > Project Settings > Cloud Messaging > Server key
	ServerKey string
}

// Get information about app instances
//
// https://developers.google.com/instance-id/reference/server#create_a_relation_mapping_for_an_app_instance
func (c *Client) GetDeviceInfo(token string) (*DeviceInfoResponse, error) {
	const API_APP_INSTANCE = "https://iid.googleapis.com/iid/info/%s?details=true"
	endpoint := fmt.Sprintf(API_APP_INSTANCE, token)
	resp, err := c.request(http.MethodGet, endpoint, nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	result := DeviceInfoResponse{}
	// If you have an io.Reader, use Decoder. Otherwise use Unmarshal.
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Send sends a Message to Firebase Cloud Messaging.
//
// The Message must specify exactly one of Token, Topic and Condition fields.
// FCM will customize the message for each target platform based on the arguments
// specified in the Message.
// Use Legacy HTTP Server Protocol
//
// https://firebase.google.com/docs/cloud-messaging/http-server-ref
func (c *Client) Send(message Message) (*MessageResponse, error) {
	const API_SEND_MESSAGE = "https://fcm.googleapis.com/fcm/send"
	body, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, API_SEND_MESSAGE, body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	result := MessageResponse{}
	// If you have an io.Reader, use Decoder. Otherwise use Unmarshal.
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SubscribeToTopic subscribes a list of registration tokens to a topic.
// The tokens list must not be empty, and have at most 1000 tokens.
func (c *Client) SubscribeToTopic(tokens []string, topic string) error {
	const API_TOPIC_BATCH_ADD = "https://iid.googleapis.com/iid/v1:batchAdd"
	return c.subscribe(API_TOPIC_BATCH_ADD, tokens, topic)
}

// UnsubscribeFromTopic unsubscribes a list of registration tokens from a topic.
// The tokens list must not be empty, and have at most 1000 tokens.
func (c *Client) UnsubscribeFromTopic(tokens []string, topic string) error {
	const API_TOPIC_BATCH_REMOVE = "https://iid.googleapis.com/iid/v1:batchRemove"
	return c.subscribe(API_TOPIC_BATCH_REMOVE, tokens, topic)
}

///////////////////////////////////////////////////////////////////

func (c *Client) request(method string, endpoint string, body []byte) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
	var err error

	req, err = http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err == nil {
		// Add headers
		req.Header.Add("Authorization", fmt.Sprintf("key=%s", c.ServerKey))
		req.Header.Add("Content-Type", "application/json")

		// Execute
		client := &http.Client{}
		resp, err = client.Do(req)
	}

	return resp, err
}

func (c *Client) subscribe(endpoint string, tokens []string, topic string) error {
	data := TopicSubscribe{
		RegistrationTokens: tokens,
		// To:                 fmt.Sprintf("/topics/%s", topic),
		To: topic,
	}

	body, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPost, endpoint, body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}
