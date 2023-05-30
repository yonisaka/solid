package presentations

// RequestUser is a struct
type RequestUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified int    `json:"verified"`
}

// ResponseUser is a struct
type ResponseUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified string `json:"verified"`
}
