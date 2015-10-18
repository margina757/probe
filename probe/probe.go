package probe

import (
	"fmt"
	"log"
	"net"
	"probe/probe/iface"
)

var (
	allActiveInterface []net.Interface
	allActiveIP        []*net.IPAddr
	allDisactiveIP     []*net.IPAddr
)

func Start() {
	var err error
	pcapInit()

	allActiveInterface, allActiveIP, allDisactiveIP, err = iface.AllInterface()
	if err != nil {
		log.Println(err)
	}

}

func printInterface() {
	fmt.Println("Active Interface:", len(allActiveInterface))
	for _, iface := range allActiveInterface {
		fmt.Println(iface.Name)
	}

	fmt.Println("Active IP:", len(allActiveIP))
	for _, ip := range allActiveIP {
		fmt.Println(ip.String())
	}

	fmt.Println("Disactive IP:", len(allActiveIP))
	for _, ip := range allDisactiveIP {
		fmt.Println(ip.String())
	}
}
