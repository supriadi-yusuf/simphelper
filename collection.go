package simphelper

type myCollection struct {
	data interface{}
}

//NewCollection is function creating object whose type is ICollection
//
//data must be array, slice, map or struct
func NewCollection(data interface{}) ICollection {
	collection := new(myCollection)
	collection.data = data
	return collection
}
