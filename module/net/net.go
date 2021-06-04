package net

import (
	"fmt"
	"time"
	"net"
	"os"
	"os/exec"
	"strings"
	"strconv"

	"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
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
		availableInterfaces()
		os.Exit(0)
	}
	fmt.Println(byNameInterface)
}
