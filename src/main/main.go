package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"taskmanager"
	"time"
)

func main() {
	log.Println(`Start Main`)

	taskmanager.Work(LongTermTask)

	log.Println(`End Main`)
}

func LongTermTask(ctx context.Context) error {
	log.Println(`Long Term Task`)

	rand.Seed(time.Now().UnixNano())

	if c := rand.Intn(200); c < 100 {
		return errors.New(`Something happened. >> ` + strconv.Itoa(c))
	}

	time.Sleep(10 * time.Second)

	return nil
}
