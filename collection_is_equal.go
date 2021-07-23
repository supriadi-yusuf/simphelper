package simphelper

import (
	"reflect"
)

func (c *myCollection) isEqualInMap(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {

		if iter.Value().Interface() != paramValue.MapIndex(iter.Key()).Interface() {
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) isEqualInSliceOfArray(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)
	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		if paramValue.Index(i).Interface() != dataValue.Index(i).Interface() {
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) isEqualInStruct(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)
	dataValue := reflect.ValueOf(c.data)

	//paramType := reflect.TypeOf(data)
	//dataType := reflect.TypeOf(c.data)

	//if paramValue.NumField() != dataValue.NumField() { //number of field
	//	return false, nil
	//}

	for i := 0; i < dataValue.NumField(); i++ {

		//if paramType.Field(i).Name != dataType.Field(i).Name { // name
		//	return false, nil
		//}

		//if paramType.Field(i).Type != dataType.Field(i).Type { // type
		//	return false, nil
		//}

		if paramValue.Field(i).Interface() != dataValue.Field(i).Interface() { // value
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) IsEqual(data interface{}) (result bool, err error) {

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
		//log.Println("param type : ", paramType)
		//log.Println("collection type : ", collectionType)
		panic("types are different")
	}

	if collectionType.Kind() == reflect.Struct {
		return c.isEqualInStruct(data)
	}

	collectionValue := reflect.ValueOf(c.data)
	paramValue := reflect.ValueOf(data)
	if collectionValue.Len() != paramValue.Len() {
		return false, nil
	}

	if collectionType.Kind() == reflect.Map {
		return c.isEqualInMap(data)
	}

	return c.isEqualInSliceOfArray(data)

}
