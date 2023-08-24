package queue

type Item interface {
	Identify() string
	Value() int // get priority of Item
}
