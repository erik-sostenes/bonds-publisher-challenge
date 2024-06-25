package domain

type Permission int

func (p Permission) Contains(p2 Permission) bool {
	return p&p2 == p2
}

const (
	Create Permission = 1 << iota
	Read
	Update
	Delete
)
