package main

import (
	"time"

	"github.com/webx-top/com"
)

func main() {
	ch := com.NewCmdChanReader()
	cmd := com.RunCmdWithReaderWriter([]string{`../process`}, ch)
	for {
		ch.Send(com.BreakLine)
		time.Sleep(5 * time.Second)
	}
	_ = cmd
}
