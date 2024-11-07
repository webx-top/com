package com

import (
	"context"
	"fmt"
	"log"
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
	delay := time.Second * 2
	d := NewDelayOnce(delay, time.Hour, true)
	defer d.Close()
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	var counter atomic.Int32
	for i := 0; i < 10; i++ {
		log.Println(`Trigger key_normal`, delay)
		copyI := i
		isNew := d.Do(ctx, `key_normal`, func() error {
			defer wg.Done()
			log.Println(`>Execute key_normal<`, copyI)
			time.Sleep(time.Second * 3)
			counter.Add(1)
			return nil
		})
		if isNew {
			log.Println(`Add new`, copyI)
			wg.Add(1)
		}
		time.Sleep(time.Second)
	}
	wg.Wait()
	assert.Equal(t, int32(1), counter.Load())
}

func TestDelayOnceTimeout(t *testing.T) {
	delay := time.Second * 2
	d := NewDelayOnce(delay, time.Second*5, true)
	defer d.Close()
	ctx := context.TODO()
	wg := sync.WaitGroup{}
	var counter atomic.Int32
	for i := 0; i < 3; i++ {
		log.Println(`Trigger key_timeout`, delay)
		copyI := i
		isNew := d.DoWithState(ctx, `key_timeout`, func(isAbort func() bool) error {
			defer wg.Done()
			for j := 0; j < 4; j++ {
				if isAbort() {
					log.Println(`------> Stop key_timeout`)
					return nil
				}
				log.Println(`Execute key_timeout`, copyI, j)
				time.Sleep(time.Second * 5)
			}
			log.Println(`>Execute key_timeout<`)
			counter.Add(1)
			return nil
		})
		if isNew {
			log.Println(`Add new`, copyI)
			wg.Add(1)
		}
		time.Sleep(time.Second)
	}
	wg.Wait()
	assert.Equal(t, int32(1), counter.Load())
	isNew := d.DoWithState(ctx, `key_timeout`, func(isAbort func() bool) error {
		log.Println(`>Execute key_timeout<`)
		counter.Add(1)
		return nil
	})
	assert.True(t, isNew)
	time.Sleep(time.Second * 3)
	assert.Equal(t, int32(2), counter.Load())
}
