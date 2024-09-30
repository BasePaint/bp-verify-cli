package events

import (
	"context"
	"fmt"
	"math/big"

	"github.com/BasePaint/bpverify/pkg/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const startingBlock = 2385188
const basepaintAddress = "0xba5e05cb26b78eda3a2f8e3b3814726305dcac83"

func GetEvents(rpcUrl string, day int) ([][]byte, error) {
    client, err := ethclient.Dial(rpcUrl)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to the client: %v", err)
    }

    eventSignature := []byte("Painted(uint256,uint256,address,bytes)")
    eventSignatureHash := crypto.Keccak256Hash(eventSignature)

    dayBytes := common.LeftPadBytes(big.NewInt(int64(day)).Bytes(), 32)
    dayHash := common.BytesToHash(dayBytes)

    address := common.HexToAddress(basepaintAddress)
    fromBlock := big.NewInt(startingBlock) // Starting block

    query := ethereum.FilterQuery{
        FromBlock: fromBlock,
        Addresses: []common.Address{address},
        Topics: [][]common.Hash{
            {eventSignatureHash},
            {dayHash},
        },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        return nil, fmt.Errorf("failed to filter logs: %v", err)
    }

    var allPixels [][]byte

    for _, vLog := range logs {
        pixels, err := processLog(vLog)
        if err != nil {
            fmt.Printf("Error processing log: %v\n", err)
            continue
        }
        allPixels = append(allPixels, pixels)
    }

    fmt.Printf("Retrieved %d paint events for BasePaint Day #%v\n", len(allPixels), day)
    return allPixels, nil
}

func processLog(vLog types.Log) ([]byte, error) {
    paintEvent := struct{
        Day *big.Int
        TokenId *big.Int
        Author common.Address
        Pixels  []byte
    }{}
    err := abi.GetContractABI().UnpackIntoInterface(&paintEvent, "Painted", vLog.Data)
    if err != nil {
        return nil, fmt.Errorf("failed to process log: %v", err)
    }
    return paintEvent.Pixels, nil
}