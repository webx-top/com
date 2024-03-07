package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/webx-top/com"
)

func main() {
	ch := com.NewCmdChanReader()
	cmd := com.RunCmdWithReaderWriter([]string{`../process`}, ch)
	i := 0
	for {
		ch.Send(com.BreakLine)
		time.Sleep(5 * time.Second)
		i++
		if i >= 3 {
			err := cmd.Process.Kill()
			if err != nil && !errors.Is(err, os.ErrProcessDone) {
				log.Println(err)
			}
			break
		}
	}
	_ = cmd
}
