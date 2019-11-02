package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"keepGuestOn"
	icon "keepGuestOn/icon"
	notify "keepGuestOn/notifyicon"
	"net"
	"os"
	"time"
	"github.com/eiannone/keyboard"
)

const IP_GUEST = "192.168.16.1" //ip guest (wifi connection)

var cosntinue bool = false;
func main () {


	notify.StartNotification("Staying alive ha! ha ha!","pinging to IGuest","[keep alive wifi connection]")


	startIntervalPinging()

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

/*	for {
		char, key, err := keyboard.GetKey()
		if (err != nil) {
			panic(err)
		} else if (key == keyboard.KeyEsc) {
			break
		}

		cosntinue = false;
		fmt.Printf("Reset por la tecla : %q\r\n", char)
	}*/

	keepGuestOn.Run(onReady, onExit)
	//-- running  :u

}

func startIntervalPinging (){

	 setInterval(func() {
	 	//if (cosntinue){
		got, result := keepAlive(IP_GUEST)
		if got != nil {
			//-
		}else {
			fmt.Print(result)
		}
	 //}
	}, 30000, true)

}

func onReady() {
	keepGuestOn.SetIcon(icon.Data)
	keepGuestOn.SetTitle("Keep Guest Alive")
	keepGuestOn.SetTooltip("Manten esa wea viva arre :u")
	mQuit := keepGuestOn.AddMenuItem("Quit", "Salir :c")

	// Sets the icon of a menu item. Only available on Mac.
	mQuit.SetIcon(icon.Data)
}

func onExit() {
	os.Exit(1)
}



func keepAlive(ipDirection string) (error,  string) {
	//await
	p := fastping.NewPinger()
	//ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	ra, err := net.ResolveIPAddr("ip4:icmp", IP_GUEST)
	if err != nil {
		return err, ""
		//os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		notify.StartNotification("Connection Online","[-----]","[keep alive wifi connection]")
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		return err, ""
	}
	return nil, ""
}

func setInterval(someFunc func(), milliseconds int, async bool) chan bool {

	// How often to fire the passed in function
	// in milliseconds
	interval := time.Duration(milliseconds) * time.Millisecond

	// Setup the ticket and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	// Put the selection in a go routine
	// so that the for loop is none blocking
	go func() {
		for {

			select {
			case <-ticker.C:
				if async {
					// This won't block
					go someFunc()
				} else {
					// This will block
					someFunc()
				}
			case <-clear:
				ticker.Stop()
				return
			}

		}
	}()

	// We return the channel so we can pass in
	// a value to it to clear the interval
	return clear

}
