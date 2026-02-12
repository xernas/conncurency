package reader

import (
	"context"
	"fmt"
	"sync"
)

func Reader(inChan chan string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for v := range inChan {
		fmt.Println(v)
	}
}
