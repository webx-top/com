package com

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	defer d.Close()
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	var lastTime time.Time
	var execTime time.Time
	var counter atomic.Int32
	for i := 0; i < 10; i++ {
		log.Println(`Trigger key_normal`)
		lastTime = time.Now()
		isNew := d.Do(ctx, `key_normal`, func() error {
			defer wg.Done()
			log.Println(`>Execute key_normal<`)
			execTime = time.Now()
			time.Sleep(time.Second * 3)
			counter.Add(1)
			return nil
		})
		if isNew {
			wg.Add(1)
		}
		//time.Sleep(time.Second * 2)
	}
	wg.Wait()
	assert.Equal(t, float64(2), math.Floor(execTime.Sub(lastTime).Seconds()))
	assert.Equal(t, int32(1), counter.Load())
}

func TestDelayOnceTimeout(t *testing.T) {
	d := NewDelayOnce(time.Second*2, time.Second*5, true)
	defer d.Close()
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	var lastTime time.Time
	var execTime time.Time
	var counter atomic.Int32
	for i := 0; i < 3; i++ {
		log.Println(`Trigger key_timeout`)
		lastTime = time.Now()
		isNew := d.DoWithState(ctx, `key_timeout`, func(isAbort func() bool) error {
			defer wg.Done()
			execTime = time.Now()
			for i := 0; i < 4; i++ {
				if isAbort() {
					log.Println(`------> Stop key_timeout`)
					return nil
				}
				log.Println(`Execute key_timeout`, i)
				time.Sleep(time.Second * 5)
			}
			log.Println(`>Execute key_timeout<`)
			counter.Add(1)
			return nil
		})
		if isNew {
			wg.Add(1)
		}
		//time.Sleep(time.Second * 6)
	}
	wg.Wait()
	assert.Equal(t, float64(2), math.Floor(execTime.Sub(lastTime).Seconds()))
	assert.Equal(t, int32(1), counter.Load())
}
