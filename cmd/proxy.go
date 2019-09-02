package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jhead/phantom/internal/proxy"
)

var bindAddressString string
var serverAddressString string
var versionString string
var nameString string
var usersInt int

func main() {
	bindArg := flag.String("bind", "0.0.0.0:19132", "IP address and port to listen on")
	serverArg := flag.String("server", "", "Bedrock/MCPE server IP address and port (ex: 1.2.3.4:19132)")
	versionArg := flag.String("version", "1.12.0", "Server version to display as")
	nameArg := flag.String("name", "Remote Server", "Name of the remote server")
	usersArg := flag.Int("users", 20, "Number of users of the remote server")

	flag.Usage = usage
	flag.Parse()

	if *serverArg == "" {
		flag.Usage()
		return
	}

	bindAddressString = *bindArg
	serverAddressString = *serverArg
	versionString = *versionArg
	nameString = *nameArg
	usersInt = *usersArg

	fmt.Printf("Starting up with remote server IP: %s\n", serverAddressString)

	proxyServer, err := proxy.New(bindAddressString, serverAddressString, versionString, nameString, usersInt)
	if err != nil {
		fmt.Printf("Failed to init server: %s\n", err)
		return
	}

	if err := proxyServer.Start(); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] -server <server-ip>\n\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}
