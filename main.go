package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

const (
	YIN  = " ________          ________ \n|________|        |________|"
	YANG = " __________________________ \n|__________________________|"
)

var (
	UnitMap = map[int]string{
		0: YIN,
		1: YANG,
	}
)

type TurtleShell struct {
	Mutex  sync.Mutex
	Result []int
}

func main() {
	var t = TurtleShell{
		sync.Mutex{},
		[]int{},
	}
	var wg = sync.WaitGroup{}
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go func(i int) {
			defer wg.Done()
			rand.Seed(int64(i))
			x := rand.Intn(2)
			t.Mutex.Lock()
			t.Result = append(t.Result, x)
			t.Mutex.Unlock()
		}(i)
	}
	wg.Wait()
	gua := ""
	for _, x := range t.Result {
		gua += strconv.Itoa(x)
		fmt.Println(UnitMap[x])
	}
	fmt.Println(SixYaoExplain[gua].Name)
	fmt.Println(SixYaoExplain[gua].Description)
	fmt.Println(SixYaoExplain[gua].Explain)
}
