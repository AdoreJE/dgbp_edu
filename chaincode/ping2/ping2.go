package main

import (
    "fmt"
    "encoding/json" 
    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

type PingPongChaincode struct {
}

type PingST struct {
   Key string 
   Value string 
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

    p := PingST{
       Key: args[0],
       Value: args[1],
    }

    pAsBytes, _ := json.Marshal(p)

    // stub.PutState(key string, value []byte)
    err := stub.PutState(p.Key, pAsBytes)
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
    fmt.Println("key: ", key)
    
    pAsBytes, err := stub.GetState(key)

    fmt.Println("pAsBytes: ", pAsBytes)
    if err != nil {
      return shim.Error(err.Error())
    }

    if pAsBytes == nil {
      return shim.Error("not found")
    }

    p := new(PingST) 
    _ = json.Unmarshal(pAsBytes, p)
    fmt.Println("p : ", p)

    valueAsBytes := []byte(p.Value)
    return shim.Success(valueAsBytes)   
}


func main() {
    err := shim.Start(new(PingPongChaincode))

    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
