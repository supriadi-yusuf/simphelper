package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeChanels_NoChannel(t *testing.T) {
	_, err := MergeChanels()
	assert.NotNil(t, err, "it should be error")
}

func Test_MergeChanels_NotChannel(t *testing.T) {
	_, err := MergeChanels(1, 2, 3)
	assert.NotNil(t, err, "it should be error")
}

func Test_MergeChanels_MixWithNotChannel(t *testing.T) {
	ch1 := make(chan int)
	defer close(ch1)

	_, err := MergeChanels(ch1, 2, 3)
	assert.NotNil(t, err, "it should be error")
}

func Test_MergeChanels_OneChannel(t *testing.T) {

	err := func() error {
		ch1 := make(chan int)
		defer close(ch1)

		_, err := MergeChanels(ch1)
		return err
	}()

	assert.Nil(t, err, "it should be not error")
}

func Test_MergeChanels_OneChannelWithValue(t *testing.T) {

	outCh, err := func() (<-chan int, error) {
		ch1 := make(chan int)

		go func() {
			defer close(ch1)
			ch1 <- 10
		}()

		outChan, err := MergeChanels(ch1)
		return outChan.(<-chan int), err
	}()

	for val := range outCh {
		val = val + 0
	}

	assert.Nil(t, err, "it should be not error")
}

func Test_MergeChanels_ManyChannels(t *testing.T) {

	err := func() error {
		ch1 := make(chan int)
		defer close(ch1)

		ch2 := make(chan int)
		defer close(ch2)

		ch3 := make(chan int)
		defer close(ch3)

		_, err := MergeChanels(ch1, ch2, ch3)
		return err
	}()

	assert.Nil(t, err, "it should be not error")
}

func Test_MergeChanels_ManyChannelsWithValues(t *testing.T) {

	outChan, err := func() (<-chan int, error) {

		ch1 := make(chan int)

		go func() {
			defer close(ch1)
			ch1 <- 1
		}()

		ch2 := make(chan int)

		go func() {
			defer close(ch2)
			ch2 <- 2
		}()

		ch3 := make(chan int)

		go func() {
			defer close(ch3)
			ch3 <- 3
		}()

		outChInt, err := MergeChanels(ch1, ch2, ch3)

		return outChInt.(<-chan int), err
	}()

	results := make([]int, 0)
	for val := range outChan {
		results = append(results, val)
	}

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, len(results), 3, "total results must be 3")
}

func Test_MergeChanels_MixChannels(t *testing.T) {

	err := func() error {
		ch1 := make(chan int)
		defer close(ch1)

		ch2 := make(chan bool)
		defer close(ch2)

		ch3 := make(chan string)
		defer close(ch3)

		_, err := MergeChanels(ch1, ch2, ch3)
		return err
	}()

	assert.NotNil(t, err, "it should be not error")
}
