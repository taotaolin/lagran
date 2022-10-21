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
		UnsetIptable(Port)
		SetIptable(Port)
		var wg sync.WaitGroup
		if SaEnable {
			p1, _ := ants.NewPoolWithFunc(128, func(i interface{}) {
				packetHandle(i.(int))
				wg.Done()
			})
			defer p1.Release()
			// Submit tasks one by one.
			log.Println("Starting Task p1")
			for i := 1000; i < 1128; i++ {
				wg.Add(1)
				_ = p1.Invoke(int(i))
			}
		}
		if AEnable {
			p2, _ := ants.NewPoolWithFunc(128, func(i interface{}) {
				packetHandle(i.(int))
				wg.Done()
			})
			defer p2.Release()
			// Submit tasks one by one.
			log.Println("Starting Task p2")
			for i := 2000; i < 2128; i++ {
				wg.Add(1)
				_ = p2.Invoke(int(i))
			}
		}
		if PaEnable {
			p3, _ := ants.NewPoolWithFunc(128, func(i interface{}) {
				packetHandle(i.(int))
				wg.Done()
			})
			defer p3.Release()
			// Submit tasks one by one.
			log.Println("Starting Task p3")
			for i := 3000; i < 3128; i++ {
				wg.Add(1)
				_ = p3.Invoke(int(i))
			}
		}
		if FaEnable {
			p4, _ := ants.NewPoolWithFunc(128, func(i interface{}) {
				packetHandle(i.(int))
				wg.Done()
			})
			defer p4.Release()
			// Submit tasks one by one.
			log.Println("Starting Task p4")
			for i := 4000; i < 4128; i++ {
				wg.Add(1)
				_ = p4.Invoke(int(i))
			}
		}
		wg.Wait()
	}
}
