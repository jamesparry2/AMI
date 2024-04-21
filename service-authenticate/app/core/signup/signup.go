package signup

type SignupRequest struct{}

type SignupResponse struct{}

func (c *Client) Signup(r *SignupRequest) *SignupResponse {
	return nil
}
