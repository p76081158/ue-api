package net

import (
	"fmt"
	"net"
	"log"
	"os"
)

// get all available network interface
func AvailableInterfaces() {
	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Println("Available network interfaces on this machine : ")
	for _, i := range interfaces {
		fmt.Printf("Name : %v \n", i.Name)
	}
}

// check interface exist or not
func CheckInterface(interface_name string) {
	byNameInterface, err := net.InterfaceByName(interface_name)

	if err != nil {
		fmt.Println(err, "[" + string(interface_name) +"]")
		fmt.Println("-----------------------------")
		AvailableInterfaces()
		os.Exit(0)
	}

	addrs, err := byNameInterface.Addrs()
	mac        := byNameInterface.HardwareAddr.String()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Info of Network Interface: ", interface_name)
	fmt.Println("IP  Address: ", addrs)
	fmt.Println("MAC Address: ", mac)
	fmt.Println("")
}