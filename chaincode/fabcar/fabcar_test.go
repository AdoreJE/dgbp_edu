package main

import (
    "fmt"
    "testing"
    "encoding/json"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func checkInit(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
    res := stub.MockInit("1", args)

    if res.Status != shim.OK {
       fmt.Println("Init failed", string(res.Message))
       t.FailNow()
    }
}

func TestFabcar_Init(t *testing.T) {
    cc := new(SmartContract)
    stub := shimtest.NewMockStub("fabcar", cc)

    checkInit(t, stub, nil)
    checkQuery(t, stub, "CAR2")
    checkQueryAllCars(t, stub)
    checkCreateCar(t, stub)
    checkQuery(t, stub, "CAR11")
    checkChangeCarOwner(t, stub)
    checkQuery(t, stub, "CAR2")
}

func checkQuery(t *testing.T, stub *shimtest.MockStub, key string) {
    res := stub.MockInvoke("1", [][]byte{[]byte("QueryCar"), []byte(key)})
    
    if res.Payload == nil {
        fmt.Println("query failed")
        t.FailNow()
    }

    car := new(Car)
    json.Unmarshal(res.Payload, car)   
    fmt.Println(car)
}

func checkQueryAllCars(t *testing.T, stub *shimtest.MockStub) {
    res := stub.MockInvoke("1", [][]byte{[]byte("QueryAllCars")})

    if res.Payload == nil {
        fmt.Println("query failed")
        t.FailNow()
    }

    results := []QueryResult{}
    json.Unmarshal(res.Payload, &results)

    fmt.Println(results[0].Record)
}

func checkCreateCar(t *testing.T, stub *shimtest.MockStub) {

    res := stub.MockInvoke("1", [][]byte{[]byte("CreateCar"), []byte("CAR11"), []byte("AA"), []byte("AA"), []byte("AA"), []byte("AA")})

    if res.Status != shim.OK {
       fmt.Println("Init failed", string(res.Message))
       t.FailNow()
    }
}


func checkChangeCarOwner(t *testing.T, stub *shimtest.MockStub) {

    res := stub.MockInvoke("1", [][]byte{[]byte("ChangeCarOwner"), []byte("CAR2"), []byte("AA")})

    if res.Status != shim.OK {
       fmt.Println("Init failed", string(res.Message))
       t.FailNow()
    }
}

