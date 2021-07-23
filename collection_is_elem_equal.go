package simphelper

import (
	"reflect"
)

func (c *myCollection) isElemEqualInMap(data interface{}) (result bool, err error) {

	return c.isEqualInMap(data)
}

func (c *myCollection) findIn(value reflect.Value, collection reflect.Value) (res bool) {

	interfaceVal := value.Interface()
	for i := 0; i < collection.Len(); i++ {
		if interfaceVal == collection.Index(i).Interface() {
			return true
		}

	}

	return false
}

func (c *myCollection) isElemEqualInSliceOfArray(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)
	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		if !c.findIn(paramValue.Index(i), dataValue) {
			return false, nil
		}

		if !c.findIn(dataValue.Index(i), paramValue) {
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) isElemEqualInStruct(data interface{}) (result bool, err error) {

	return c.isEqualInStruct(data)
}

func (c *myCollection) IsElemEqual(data interface{}) (result bool, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map && collectionType.Kind() != reflect.Struct {
		panic("collection must be array, slice, struct or map")
	}

	//check type of input parameter
	paramType := reflect.TypeOf(data)
	/*if paramType.Kind() != reflect.Array && paramType.Kind() != reflect.Slice &&
		paramType.Kind() != reflect.Map && paramType.Kind() != reflect.Struct {
		panic("input param must be array, slice, struct or map")
	}*/

	if paramType != collectionType {
		panic("types are different")
	}

	if collectionType.Kind() == reflect.Struct {
		return c.isElemEqualInStruct(data)
	}

	collectionValue := reflect.ValueOf(c.data)
	paramValue := reflect.ValueOf(data)
	if collectionValue.Len() != paramValue.Len() {
		return false, nil
	}

	if collectionType.Kind() == reflect.Map {
		return c.isElemEqualInMap(data)
	}

	return c.isElemEqualInSliceOfArray(data)

}
