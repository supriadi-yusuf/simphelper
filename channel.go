package simphelper

import (
	"errors"
	"reflect"
	"sync"
)

// MergeChanels is function to merge several receiver channels into one receiver channel.
func MergeChanels(channels ...interface{}) (chn interface{}, err error) {

	defer GetErrorOnPanic(&err)

	totChan := len(channels) //total channel element that we want to join
	if totChan == 0 {
		return nil, errors.New("number of joined channel may not be zero")
	}

	var refVal0 reflect.Value
	var elType0 reflect.Type

	for i, channel := range channels {
		refVal := reflect.ValueOf(channel)
		if i == 0 {
			refVal0 = refVal
		}

		if refVal.Kind() != reflect.Chan {
			return nil, errors.New("input parameter must be channel")
		}

		elType := refVal.Type().Elem()
		if i == 0 {
			elType0 = elType
		}

		if elType != elType0 {
			return nil, errors.New("channel's type is different")
		}
	}

	elType := refVal0.Type().Elem()                      //type of channel
	chType := reflect.ChanOf(reflect.BothDir, elType)    //channel type
	chRcvType := reflect.ChanOf(reflect.RecvDir, elType) //channel type : receiver channel

	outChan := reflect.MakeChan(chType, 0)      // create channel
	outRcvChan := reflect.New(chRcvType).Elem() //create empty receiver channel
	outRcvChan.Set(outChan)                     //convert channel into receiver channel

	wg := new(sync.WaitGroup)
	wg.Add(totChan)

	for _, chanel := range channels {

		go func(ch reflect.Value) {

			defer func() {
				//fmt.Println("selesai - START")
				wg.Done()
				recover()
				//fmt.Println("selesai - DONE")
			}()

			for {

				val, ok := ch.Recv() //receive value from input channel
				if ok {
					outChan.Send(val) //send value into output channel
				} else {
					// chanel is closed
					break
				}

			}

		}(reflect.ValueOf(chanel))

	}

	go func() {
		wg.Wait()
		outChan.Close() //close channel
	}()

	return outRcvChan.Interface(), nil
}
