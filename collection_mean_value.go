package simphelper

import (
	"errors"
	"reflect"
)

func (c *myCollection) meanValueFromMap() (result float64, err error) {

	var total float64
	var currMap = reflect.ValueOf(c.data)

	if currMap.Len() == 0 {
		return 0, nil
	}

	iter := currMap.MapRange()
	for iter.Next() {

		res, _ := NewValue(iter.Value().Interface()).ToFloat()
		//res, err := NewValue(iter.Value().Interface()).ToFloat()
		//if err != nil {
		//	return 0, err
		//}

		total += res
	}

	return total / float64(currMap.Len()), nil
}

func (c *myCollection) meanValueFromSliceOfArray() (result float64, err error) {

	var total float64
	var currSlice = reflect.ValueOf(c.data)

	if currSlice.Len() == 0 {
		return 0, nil
	}

	for i := 0; i < currSlice.Len(); i++ {

		res, _ := NewValue(currSlice.Index(i).Interface()).ToFloat()
		//res, err := NewValue(currSlice.Index(i).Interface()).ToFloat()
		//if err != nil {
		//	return 0, err
		//}

		total += res
	}

	return total / float64(currSlice.Len()), nil
}

//Mean is count average of data in collection
func (c *myCollection) MeanValue() (result float64, err error) {

	defer GetErrorOnPanic(&err)

	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	// check element type
	switch collectionType.Elem().Name() {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		break
	default:
		return 0, errors.New("element type must be number")
	}

	if collectionType.Kind() == reflect.Map {
		return c.meanValueFromMap()
	}

	return c.meanValueFromSliceOfArray()
}
