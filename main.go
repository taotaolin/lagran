package main

import (
	"fmt"
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
		p := NewPool()
		StartPool(p)
		var wg sync.WaitGroup
		wg.Add(4)
		task1 := func() {
			log.Println("Starting goroutine 1.")
			packetHandle(2)
			wg.Done()
		}
		task2 := func() {
			log.Println("Starting goroutine 2.")
			defer wg.Done()
			packetHandle(4)
		}
		task3 := func() {
			log.Println("Starting goroutine 3.")
			packetHandle(6)
			wg.Done()
		}
		task4 := func() {
			log.Println("Starting goroutine 4.")
			defer wg.Done()
			packetHandle(8)
		}
		p.Submit(task1)
		p.Submit(task2)
		p.Submit(task3)
		p.Submit(task4)
		wg.Wait()
	}
}
