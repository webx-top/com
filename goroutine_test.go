package com

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLoop(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go Loop(ctx, func() error {
		fmt.Println(`Loop`, time.Now())
		return nil
	}, time.Second*1)
	time.Sleep(time.Second * 5)
	cancel()
}

func TestDelayOnce(t *testing.T) {
	d := NewDelayOnce(time.Second*2, time.Hour)
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		fmt.Println(`Trigger`, time.Now())
		isNew := d.Do(ctx, `key`, func() error {
			defer wg.Done()
			fmt.Println(`Execute`, time.Now())
			return nil
		})
		if isNew {
			wg.Add(1)
		}
		time.Sleep(time.Second * 2)
	}
	wg.Wait()
}
