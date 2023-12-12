package model

type Role uint

const (
	Admin Role = iota + 1
	Customer
)

func (r Role) String() string {
	return "role"
}

func (r Role) Name() string {
	switch r {
	case Admin:
		return "Admin"
	case Customer:
		return "Customer"
	}

	return ""
}
