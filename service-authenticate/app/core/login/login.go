package login

type LoginRequest struct{}

type LoginResponse struct{}

func (c *Client) Login(r *LoginRequest) *LoginResponse {
	c.repo.Get(nil)
	return nil
}
