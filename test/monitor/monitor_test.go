package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/webx-top/com"
)

func newLogWriter() *logWriter {
	return &logWriter{
		b: []byte{},
	}
}

type logWriter struct {
	b []byte
	l sync.RWMutex
}

func (l *logWriter) Write(b []byte) (n int, e error) {
	l.b = append(l.b, b...)
	return len(b), nil
}

func (l *logWriter) String() string {
	l.l.Lock()
	r := string(l.b)
	l.b = []byte{}
	l.l.Unlock()
	return r
}

func TestMonitor(t *testing.T) {
	testDataDir := `../../testdata`
	var err error
	w := newLogWriter()
	log.SetOutput(w)
	os.Mkdir(testDataDir, os.ModePerm)
	me := &com.MonitorEvent{
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
	me.AddDir(testDataDir)
	time.Sleep(time.Second * 1)

	fmt.Println(``)
	fmt.Println(`CREATE: ` + testDataDir + `/aa`)
	os.Mkdir(testDataDir+`/aa`, 0666)
	time.Sleep(time.Second * 1)
	s := w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`CHMOD: ` + testDataDir + `/aa`)
	os.Chmod(testDataDir+`/aa`, os.ModePerm)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: CHMOD`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CHMOD`)
	}

	fmt.Println(``)
	fmt.Println(`WRITE: ` + testDataDir + `/aa/a.log`)
	err = ioutil.WriteFile(testDataDir+`/aa/a.log`, []byte(`test`), 0666)
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
	fmt.Println(`CHMOD: ` + testDataDir + `/aa/a.log`)
	os.Chmod(testDataDir+`/aa/a.log`, os.ModePerm)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: CHMOD`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CHMOD`)
	}

	fmt.Println(``)
	fmt.Println(`CREATE: ` + testDataDir + `/bb`)
	os.Mkdir(testDataDir+`/bb`, os.ModePerm)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`WRITE: ` + testDataDir + `/bb/b.log`)
	ioutil.WriteFile(testDataDir+`/bb/b.log`, []byte(`test`), 0666)
	time.Sleep(time.Second * 2)
	s = w.String()
	if !strings.HasSuffix(strings.TrimSpace(s), `: CREATE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: CREATE`)
	}

	fmt.Println(``)
	fmt.Println(`RENAME: ` + testDataDir + `/bb/b.log`)
	err = os.Rename(testDataDir+`/bb/b.log`, testDataDir+`/bb/bb.log`)
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
	fmt.Println(`REMOVE: ` + testDataDir + `/bb/bb.log`)
	os.Remove(testDataDir + `/bb/bb.log`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: ` + testDataDir + `/bb`)
	os.Remove(testDataDir + `/bb`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: ` + testDataDir + `/aa/a.log`)
	os.Remove(testDataDir + `/aa/a.log`)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}

	fmt.Println(``)
	fmt.Println(`REMOVE: ` + testDataDir + ``)
	os.RemoveAll(testDataDir + ``)
	time.Sleep(time.Second * 1)
	s = w.String()
	if !strings.Contains(s, `: REMOVE`) {
		fmt.Print(`[log]`, s)
		fmt.Println(`Fail: REMOVE`)
	}
}
