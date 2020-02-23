package golang

import (
	"fmt"
	"time"
)

func FirstWrite(writeChan <-chan bool, readChan chan<- bool, endChan <-chan bool) {
	for {
		select {
		case <-endChan:
			fmt.Println("close")
			return
		case <-writeChan:
			fmt.Println("write")
			time.Sleep(time.Second)
			readChan <- true
		}
	}

}

func SecondRead(writeChan chan<- bool, readChan <-chan bool, endChan <-chan bool) {
	for {
		select {
		case <-endChan:
			fmt.Println("close")
			return
		case <-readChan:
			fmt.Println("read")
			time.Sleep(time.Second)
			writeChan <- true
		}
	}

}

func Main() {
	//var wg sync.WaitGroup
	//wChan, rChan, eChan := make(chan bool, 1), make(chan bool, 1), make(chan bool)
	//finalChan := make(chan bool, 2)
	//go func(writeChan chan bool, readChan chan bool, endChan chan bool) {
	//	defer func() {
	//		finalChan <- true
	//	}()
	//	FirstWrite(writeChan, readChan, endChan)
	//}(wChan, rChan, eChan)
	//go func(writeChan chan bool, readChan chan bool, endChan chan bool) {
	//	defer func() {
	//		finalChan <- true
	//	}()
	//	SecondRead(writeChan, readChan, endChan)
	//}(wChan, rChan, eChan)
	//wChan <- true
	//time.Sleep(10 * time.Second)
	//close(eChan)
	//for i := 0; i < 2; i++ {
	//	<-finalChan
	//}
	//var pool sync.Pool
	//count := 0
	//pool.New = func() interface{} {
	//	count += 1
	//	return count
	//}
	//var wg sync.WaitGroup
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		defer wg.Done()
	//		a := pool.Get()
	//		fmt.Println(a)
	//		time.Sleep(time.Second * 3)
	//		pool.Put(a)
	//	}()
	//	time.Sleep(time.Second * 2)
	//}
	//wg.Wait()
	L := Constructor(2)
	L.Put(1, 1)
	L.Put(2, 2)
	L.Get(1)
	L.Put(3, 3)
	L.Get(2)
	L.Put(4, 4)
	L.Get(1)
	L.Get(3)
	L.Get(4)
}
