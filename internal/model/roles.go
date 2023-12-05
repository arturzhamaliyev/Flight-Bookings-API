package model

type Role uint

func (r Role) String() string {
	return "role"
}

const (
	Admin Role = iota + 1
	Customer
)
