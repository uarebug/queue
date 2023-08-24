package queue

type Item interface {
	ID() string
	Value() int // get priority of Item
}
