package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/machinefi/w3bstream-wasm-golang-sdk/blockchain"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/log"
	"github.com/machinefi/w3bstream-wasm-golang-sdk/stream"
)

var ZERO = big.NewInt(0)

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

func ReadBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

func PaddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	ReadBits(bigint, ret)
	return ret
}

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
	tokenId := gjson.Get(res, "tokenId")

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

	tokenIdBig, ok := new(big.Int).SetString(tokenId.String(), 10)
	if !ok {
		log.Log(fmt.Sprintf("convert tokenId to BigInt failed: %s", tokenId.String()))
		return -1
	}
	tokenIdHex := hex.EncodeToString(PaddedBigBytes(tokenIdBig, 32))

	data, err := blockchain.CallContract(
		uint32(chainId), // chain id
		contract,        // contract address
		fmt.Sprintf("00fdd58e000000000000000000000000%s%s", account.String(), tokenIdHex),
	)
	if err != nil {
		log.Log(fmt.Sprintf("Read balanceOf failed: %v", err))
		return -1
	}
	balance := new(big.Int).SetBytes(data)

	if balance.Cmp(ZERO) == 0 {
		log.Log("Sending Vote SBT(" + contract + ":" + network + ") to " + account.String())
		blockchain.SendTx(
			uint32(chainId), // chain id
			contract,        // contract address
			big.NewInt(0),
			fmt.Sprintf("40c10f19000000000000000000000000%s%s", account.String(), tokenIdHex),
		)
		log.Log("Vote SBT has been sent")
	} else {
		log.Log(fmt.Sprintf("%s already have Vote SBT", account.String()))
	}

	return 0
}
