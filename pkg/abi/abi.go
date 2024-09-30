package abi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var ContractABI abi.ABI

const abiJSON = `[
  {
    "anonymous": false,
    "inputs": [
      { "indexed": true, "name": "day", "type": "uint256" },
      { "indexed": false, "name": "tokenId", "type": "uint256" },
      { "indexed": false, "name": "author", "type": "address" },
      { "indexed": false, "name": "pixels", "type": "bytes" }
    ],
    "name": "Painted",
    "type": "event"
  }
]`

func init() {
	var err error
	ContractABI, err = abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic("Failed to parse ABI: " + err.Error())
	}
}

func GetContractABI() abi.ABI {
	return ContractABI
}
