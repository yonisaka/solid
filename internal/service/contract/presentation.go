package contract

// Response is a struct
type Response struct {
	Message string
	Data    interface{}
}

// Request is a struct
type Request struct {
	Data interface{}
}
