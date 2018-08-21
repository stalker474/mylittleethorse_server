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
	conn, err := ethclient.Dial(ipc)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	race, err := NewBettingCaller(common.HexToAddress("0xd385915904b2bd6502b244195c802c7b90aa261b"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	reward, err := race.RewardTotal(nil)
	if err != nil {
		log.Fatalf("Failed to get reward: %v", err)
	}
	fmt.Println("Reward total:", reward)
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {

	races := fetchArchive()

	fmt.Fprintln(w, races[0].ID)
}
