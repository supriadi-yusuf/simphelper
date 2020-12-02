package simphelper

import (
	"fmt"
	"log"
)

func ExampleMergeChanels() {

	sendValInt := func(unity []int) (ch <-chan int) {

		outCh := make(chan int)

		go func() {
			for _, val := range unity {

				outCh <- val
			}

			close(outCh)
		}()

		return outCh
	}

	data1 := []int{1, 2, 3, 4}
	data2 := []int{5, 6}
	data3 := []int{7, 6, 9}

	ch1 := sendValInt(data1)
	ch2 := sendValInt(data2)
	ch3 := sendValInt(data3)

	result, err := MergeChanels(ch1, ch2, ch3)
	if err != nil {
		log.Panicln(err.Error())
	}

	dataOut := make([]int, 0)
	for out := range result.(<-chan int) {
		dataOut = append(dataOut, out)
	}

	eq, err := NewCollection([]int{1, 2, 3, 4, 5, 6, 7, 6, 9}).IsElemEqual(dataOut)
	if err != nil {
		log.Panicln(err.Error())
	}

	fmt.Println(eq)
	//Output:
	//true
}
