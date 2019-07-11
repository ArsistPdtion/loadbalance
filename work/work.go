package work

import (
	"fmt"
	"sync"
	"time"
)

type Work struct {
	M           *sync.RWMutex
	Channel chan int
	//Process     int
	//FreeProcess int

}

func NewWork(process int)*Work{
	return &Work{
		//Process:process,
		//FreeProcess:process,
		Channel:make(chan int,process),
	}
}

func (w *Work)Running(i int, wg sync.WaitGroup){
	defer wg.Done()
	select {
	case <-w.Channel:
		time.Sleep(time.Second)
		fmt.Printf("The %d Work Processing.\n", i)
	default:
		fmt.Println("w.Channel is kong")
	}

	//<-w.Channel

	//w.M.RLocker()
	//w.FreeProcess += 1
	//w.M.RUnlock()
}


