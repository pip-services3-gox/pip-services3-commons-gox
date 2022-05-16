package data

// IIdentifiable interface for data objects to operate with ids.
//	Example
//		type MyStruct struct {
//			...
//			id string
//		}
//
//		func (c *MyStruct) GetId() string {
//			return c.id
//		}
//		func (c *MyStruct) IsZeroId() bool {
//			return c.id == ""
//		}
//		func (c *MyStruct) IsEqualId(id string) bool {
//			return c.id == id
//		}
//		func (c *MyStruct) WithGeneratedId() *MyStruct {
//			c.id = cdata.IdGenerator.NextLong()
//			return c
//		}
type IIdentifiable[T any, K any] interface {
	GetId() K

	IsZeroId() bool

	WithGeneratedId() T

	IsEqualId(K) bool
}
