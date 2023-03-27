package main

import (
	"fmt"
	"math/big"

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

	Account := gjson.Get(res, "Account")

	log.Log("wasm get Account from json: " + Account.String())

	// TODO how to read state from chain?
	blockchain.SendTx(
		4690, // chain id
		"0x8AA6C8023EEe0Fc89a4e7eBfcAE7F21876311410", // contract address
		big.NewInt(0),
		fmt.Sprintf("6a627842000000000000000000000000%s", Account.String()),
	)
	log.Log("Vote SBT has been sent to :" + Account.String())

	return 0
}
