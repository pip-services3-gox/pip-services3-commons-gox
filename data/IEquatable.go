package data

type IEquatable[T any] interface {
	Equals(value T) bool
}
