package main

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
)

const IP_TO_FIND = "198.18.0.1"

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	var foundInterface *net.Interface
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.String() == IP_TO_FIND {
				foundInterface = &i
				break
			}
		}

		if foundInterface != nil {
			break
		}
	}

	if foundInterface != nil {
		fmt.Printf("Found the interface: %v\n", foundInterface.Name)

		cmd := exec.Command("sudo", "ip", "link", "set", foundInterface.Name, "mtu", "1500")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return
		}
		fmt.Println("Result: " + out.String())
	} else {
		fmt.Printf("No interface found with IP %v\n", IP_TO_FIND)
	}
}
