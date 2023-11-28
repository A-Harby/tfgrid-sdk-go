package main

import (
	"context"
	"fmt"
	"time"

	"log"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/farmerbot/models"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
)

func findNode() (uint32, error) {
	mnemonics := "<mnemonics goes here>"
	subManager := substrate.NewManager("wss://tfchain.dev.grid.tf/ws")
	sub, err := subManager.Substrate()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to substrate: %w", err)
	}
	defer sub.Close()

	client, err := peer.NewRpcClient(
		context.Background(),
		peer.KeyTypeSr25519,
		mnemonics,
		"wss://relay.dev.grid.tf",
		"test-find-node",
		sub,
		true,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create rpc client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	farmID := 53 // <- replace this with the farm id of the farmerbot
	service := fmt.Sprintf("farmerbot-%d", farmID)
	const farmerbotTwinID = 164 // <- replace this with the twin id of where the farmerbot is running

	options := models.NodeOptions{
		NodeExclude:  []uint32{},
		HasGPUs:      0,
		GPUVendors:   []string{},
		GPUDevices:   []string{},
		Certified:    false,
		Dedicated:    false,
		PublicConfig: false,
		PublicIPs:    0,
		Capacity:     models.Capacity{},
	}
	var output uint32
	if err := client.Call(ctx, farmerbotTwinID, &service, "nodemanager.findnode", options, &output); err != nil {
		return 0, err
	}

	return 0, nil
}

func main() {
	nodeID, err := findNode()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("nodeID: %v\n", nodeID)
}
