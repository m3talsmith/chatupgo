package main

import (
	"bufio"
	"chatupgo/app"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

var (
	peer string
	port int
	name string
)

func init() {
	flag.StringVar(&peer, "peer", "localhost:3000", "Chatup user peer")
	flag.IntVar(&port, "port", 0, "Chatup service port")
	flag.StringVar(&name, "name", "anonymous", "Chatup user name")
	flag.Parse()
}

func main() {
	if port == 0 {
		port = randomPort()
	}

	host := fmt.Sprintf("localhost:%d", port)

	fmt.Fprintf(os.Stdout, "peer: %s, port: %d, host: %s, name: %s\n", peer, port, host, name)

	go func() {
		server := app.NewServer(host)
		server.Listen()
	}()

	client := app.NewClient(peer, name)

	for {
		reader := bufio.NewReader(os.Stdin)
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %s\n", err)
		}
		client.SendMessage(content)
	}
}

func randomPort() int {
	return rand.Intn(100) + 40000
}
