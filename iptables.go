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
	if SaEnable {
		_ = ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "SYN,ACK", "-j", "NFQUEUE", "--queue-balance", "1000:1127")
	}
	if AEnable {
		_ = ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "ACK", "-j", "NFQUEUE", "--queue-balance", "2000:2127")
	}
	if PaEnable {
		_ = ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "PSH,ACK", "-j", "NFQUEUE", "--queue-balance", "3000:3127")
	}
	if FaEnable {
		_ = ipt.AppendUnique("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "FIN,ACK", "-j", "NFQUEUE", "--queue-balance", "4000:4127")
	}
}
func UnsetIptable(sport string) {
	ipt, err := iptables.New()
	if err != nil {
		fmt.Printf("Iptabels new error:%v", err)
		os.Exit(1)
	}
	if SaEnable {
		_ = ipt.Delete("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "SYN,ACK", "-j", "NFQUEUE", "--queue-balance", "1000:1127")
	}
	if AEnable {
		_ = ipt.Delete("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "ACK", "-j", "NFQUEUE", "--queue-balance", "2000:2127")
	}
	if PaEnable {
		_ = ipt.Delete("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "PSH,ACK", "-j", "NFQUEUE", "--queue-balance", "3000:3127")
	}
	if FaEnable {
		_ = ipt.Delete("filter", "OUTPUT", "-p", "tcp", "-m", "multiport", "--sport", sport, "--tcp-flags", "SYN,RST,ACK,FIN,PSH", "FIN,ACK", "-j", "NFQUEUE", "--queue-balance", "4000:4127")
	}
}
