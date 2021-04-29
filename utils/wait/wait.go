package wait

import "sync"

func Block(num int) {
	w := sync.WaitGroup{}
	w.Add(num)
	w.Wait()
}
