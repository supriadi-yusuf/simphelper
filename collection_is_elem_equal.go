package simphelper

import (
	"reflect"
)

func (c *myCollection) isElemEqualInMap(data interface{}) (result bool, err error) {

	return c.isEqualInMap(data)
}

func (c *myCollection) findIn(value reflect.Value, collection reflect.Value,
	start int, swapperFunc func(int, int)) (res bool) {

	interfaceVal := value.Interface()
	for i := start; i < collection.Len(); i++ {
		if interfaceVal == collection.Index(i).Interface() {
			// do swapping
			if i != start {
				//startValue := collection.Index(start)
				//iValue := collection.Index(i)
				//collection.Index(i).Set(startValue)
				//collection.Index(start).Set(iValue)
				swapperFunc(start, i)
			}

			return true
		}

	}

	return false
}

func (c *myCollection) isElemEqualInSliceOfArray(param interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(param)
	dataValue := reflect.ValueOf(c.data)

	dataType := reflect.TypeOf(c.data)
	dataLen := dataValue.Len()
	duplicateCollection := reflect.MakeSlice(dataType, dataLen, dataLen)
	reflect.Copy(duplicateCollection, dataValue)

	swapperFunc := reflect.Swapper(duplicateCollection.Interface())

	paramLen := paramValue.Len()
	for i := 0; i < paramLen; i++ {

		if !c.findIn(paramValue.Index(i), duplicateCollection, i, swapperFunc) {
			return false, nil
		}

		//if !c.findIn(dataValue.Index(i), paramValue) {
		//return false, nil
		//}
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

	if !func() bool {
		collectionValue := reflect.ValueOf(c.data)
		paramValue := reflect.ValueOf(data)
		return collectionValue.Len() == paramValue.Len()
	}() {
		return false, nil //errors.New("size is different")
	}

	if collectionType.Kind() == reflect.Map {
		return c.isElemEqualInMap(data)
	}

	return c.isElemEqualInSliceOfArray(data)

}
