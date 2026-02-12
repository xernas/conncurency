package reader

import (
	"context"
	"fmt"
	"sync"
)

func Reader(inChan chan string, mtx *sync.Mutex, wg *sync.WaitGroup, ctx context.Context) {
	defer func() {
		wg.Done()
		ctx.Done()
	}()

	for v := range inChan {
		mtx.Lock()
		fmt.Println(v)
		mtx.Unlock()
	}
}
