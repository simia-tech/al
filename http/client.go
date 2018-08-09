package http

var DefaultClient = NewClient()

type Client struct{}

func NewClient() *Client {
	return &Client{}
}
