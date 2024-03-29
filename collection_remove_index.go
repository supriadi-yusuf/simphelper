package simphelper

import (
	"reflect"
)

func (c *myCollection) removeIndexInMap(index interface{}) (result interface{}, err error) {

	newMap := reflect.MakeMap(reflect.TypeOf(c.data))

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {
		key := iter.Key()
		if key.Type() != reflect.TypeOf(index) {
			panic("map key and index have different data type")
		}

		if key.Interface() == index {
			continue
		}

		newMap.SetMapIndex(key, iter.Value())
	}

	return newMap.Interface(), nil
}

func (c *myCollection) removeIndexInSliceOfArray(index int) (result interface{}, err error) {

	newSlice := reflect.New(reflect.TypeOf(c.data)).Elem()

	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {
		if i == index {
			continue
		}

		newSlice = reflect.Append(newSlice, dataValue.Index(i))
	}

	return newSlice.Interface(), nil
}

func (c *myCollection) RemoveIndex(index interface{}) (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	indexType := reflect.TypeOf(index)
	if indexType.Kind() == reflect.Array || indexType.Kind() == reflect.Slice ||
		indexType.Kind() == reflect.Map || indexType.Kind() == reflect.Struct {
		panic("index must be not array, slice, map or struct")
	}

	if collectionType.Kind() == reflect.Map {
		return c.removeIndexInMap(index)
	}

	return c.removeIndexInSliceOfArray(index.(int))
}
