package main

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
	"os"
)

func SetIptable(sport string) {
	ipt, err := iptables.New()
	if err != nil {
		fmt.Printf("Iptabels new error:%v", err)
		os.Exit(1)
	}
	ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "SYN,ACK", "-j", "NFQUEUE", "--queue-num", "2")
	ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "ACK", "-j", "NFQUEUE", "--queue-num", "4")
	ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "PSH,ACK", "-j", "NFQUEUE", "--queue-num", "6")
	ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "FIN,ACK", "-j", "NFQUEUE", "--queue-num", "8")
}
func UnsetIptable(sport string) {
	ipt, err := iptables.New()
	if err != nil {
		fmt.Printf("Iptabels new error:%v", err)
		os.Exit(1)
	}
	ipt.Delete("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "SYN,ACK", "-j", "NFQUEUE", "--queue-num", "2")
	ipt.Delete("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "ACK", "-j", "NFQUEUE", "--queue-num", "4")
	ipt.Delete("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "PSH,ACK", "-j", "NFQUEUE", "--queue-num", "6")
	ipt.Delete("filter", "OUTPUT", "-p", "tcp", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "FIN,ACK", "-j", "NFQUEUE", "--queue-num", "8")
}
