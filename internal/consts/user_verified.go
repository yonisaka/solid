package consts

const (
	// UserVerified is a constant
	UserVerified = 1
	// UserRegistered is a constant
	UserRegistered = 2
)

var (
	// CodeToStrVerified is a map to convert code to string verified status
	CodeToStrVerified = map[int]string{
		UserVerified:   "Verified",
		UserRegistered: "Registered",
	}
)
