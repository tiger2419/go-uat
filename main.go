package main

import (
	"fmt"
	"runtime"
	"sync"
	"uat/entity"
)

var maps map[int]int = make(map[int]int)
var resData = 0
var lock sync.Mutex
var group = sync.WaitGroup{}
var index = 10
var channel chan int = make(chan int)

func main() {

	fmt.Println("hello world")

	runtime.GOMAXPROCS(8)

	//group.Add(10)

	for i := 1; i <= 10; i++ {
		go goruntineDemo(i)
	}
	//time.Sleep(time.Second * 1)
	//group.Wait()
	for i := range maps {
		fmt.Println(i, "->", maps[i])
	}
	//index <- channel
	//fmt.Println(<-channel)
	go catChannel()
	<-channel

	fmt.Println("block.......")
	go func() {
		channel <- 20
	}()
	for i := range maps {
		fmt.Println(i, "->", maps[i])
	}

}

//func checkone(i int) int {
//	if i > 1 {
//		return checkone(i - 1)
//	}
//	return i
//}

func goruntineDemo(i int) {

	result := 1
	key := i
	for i > 0 {
		result = result * i
		i--
	}
	lock.Lock()
	maps[key] = result
	lock.Unlock()
	resData = result
	//group.Done()

	if key == index {

		channel <- 10
	}
}

func catChannel() {

	var cats chan interface{}
	cats = make(chan interface{}, 2)

	cats <- entity.Cat{"Tom", 5}
	cats <- entity.Cat{"Jerry", 3}

	cats2 := <-cats
	// 类型强转
	cat := cats2.(entity.Cat)

	fmt.Println(cat.Name)

	var f1 = 1
	var f2 = 1

	fmt.Println(sums(float64(f1), float64(f2)))
}

func sums(
	f1 float64, f2 float64) float64 {

	return f1 + f2
}
