package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/appengine"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("/home/antonpossylkine/.ethereum/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	race, err := NewBetting(common.HexToAddress("0xdbbe7e0db7c8819925843f73a03c94b495fbaa9a"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	chronus, err := race.Chronus(nil)
	fmt.Println("Starting time:", chronus.StartingTime)
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
