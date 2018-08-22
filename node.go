package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// Node represents a local node
type Node struct {
	Conn *ethclient.Client
}

// NewNode a node constructor
func NewNode(url string) (*Node, error) {
	node := new(Node)
	var err error
	// Create an IPC based RPC connection to a remote node
	node.Conn, err = ethclient.Dial(url)

	return node, err
}
