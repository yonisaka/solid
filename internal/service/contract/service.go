package contract

// Service is an interface
type Service interface {
	Serve(request Request) Response
}
