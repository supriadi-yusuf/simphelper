package simphelper

import (
	"reflect"
)

func (c *myCollection) mappingValueInMap(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newMapType := reflect.MapOf(reflect.TypeOf(c.data).Key(), funcValue.Type().Out(0))
	newMap := reflect.MakeMap(newMapType)

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {

		outPrms := funcValue.Call([]reflect.Value{iter.Value()})

		newMap.SetMapIndex(iter.Key(), outPrms[0])
	}

	return newMap.Interface(), nil
}

func (c *myCollection) mappingValueInSliceOfArray(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newSliceType := reflect.SliceOf(funcValue.Type().Out(0))
	newSlice := reflect.New(newSliceType).Elem()

	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		outPrms := funcValue.Call([]reflect.Value{dataValue.Index(i)})

		newSlice = reflect.Append(newSlice, outPrms[0])
	}

	return newSlice.Interface(), nil
}

//Mapping is map every value on collection based on function fmapping.
func (c *myCollection) MappingValue(fmapping interface{}) (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	//fcriteria must function with 1 input argument and 1 output argument.
	ftype := reflect.TypeOf(fmapping)
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

	if collectionType.Kind() == reflect.Map {
		return c.mappingValueInMap(fmapping)
	}

	return c.mappingValueInSliceOfArray(fmapping)
}
