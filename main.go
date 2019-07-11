package main

import (
	"github.com/ArsistPdtion/workbook/LoadBalance/lb"
)

func main() {

	loadbalance := lb.LB([]int{20, 5, 8, 4})

	for i := 0; i < 18; i++ {
		loadbalance.Submit(i)
	}
	loadbalance.Wait()
}
