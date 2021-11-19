package main

import (
    "fmt"
    
    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

type PingPongChaincode struct {
}

func (t *PingPongChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Println("########## example_cc Init ##########")
    return shim.Success(nil)
}

// ping, pong
func (t *PingPongChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Println("################# example_cc Invoke ############")
    fnc, args := stub.GetFunctionAndParameters()  

    if fnc == "ping" {
       return t.ping(stub, args)
    } else if fnc == "pong" {
       return t.pong(stub, args)
    }

    return shim.Error("Unknown action, check thre first argument, must be one of 'ping', 'pong'")
}


// args : key, value
func (t *PingPongChaincode) ping(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 2 {
        return shim.Error("invalid arguments, expecting 2")
    }

    key := args[0]
    value := args[1]

    // stub.PutState(key string, value []byte)
    err := stub.PutState(key, []byte(value))
    if err != nil {
        return shim.Error(err.Error())
    }

    return shim.Success(nil)

}

// args : key 
func (t *PingPongChaincode) pong(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) != 1 {
        return shim.Error("invalid arguments, expecting 1")
    }

    key := args[0]

    valueAsBytes, err := stub.GetState(key)

    if err != nil {
      return shim.Error(err.Error())
    }

    if valueAsBytes == nil {
      return shim.Error("not found")
    }

    return shim.Success(valueAsBytes)   
}


func main() {
    err := shim.Start(new(PingPongChaincode))

    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
