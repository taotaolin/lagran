package main

import (
	"flag"
)

var Port string
var SaEnable bool
var WindowSa uint16
var AEnable bool
var WindowA uint16
var PaEnable bool
var WindowPa uint16
var FaEnable bool
var WindowFa uint16
var Debug bool
var Daemon bool
var Forever bool

func InitParams() (string, bool, uint16, bool, uint16, bool, uint16, bool, uint16, bool, bool, bool) {
	flag.StringVar(&Port, "p", "80", "The port of geneva, multi ports should be like \"80,443\"")

	flag.BoolVar(&SaEnable, "sa", true, "Enable to handle the packet when TCP flag SYN and ACK are 1.")
	var wsa int
	flag.IntVar(&wsa, "wsa", 1, "The window size of packet when TCP flag SYN and ACK are 1, available only when param -sa is true")

	flag.BoolVar(&AEnable, "a", true, "Enable to handle the packet when TCP flag ACK is 1")
	var wa int
	flag.IntVar(&wa, "wa", 1, "The window size of packet when TCP flag ACK is 1, available only when param -a is true")

	flag.BoolVar(&PaEnable, "pa", true, "Enable to handle the packet when TCP flag PSH and ACK are 1.")
	var wpa int
	flag.IntVar(&wpa, "wpa", 1, "The window size of packet when TCP flag PSH and ACK are 1, available only when param -pa is true")

	flag.BoolVar(&FaEnable, "fa", true, "Enable to handle the packet when TCP flag FIN and ACK are 1.")
	var wfa int
	flag.IntVar(&wfa, "wfa", 1, "The window size of packet when TCP flag FIN and ACK are 1, available only when param -fa is true")

	flag.BoolVar(&Debug, "debug", false, "Debug mode.")
	flag.BoolVar(&Daemon, "daemon", false, "Run in daemon")
	flag.BoolVar(&Forever, "forever", false, "Run forever")

	flag.Parse()
	WindowSa = uint16(wsa)
	WindowA = uint16(wa)
	WindowPa = uint16(wpa)
	WindowFa = uint16(wfa)
	return Port, SaEnable, WindowSa, AEnable, WindowA, PaEnable, WindowPa, FaEnable, WindowFa, Debug, Daemon, Forever
}
