package com

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoOnce(t *testing.T) {
	doo := NewOnce()
	defer doo.Close()
	var count int

	wg := &sync.WaitGroup{}
	n := 100
	wg.Add(n)
	g := func(uid string) {
		defer wg.Done()
		flag := uid
		if doo.CanSet(flag) { // Step 1: CanSet
			t.Log(`DO_ONCE------->do: `, uid)
			time.Sleep(time.Millisecond * 20)
			count++
			doo.Release(flag) // Step 3: Release
			return
		}
		t.Log(`DO_ONCE------->wait: `, uid)
		doo.Wait(flag) // Step 2: Wait
		assert.Equal(t, 1, count)
	}
	for i := 0; i < n; i++ {
		go g(`tester`)
	}
	wg.Wait()
	assert.Equal(t, 1, count)
}
