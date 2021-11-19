package main

import (
    "fmt"
    "testing"

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

func TestPing_Init(t *testing.T) {
    cc := new(PingPongChaincode)
    stub := shimtest.NewMockStub("ping", cc)

    checkInit(t, stub, nil)
}

func checkInvoke(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
    res := stub.MockInvoke("1", args)

    if res.Status != shim.OK {
      fmt.Println("Invoke", args, "failed", string (res.Message))
      t.FailNow()
    }
}

func TestPing_Invoke(t *testing.T) {
    cc := new(PingPongChaincode)
    stub := shimtest.NewMockStub("ping", cc)
    
//    a := []int{1,2,3}
   
  //  a := [][]byte{[]byte("ping"), []byte("a"), []byte("10")}

    checkInvoke(t, stub, [][]byte{[]byte("ping"), []byte("a"), []byte("10")})

}

func checkQuery(t *testing.T, stub *shimtest.MockStub, key string, value string) {
    res := stub.MockInvoke("1", [][]byte{[]byte("pong"), []byte(key)})
    if res.Status != shim.OK {
        fmt.Println("query pong api is failed", string(res.Message))
        t.FailNow()
    }

    if res.Payload == nil {
        fmt.Println("Query pong api failed to get value")
        t.FailNow()
    }

    if string(res.Payload) != value {
       fmt.Println("Query value", key, "was not", value, "as expected")
       t.FailNow()
    }
}

func TestPing_Query(t *testing.T) {
    cc := new(PingPongChaincode)
    stub := shimtest.NewMockStub("ping", cc)
    checkInvoke(t, stub, [][]byte{[]byte("ping"), []byte("a"), []byte("20")})
    checkQuery(t, stub, "a", "10")
}

func TestPing_InvokeWithIncorrectArugments(t *testing.T) {
    cc := new(PingPongChaincode)
    stub := shimtest.NewMockStub("ping", cc)

    res := stub.MockInvoke("1", [][]byte{[]byte("ping"), []byte("a"), []byte("10"), []byte("1")})

    if res.Status != shim.ERROR {
       fmt.Println("invalid invoke")
       t.FailNow()
    }

    if res.Message != "invalid arguments, expecting 2" {
       fmt.Println("Unexpected Error message:", string(res.Message))
       t.FailNow()
    }
}

func TestPing_QueryWithIncorrectArguments(t *testing.T) {
cc := new(PingPongChaincode)
    stub := shimtest.NewMockStub("ping", cc)

    checkInvoke(t, stub, [][]byte{[]byte("ping"), []byte("a"), []byte("10")})

    res := stub.MockInvoke("1", [][]byte{[]byte("pong"), []byte("a"), []byte("10"), []byte("1")})

    if res.Status != shim.ERROR {
       fmt.Println("invalid invoke")
       t.FailNow()
    }

    if res.Message != "invalid arguments, expecting 1" {
       fmt.Println("Unexpected Error message:", string(res.Message))
       t.FailNow()
    }

    res2 := stub.MockInvoke("1", [][]byte{[]byte("pong"), []byte("b")})

    if res2.Status != shim.ERROR {
       fmt.Println("invalid query")
       t.FailNow()
    }

    if res2.Message != "not found" {
       fmt.Println("invalid query")
       t.FailNow()
    }
}










