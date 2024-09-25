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

func TestParseArgs(t *testing.T) {
	parts := ParseArgs(`tower.exe -c tower.yaml -p "eee\"ddd" -t aaaa`)
	expected := []string{"tower.exe", "-c", "tower.yaml", "-p", "eee\"ddd", "-t", "aaaa"}
	assert.Equal(t, expected, parts)

	parts = ParseArgs(`tower.exe -c    tower.yaml -p "eee\"ddd" -t aaaa`) // more space
	assert.Equal(t, expected, parts)

	parts = ParseArgs(`tower.exe -c=tower.yaml -p="eee\"ddd" -t=aaaa`)
	assert.Equal(t, expected, parts)
	parts = ParseArgs(`tower.exe -c=tower.yaml -p='eee"ddd' -t=aaaa`)
	assert.Equal(t, expected, parts)

	parts = ParseArgs(`tower.exe -c	  	tower.yaml 		-p 	"eee\"ddd" 	-t 	aaaa`) // space and tab
	assert.Equal(t, expected, parts)

	parts = ParseArgs(`tower.exe -c tower.yaml -p   'eee\'ddd' -t aaaa`)
	assert.Equal(t, []string{"tower.exe", "-c", "tower.yaml", "-p", "eee'ddd", "-t", "aaaa"}, parts)
	parts = ParseArgs(`tower.exe -c tower.yaml -p   '	eee\'ddd ' -t aaaa`)
	assert.Equal(t, []string{"tower.exe", "-c", "tower.yaml", "-p", "	eee'ddd ", "-t", "aaaa"}, parts)
}

func TestParseFields(t *testing.T) {
	parts := ParseFields(`drwxr-xr-x   1 root root    0 2023-11-19 04:18 'test test2'`)
	expected := []string{"drwxr-xr-x", "1", "root", "root", "0", "2023-11-19", "04:18", "test test2"}
	assert.Equal(t, expected, parts)
}

func TestParseEnvVar(t *testing.T) {
	os.Setenv(`TESTENV`, `1`)
	v := ParseEnvVar(`ab{$TESTENV}c`)
	assert.Equal(t, `ab1c`, v)
	v = ParseEnvVar(`ab{$NOTEXISTS:ok}c`)
	assert.Equal(t, `abokc`, v)
}

func TestParseWindowsEnvVar(t *testing.T) {
	os.Setenv(`TESTENV`, `2`)
	v := ParseWindowsEnvVar(`ab{%TESTENV%}c`)
	assert.Equal(t, `ab2c`, v)
	v = ParseWindowsEnvVar(`ab{%NOTEXISTS:ok%}c`)
	assert.Equal(t, `abokc`, v)
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

func TestParseCmdArgs(t *testing.T) {
	args := ParseCmdArgs(`c.exe`, `--a`, `b`)
	assert.Equal(t, map[string]string{"a": "b"}, args)
	args = ParseCmdArgs(`c.exe`, `--a`)
	assert.Equal(t, map[string]string{"a": ""}, args)
}
