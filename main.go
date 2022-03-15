package main

import (
	"fmt"
	"log"
)

//func init() {
//	rand.Seed(time.Now().Unix())
//}
//func sleep() {
//	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//}
//
//func reader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
//	sleep()
//	m.RLock()
//	c <- 1
//	sleep()
//	c <- -1
//	m.RUnlock()
//	wg.Done()
//}
//
//func writer(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
//	sleep()
//	m.Lock()
//	c <- 1
//	sleep()
//	c <- -1
//	m.Unlock()
//	wg.Done()
//}

func divide(a, b int) int {
	if b == 0 {
		panic(nil)
	}
	return a / b
}

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic occurred:", err)
		}
	}()
	fmt.Println(divide(1, 0))
}

func main() {

	divideByZero()
	fmt.Println("we survived dividing by zero")
	//var m sync.RWMutex
	//var rs, ws int
	//rsch := make(chan int)
	//wsch := make(chan int)
	//go func() {
	//	for {
	//
	//		select {
	//		case n := <-rsch:
	//			rs += n
	//
	//		case n := <-wsch:
	//			ws += n
	//		}
	//		fmt.Printf("%s%s\n", strings.Repeat("R", rs), strings.Repeat("W", ws))
	//	}
	//}()
	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go reader(rsch, &m, &wg)
	//
	//}
	//for i := 0; i < 3; i++ {
	//	wg.Add(1)
	//	go writer(wsch, &m, &wg)
	//}
	//wg.Wait()

}
