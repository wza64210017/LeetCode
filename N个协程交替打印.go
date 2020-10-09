package main

import "fmt"

const COUNT = 3
const NUM = 5

func main() {
	chSlices := make([]chan int, COUNT)
	exitCh := make(chan int)
	result := 0

	for i := 0; i < COUNT; i++ {
		chSlices[i] = make(chan int, 1)
	}

	chSlices[COUNT-1] <- 1

	for i := 0; i < COUNT; i++ {
		curCh := chSlices[i]

		lastCh := chSlices[COUNT-1]
		if i != 0 {
			lastCh = chSlices[i-1]
		}

		go func(i int, curCh, lastCh chan int) {
			for {
				// 通过上个channel是否可写来控制本goroutine输出
				<-lastCh
				fmt.Println(fmt.Sprintf("thread%d: %d", i, result))

				result = result + 1
				if result > NUM {
					close(exitCh)
					break
				}

				curCh <- 1
			}
		}(i, curCh, lastCh)
	}

	<-exitCh
}
