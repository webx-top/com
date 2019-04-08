package com

import (
	"context"
	"fmt"
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
