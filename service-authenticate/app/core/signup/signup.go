package signup

import (
	"errors"

	"github.com/jamesparry2/aim/service-authenticate/app/repository"
)

var (
	ErrMissingRequest = errors.New("missing request body")
	ErrAcountExists   = errors.New("account already present")
)

type SignupRequest struct {
	Identity string
	Password string
}

func (c *Client) Signup(r *SignupRequest) error {
	if r == nil {
		return ErrMissingRequest
	}

	current := c.repo.Get(&repository.Request{PrimaryID: r.Identity})
	if current.Item != nil {
		return ErrAcountExists
	}

	// Apply salt and hasing here before saving. Make a generic function that can be resued in the login
	// when you aren't brain dead
	if r := c.repo.Insert(&repository.Request{
		PrimaryID: r.Identity,
		Value: map[string]any{
			"hashed_password": "somepassword",
		},
	}); r.Error != nil {
		return r.Error
	}

	return nil
}
