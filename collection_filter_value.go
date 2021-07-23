package simphelper

import (
	"reflect"
)

func (c *myCollection) filterValueInMap(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newMap := reflect.MakeMap(reflect.TypeOf(c.data))

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {

		outPrms := funcValue.Call([]reflect.Value{iter.Value()})
		if !outPrms[0].Bool() {
			continue
		}

		newMap.SetMapIndex(iter.Key(), iter.Value())
	}

	return newMap.Interface(), nil
}

func (c *myCollection) filterValueInSliceOfArray(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newSlice := reflect.New(reflect.TypeOf(c.data)).Elem()

	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		outPrms := funcValue.Call([]reflect.Value{dataValue.Index(i)})
		if !outPrms[0].Bool() {
			continue
		}

		newSlice = reflect.Append(newSlice, dataValue.Index(i))
	}

	return newSlice.Interface(), nil
}

func (c *myCollection) FilterValue(fcriteria interface{}) (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	//fcriteria must function with 1 input argument and 1 output argument.
	//output argument must be boolean
	ftype := reflect.TypeOf(fcriteria)
	if ftype.Kind() != reflect.Func {
		panic("fcriteria argument must be function")
	}

	if ftype.NumIn() != 1 {
		panic("fcritera must have only one input parameter")
	}

	if ftype.In(0) != collectionType.Elem() {
		panic("element of collection and input parameter for fcriteria are different int type")
	}

	if ftype.NumOut() != 1 {
		panic("fcritera must have only one output parameter")
	}

	if ftype.Out(0) != reflect.TypeOf(true) {
		panic("output parameter for fcriteria must be boolean")
	}

	if collectionType.Kind() == reflect.Map {
		return c.filterValueInMap(fcriteria)
	}

	return c.filterValueInSliceOfArray(fcriteria)
}
