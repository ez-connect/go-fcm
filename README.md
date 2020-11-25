### DEPRECATED - no longer actively maintained

---

# Firebase Cloud Messaging

Golang client library for Firebase Cloud Messaging.
It uses Firebase Cloud Messaging HTTP protocol: https://firebase.google.com/docs/cloud-messaging/http-server-ref.

## Getting Started

```sh
$ go get github.com/ez-connect/go-fcm
```

## Send message

```go
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
```
## Subscribe

```go
// Subscribe
err := client.SubscribeToTopic([]string{"device-token"}, "/topics/test-topic-name")

// Unsubscribe
err = client.UnsubscribeFromTopic([]string{"device-token"}, "/topics/test-topic-name")
```

## Device Info

```go
// Get device info
info, err := client.GetDeviceInfo("device-token")
if err != nil {
    fmt.Println(info)
}
```
