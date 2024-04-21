package repository

type Request struct {
	PrimaryID string
	Value     any
}

type Response struct {
	Item  any
	Error error
}

type Repository interface {
	Insert(*Request) *Response
	Get(*Request) *Response
}
