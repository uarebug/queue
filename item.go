package queue

type Item interface {
	Value() int // get priority of Item
}
