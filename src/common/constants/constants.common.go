package constants

type Gender string
type UserRole string

const (
	USER  UserRole = "USER"
	ADMIN UserRole = "ADMIN"
)

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
	OTHER  Gender = "OTHER"
)
