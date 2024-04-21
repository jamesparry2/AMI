package login

import "github.com/jamesparry2/aim/service-authenticate/app/repository"

type LoginI interface {
	Login(*LoginRequest) *LoginResponse
}

type Client struct {
	repo repository.Repository
}

type ClientOptions struct {
	Repo repository.Repository
}

func New(opts *ClientOptions) *Client {
	return &Client{
		repo: opts.Repo,
	}
}
