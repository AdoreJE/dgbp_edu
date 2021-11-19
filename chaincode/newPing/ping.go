package main

import (
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PingPongChaincode struct {
    contractapi.Contract
}

func (p *PingPongChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
    fmt.Println("############ initLedger ############")
    return nil
}

func (p *PingPongChaincode) Ping(ctx contractapi.TransactionContextInterface, key string, value string) error {
   return ctx.GetStub().PutState(key, []byte(value))
}

func (p *PingPongChaincode) Pong(ctx contractapi.TransactionContextInterface, key string) (string, error) {
    valueAsBytes, err := ctx.GetStub().GetState(key)
   
    if err != nil {
        return "error", fmt.Errorf("Fail to read from world state. %s", err.Error())
    }

    return string(valueAsBytes[:]), nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(PingPongChaincode))

    if err != nil {
        fmt.Printf("Error create pingpong chaincode: %s", err.Error())
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting pingpong chaincode: %s", err.Error())
    }
}
