package inmemory

type Client struct {
	values map[string]any
}

func New() *Client {
	return &Client{
		values: map[string]any{},
	}
}
