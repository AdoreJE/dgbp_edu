/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

//	"github.com/hyperledger/fabric-contract-api-go/contractapi"
        "github.com/hyperledger/fabric-chaincode-go/shim"
        pb "github.com/hyperledger/fabric-protos-go/peer"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
}

// Car describes basic details of what makes up a car
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

// Init adds a base set of cars to the ledger
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := stub.PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return shim.Error("Failed to put to world state.")
		}
	}

	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fnc, args := stub.GetFunctionAndParameters()

    if fnc == "CreateCar" {
        return s.CreateCar(stub, args)
    } else if fnc == "QueryCar" {
        return s.QueryCar(stub, args)
    } else if fnc == "QueryAllCars" {
        return s.QueryAllCars(stub)
    } else if fnc == "ChangeCarOwner" {
        return s.ChangeCarOwner(stub, args)
    }

    return shim.Error("Unknown action")
}
// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateCar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	car := Car{
		Make:   args[1],
		Model:  args[2],
		Colour: args[3],
		Owner:  args[4],
	}

	carAsBytes, _ := json.Marshal(car)

	stub.PutState(args[0], carAsBytes)
        return shim.Success(nil)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        carNumber := args[0]
	carAsBytes, err := stub.GetState(carNumber)

	if err != nil {
		return shim.Error("Failed to read from world state.")
	}

	if carAsBytes == nil {
		return shim.Error("does not exist")
	}

//	car := new(Car)
//	_ = json.Unmarshal(carAsBytes, car)

	return shim.Success(carAsBytes)
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllCars(stub shim.ChaincodeStubInterface) pb.Response {
	startKey := ""
	endKey := ""

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)

	if err != nil {
                return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
                        return shim.Error(err.Error())
		}

		car := new(Car)
		_ = json.Unmarshal(queryResponse.Value, car)

		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
		results = append(results, queryResult)
	}

        resultsAsBytes, _ := json.Marshal(results)
        return shim.Success(resultsAsBytes)
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) ChangeCarOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        carNumber := args[0]
        newOwner := args[1]

//        newArgs := []string{carNumber}
	
        carAsBytes, err := stub.GetState(carNumber)
	//car, err := s.QueryCar(stub, newArgs)

	if err != nil {
		return shim.Error(err.Error())
	}
        car := new(Car)
        json.Unmarshal(carAsBytes, car)

	car.Owner = newOwner

	carAsBytes, _ = json.Marshal(car)
	stub.PutState(carNumber, carAsBytes)

	return shim.Success(nil)
}

func main() {

        err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}
}
