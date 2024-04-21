package inmemory

import "github.com/jamesparry2/aim/service-authenticate/app/repository"

func (c *Client) Insert(request *repository.Request) *repository.Response {
	c.values[request.PrimaryID] = request.Value

	return &repository.Response{
		Item: c.values[request.PrimaryID],
	}
}
