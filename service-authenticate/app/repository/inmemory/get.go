package inmemory

import "github.com/jamesparry2/aim/service-authenticate/app/repository"

func (c *Client) Get(r *repository.Request) *repository.Response {
	return &repository.Response{
		Item: c.values[r.PrimaryID],
	}
}
