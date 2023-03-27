package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/blockchain"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

func main() {}

//export start
func _start(rid uint32) int32 {
	log.Log(fmt.Sprintf("start received: %d", rid))
	message, err := stream.GetDataByRID(rid)
	if err != nil {
		log.Log("error: " + err.Error())
		return -1
	}
	res := string(message)

	account := gjson.Get(res, "account")

	network, err := stream.GetEnv("NETWORK")
	if err != nil {
		log.Log(fmt.Sprintf("get network from host failed: %v", err))
		return -1
	}
	chainId, err := strconv.ParseUint(network, 10, 32)
	if err != nil {
		log.Log(fmt.Sprintf("convert network to chainId failed: %v", err))
		return -1
	}
	contract, err := stream.GetEnv("CONTRACT")
	if err != nil {
		log.Log(fmt.Sprintf("get contract from host failed: %v", err))
		return -1
	}
	log.Log("Sending Vote SBT(" + contract + ":" + network + ") to " + account.String())

	// TODO how to read state from chain?
	blockchain.SendTx(
		uint32(chainId), // chain id
		contract,        // contract address
		big.NewInt(0),
		fmt.Sprintf("6a627842000000000000000000000000%s", account.String()),
	)
	log.Log("Vote SBT has been sent")

	return 0
}
