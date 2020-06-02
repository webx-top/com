package com

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func newLogWriter() *logWriter {
	return &logWriter{
		b: []byte{},
	}
}

type logWriter struct {
	b []byte
}

func (l *logWriter) Write(b []byte) (n int, e error) {
	l.b = append(l.b, b...)
	return len(b), nil
}

func (l *logWriter) String() string {
	defer func() {
		l.b = []byte{}
	}()
	return string(l.b)
}

func TestMonitor(t *testing.T) {
	var err error
	w := newLogWriter()
	log.SetOutput(w)
	os.Mkdir(`testdata`, os.ModePerm)
	me := &MonitorEvent{
		Debug: true,
		Create: func(file string) {
			fmt.Println(`create----------->`, file)
		},
		Delete: func(file string) {
			fmt.Println(`delete----------->`, file)
		},
		Modify: func(file string) {
			fmt.Println(`modify----------->`, file)
		},
		Chmod: func(file string) {
			fmt.Println(`chmod----------->`, file)
		},
		Rename: func(file string) {
			fmt.Println(`rename----------->`, file)
		},
	}
	me.Watch()
	me.AddDir(`testdata`)
	time.Sleep(time.Second * 1)

	fmt.Println(``)
	fmt.Println(`CREATE: testdata/aa`)
	os.Mkdir(`testdata/aa`, 0666)
	time.Sleep(time.Second * 1)
	s := w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`CHMOD: testdata/aa`)
	os.Chmod(`testdata/aa`, os.ModePerm)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: CHMOD`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CHMOD`)
	}

	fmt.Println(``)
	fmt.Println(`WRITE: testdata/aa/a.log`)
	err = ioutil.WriteFile(`testdata/aa/a.log`, []byte(`test`), 0666)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`CHMOD: testdata/aa/a.log`)
	os.Chmod(`testdata/aa/a.log`, os.ModePerm)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: CHMOD`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CHMOD`)
	}

	fmt.Println(``)
	fmt.Println(`CREATE: testdata/bb`)
	os.Mkdir(`testdata/bb`, os.ModePerm)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`WRITE: testdata/bb/b.log`)
	ioutil.WriteFile(`testdata/bb/b.log`, []byte(`test`), 0666)
	time.Sleep(time.Second * 2)
	s = w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`RENAME: testdata/bb/b.log`)
	err = os.Rename(`testdata/bb/b.log`, `testdata/bb/bb.log`)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: RENAME`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: RENAME`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: testdata/bb/bb.log`)
	os.Remove(`testdata/bb/bb.log`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: testdata/bb`)
	os.Remove(`testdata/bb`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: testdata/aa/a.log`)
	os.Remove(`testdata/aa/a.log`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: testdata`)
	os.RemoveAll(`testdata`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}
}
