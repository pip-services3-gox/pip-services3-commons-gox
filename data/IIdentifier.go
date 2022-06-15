package data

type IIdentifier[T any] interface {
	Empty() bool
	Equals(T) bool
}
