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
	l.b = b
	return len(b), nil
}

func (l *logWriter) String() string {
	return string(l.b)
}

func TestMonitor(t *testing.T) {
	w := newLogWriter()
	log.SetOutput(w)
	os.Mkdir(`testdata`, os.ModePerm)
	me := &MonitorEvent{
		Debug: true,
		Create: func(file string) {
			fmt.Println(`create-file----------->`, file)
		},
		Delete: func(file string) {
			fmt.Println(`delete-file----------->`, file)
		},
		Modify: func(file string) {
			fmt.Println(`modify-file----------->`, file)
			if _, err := os.Stat(file); err != nil {
				if err == os.ErrNotExist {
					fmt.Println(`[Monitor]`, `Remove Watch:`, file)
				}
				fmt.Println(`[Monitor]`, err)
			} else {
				fmt.Println(`[Monitor]`, file)
			}
		},
		Chmod: func(file string) {
			fmt.Println(`chmod-file----------->`, file)
		},
		Rename: func(file string) {
			fmt.Println(`rename-file----------->`, file)
		},

		//文件夹事件
		DirCreate: func(file string) {
			fmt.Println(`create-dir----------->`, file)
		},
		DirDelete: func(file string) {
			fmt.Println(`delete-dir----------->`, file)
		},
		DirModify: func(file string) {
			fmt.Println(`modify-dir----------->`, file)
			if _, err := os.Stat(file); err != nil {
				if err == os.ErrNotExist {
					fmt.Println(`[Monitor]`, `Remove Watch:`, file)
				}
				fmt.Println(`[Monitor]`, err)
			} else {
				fmt.Println(`[Monitor]`, file)
			}
		},
		DirChmod: func(file string) {
			fmt.Println(`chmod-dir----------->`, file)
		},
		DirRename: func(file string) {
			fmt.Println(`rename-dir----------->`, file)
		},
	}
	me.Watch(`testdata`)
	time.Sleep(time.Second * 1)

	os.Mkdir(`testdata/aa`, os.ModePerm)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: CREATE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: CREATE`)
	}

	os.Chmod(`testdata/aa`, 0666)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: CHMOD`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: CHMOD`)
	}

	ioutil.WriteFile(`testdata/aa/a.log`, []byte(`test`), 0666)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: WRITE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: WRITE`)
	}

	os.Chmod(`testdata/aa/a.log`, os.ModePerm)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: CHMOD`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: CHMOD`)
	}

	os.Mkdir(`testdata/bb`, os.ModePerm)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: CREATE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: CREATE`)
	}

	ioutil.WriteFile(`testdata/bb/b.log`, []byte(`test`), 0666)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: WRITE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: WRITE`)
	}

	os.Rename(`testdata/bb/b.log`, `testdata/bb/bb.log`)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: RENAME`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: RENAME`)
	}

	os.Remove(`testdata/bb/bb.log`)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: REMOVE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: REMOVE`)
	}

	os.Remove(`testdata/bb`)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: REMOVE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: REMOVE`)
	}

	os.Remove(`testdata/aa/a.log`)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: REMOVE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: REMOVE`)
	}

	os.RemoveAll(`testdata`)
	time.Sleep(time.Second * 1)
	if strings.HasSuffix(strings.TrimSpace(w.String()), `: REMOVE`) {
		fmt.Println(`OK`)
	} else {
		fmt.Print(`[log]`, w.String())
		fmt.Println(`Fail: REMOVE`)
	}
}
