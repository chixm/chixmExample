package taskmanager

import (
	"context"
	"log"
	"sync"
)

func Work(r Runnable) {
	log.Println(`Work Start`)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println(`Runnable Start`)
		if err := r(ctx); err != nil {
			log.Println(`Error Cancel of the task.[` + err.Error() + `]`)
			cancel()
		}
	}()
	wg.Wait()
}

type Runnable func(ctx context.Context) error
