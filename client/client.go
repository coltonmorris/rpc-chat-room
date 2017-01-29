package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strings"
	"time"
)

func main() {
	// handle command line arguments
	handlePtr := flag.String("handle", "Colton", "-handle=\"your handle/name\"")
	hostPtr := flag.String("host", "127.0.0.1", "-host=\"the host address\"")
	portPtr := flag.String("port", "3410", "-port=\"the port number\"")
	flag.Parse()

	// initialize RPC client
	client, err := rpc.DialHTTP("tcp", net.JoinHostPort(*hostPtr, *portPtr))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	fmt.Printf("%T", client)
	fmt.Println("****************************************************")
	fmt.Printf("\tConnecting to %v:%v...\n", *hostPtr, *portPtr)
	fmt.Println("****************************************************")

	// log us in first
	Login(client, *handlePtr)

	// start a routine that always checks for incoming messages
	go func() {
		for {
			CheckMessages(client, *handlePtr)
			time.Sleep(time.Second)
		}
	}()

	// parse the commandline
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		line := strings.SplitAfterN(input, " ", 3)

		switch strings.ToLower(strings.TrimSpace(line[0])) {
		case "tell":
			if len(line) == 3 {
				TellUser(client, *handlePtr, strings.TrimSpace(line[1]), line[2])
			}
		case "say":
			if len(line) == 2 {
				Say(client, *handlePtr, line[1])
			}
			if len(line) == 3 {
				Say(client, *handlePtr, line[1]+line[2])
			}
		case "list":
			ListUsers(client, *handlePtr)
		case "quit":
			Logout(client, *handlePtr)
		case "help":
			Help()
		case "shutdown":
			Shutdown(client, *handlePtr)
		default:
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}
