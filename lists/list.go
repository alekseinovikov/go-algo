package lists

type List interface {
	Add(elem interface{})
	GetAt(index int) (interface{}, bool)
	AddAt(index int, elem interface{}) error
	RemoveLast() bool
	RemoveFirst() bool
}
