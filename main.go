package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

const (
	DAEMON  = "daemon"
	FOREVER = "forever"
)

func init() {
	if getProcessOwner() != "root\n" {
		log.Fatalln("Please run this program with root.")
	}
}

func main() {
	InitParams()
	if !Debug {
		log.SetOutput(ioutil.Discard)
	}
	if Daemon {
		SubProcess(StripSlice(os.Args, "-"+DAEMON))
		fmt.Printf("[*] Daemon running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
		os.Exit(0)
	} else if Forever {
		for {
			cmd := SubProcess(StripSlice(os.Args, "-"+FOREVER))
			fmt.Printf("[*] Forever running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
			cmd.Wait()
		}
		os.Exit(0)
	} else {
		SetIptable(Port)
		var wg sync.WaitGroup
		p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
			packetHandle(i.(int)*2 + 2)
			wg.Done()
		})
		defer p.Release()
		// Submit tasks one by one.
		for i := 0; i < 4; i++ {
			log.Println("Starting Task ", i)
			wg.Add(1)
			_ = p.Invoke(int(i))
		}
		wg.Wait()
	}
}
