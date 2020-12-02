package simphelper

import (
	"log"
	"testing"
)

func Test_Convert_01(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).ConvElmToInterface()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := NewCollection(res.([]interface{})).IsEqual([]interface{}{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Convert_02(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection(
		map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}).ConvElmToInterface()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := NewCollection(res.(map[string]interface{})).IsEqual(
		map[string]interface{}{"one": 1, "two": 2, "three": 3, "four": 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}
