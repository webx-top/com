package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/webx-top/com"
)

func main() {
	os.Mkdir(`testdata`, os.ModePerm)
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
	me.AddDir(`testdata`)

	os.Mkdir(`testdata/aa`, os.ModePerm)

	ioutil.WriteFile(`testdata/aa/a.log`, []byte(`test`), 0666)

	os.Chmod(`testdata/aa/a.log`, os.ModePerm)

	os.Mkdir(`testdata/bb`, os.ModePerm)

	ioutil.WriteFile(`testdata/bb/b.log`, []byte(`test`), 0666)

	os.Rename(`testdata/bb/b.log`, `testdata/bb/bb.log`)
	<-make(chan int)

}
