package lb

import (
	"errors"
	"fmt"
	"github.com/ArsistPdtion/workbook/LoadBalance/work"
	"sync"
	"time"
)

var wg sync.WaitGroup

type LoadBlance struct {
	Works []work.Work
}

//type LBInterface interface {
//	submit(f func())
//	wait()
//}

func LB(args []int) *LoadBlance {
	var lb LoadBlance
	for _, arg := range args {
		tWork := work.NewWork(arg)
		lb.Works = append(lb.Works, *tWork)
	}
	return &lb
}

func (lb *LoadBlance) Submit(i int) {
	maxIndex, err := MaxListIndex(lb.Works)
	if err != nil {
		time.Sleep(time.Second)
		maxIndex, err = MaxListIndex(lb.Works)
	}
	fmt.Println("maxIndex is: ", maxIndex)
	//lb.Works[maxIndex].M.RLocker()
	select {
	case lb.Works[maxIndex].Channel <- 1:
		wg.Add(1)
		go lb.Works[maxIndex].Running(maxIndex, wg)
	default:
		fmt.Printf("lb.works %d work is full\n", maxIndex)
	}
	//lb.Works[maxIndex].Channel <- 1
	//lb.Works[maxIndex].M.RUnlock()

}

func (lb *LoadBlance) Wait() {
	//time.Sleep(time.Second * 5)
	wg.Wait()
}

func MaxListIndex(lists []work.Work) (int, error) {
	maxIndex := 0
	maxNumber := cap(lists[0].Channel) - len(lists[0].Channel)
	for i, l := range lists {
		fmt.Println("cap and len is:", cap(l.Channel), len(l.Channel))
		if cap(l.Channel)-len(l.Channel) > maxNumber && maxNumber > 0 {
			maxNumber = cap(l.Channel) - len(l.Channel)
			maxIndex = i
		}
	}
	fmt.Println("maxNUmber is: ", maxNumber)
	if maxNumber > 0 {
		return maxIndex, nil
	} else {
		return 0, errors.New("all is full")
	}

}
