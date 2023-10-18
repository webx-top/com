package com

import (
	"bufio"
	"context"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnvVar(t *testing.T) {
	os.Setenv(`TESTENV`, `1`)
	v := ParseEnvVar(`ab{$TESTENV}c`)
	assert.Equal(t, `ab1c`, v)
}

func TestParseWindowsEnvVar(t *testing.T) {
	os.Setenv(`TESTENV`, `2`)
	v := ParseWindowsEnvVar(`ab{%TESTENV%}c`)
	assert.Equal(t, `ab2c`, v)
}

func TestCmdChanReader(t *testing.T) {
	c := NewCmdChanReader()
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		c.Close()
	}()
	wg := sync.WaitGroup{}
	wg.Add(5)
	go func() {
		in := bufio.NewReader(c)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				input, err := in.ReadString(LF)
				if err != nil && err != io.EOF {
					t.Log(err.Error())
					return
				}
				wg.Done()
				t.Log(input)
			}
		}
	}()
	c.SendStringAndWait("OK\n")
	c.SendStringAndWait("OK1\n")
	c.SendStringAndWait("OK2\n")
	c.SendStringAndWait("OK3\n")
	c.SendStringAndWait("OK4\n")
	wg.Wait()
}
