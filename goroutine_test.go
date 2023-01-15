package com

import (
	"context"
	"fmt"
	"log"
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

func TestDelayOnceNormal(t *testing.T) {
	d := NewDelayOnce(time.Second*2, time.Hour, true)
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		log.Println(`Trigger key_normal`)
		isNew := d.Do(ctx, `key_normal`, func() error {
			defer wg.Done()
			log.Println(`Execute key_normal`)
			return nil
		})
		if isNew {
			wg.Add(1)
		}
		time.Sleep(time.Second * 2)
	}
	wg.Wait()
}

func TestDelayOnceTimeout(t *testing.T) {
	d := NewDelayOnce(time.Second*2, time.Second*5, true)
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		log.Println(`Trigger key_timeout`)
		isNew := d.DoWithState(ctx, `key_timeout`, func(isAbort func() bool) error {
			defer wg.Done()
			for i := 0; i < 4; i++ {
				if isAbort() {
					log.Println(`------> Stop key_timeout`)
					return nil
				}
				log.Println(`Execute key_timeout`, i)
				time.Sleep(time.Second * 5)
			}
			log.Println(`Execute key_timeout`)
			return nil
		})
		if isNew {
			wg.Add(1)
		}
		time.Sleep(time.Second * 6)
	}
	wg.Wait()
}
