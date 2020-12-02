package simphelper

import (
	"fmt"
	"log"
)

func ExampleNewCollection_removeIndex() {

	res, err := NewCollection([]int{1, 2, 3, 4}).RemoveIndex(0)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//[2 3 4]
}

func ExampleNewCollection_filterValue() {

	res, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(
		func(i int) bool {
			return i%2 == 1
		})
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//[1 3 5]
}

func ExampleNewCollection_mappingValue() {

	res, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(i int) bool {
		return i%2 == 1
	})
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//[true false true false true false]
}

func ExampleNewCollection_meanValue() {

	res, err := NewCollection([]int{1, 2, 3, 4}).MeanValue()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//2.5
}

func ExampleNewCollection_isEqual() {

	res, err := NewCollection([]int{1, 2, 3, 4}).IsEqual([]int{1, 2, 3, 4})
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//true
}

func ExampleNewCollection_isElemEqual() {

	res, err := NewCollection([]int{1, 2, 3, 4}).IsElemEqual([]int{1, 2, 3, 4})
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	res, err = NewCollection([]int{1, 2, 3, 4}).IsElemEqual([]int{1, 2, 4, 3})
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//true
	//true
}

func ExampleNewCollection_convElmToInterface() {

	res, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).ConvElmToInterface()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	fmt.Println(res)

	//Output:
	//[1 2 3 4 5 6]

}
