package main

import (
	"bufio"
	"context"
	"os"
	"time"

	"github.com/admpub/color"
	"github.com/webx-top/com"
)

func main() {
	args := com.ParseCmdArgs()
	com.Dump(args)
	ctx := context.Background()
	// Listen to keypress of "return" and restart the app automatically
	go func() {
		color.Blue(`[process] listen return ==> ` + com.NowStr())
		in := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				color.Yellow(`[process] reading ==> ` + com.NowStr())
				input, _ := in.ReadString(com.LF)
				if input == com.StrLF || input == com.StrCRLF {
					color.Green(`[process] restart ==> ` + com.NowStr())
					//panic(`[panic]`)
				} else {
					color.Blue(`[process] waiting ==> ` + com.NowStr())
				}
				time.Sleep(20 * time.Second)
				//os.Exit(1)
			}
		}
	}()

	<-make(chan struct{})
}
