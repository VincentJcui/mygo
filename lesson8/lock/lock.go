package main

import (
	"fmt"
	"sync"
	"time"
)

type Accout struct {
	lock  sync.Mutex
	money int
}

func (a *Accout) DoPrepare() {
	time.Sleep(time.Second)
}

func (a *Accout) GetGongZi(n int) {
	a.money += n
}

func (a *Accout) GiveWife(n int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
		fmt.Println("give wife succes")
	} else {
		fmt.Println("give wife faild")
	}
}

func (a *Accout) Buy(n int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.money > n {
		a.DoPrepare()
		a.money -= n
		fmt.Println("bug succes")
	} else {
		fmt.Println("bug faild")
	}
}
func (a *Accout) left() int {
	return a.money
}

//方法一

func main() {
	var account Accout
	ch := make(chan int)
	account.GetGongZi(10)
	//定义channel
	go func() {
		go account.GiveWife(6)
		ch <- 0
	}()
	go func() {
		go account.Buy(5)
		ch <- 0
	}()
	deadline := time.After(time.Millisecond * 100)
	for i := 0; i < 2; i++ {
		select {
		case <-ch:
		case <-deadline:
			fmt.Println("deadline reach")
			return
		}
	}
	//从channel接收数据,如果收到两条数据,等待结束
	fmt.Println(account.left())

}

/*
//方法二
func main() {
	var account Accout
	account.GetGongZi(10)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//定义channel
	go func() {
		go account.GiveWife(6)
		wg.Done()
	}()
	go func() {
		go account.Buy(5)
		wg.Done()
	}()
	//从channel接收数据,如果收到两条数据,等待结束
	wg.Wait()
	fmt.Println(account.left())

}
*/
