package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Person struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Gender  string `json:"gender"`
	Balance int    `json:"balance"`
	TimeIn  string `json:"timein"`
	TimeOut string `json:"timeout"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	args := APIstub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	// Set up any variables or assets here by calling stub.PutState()

	// We store the key and the value on the ledger
	err := APIstub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryUser" {
		return s.queryUser(APIstub, args)
	} else if function == "createUser" {
		return s.createUser(APIstub, args)
	} else if function == "queryAllUsers" {
		return s.queryAllUsers(APIstub)
	} else if function == "In" {
		return s.In(APIstub, args)
	} else if function == "Out" {
		return s.Out(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	users := []Person{
		Person{Name: "Thao", Id: "1", Gender: "male", Balance: 999, TimeIn: "123", TimeOut: "234"},
	}

	i := 0
	for i < len(users) {
		fmt.Println("i is ", i)
		userAsBytes, _ := json.Marshal(users[i])
		APIstub.PutState("User"+strconv.Itoa(i), userAsBytes)
		fmt.Println("Added", users[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(userAsBytes)
}

func (s *SmartContract) createUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	balance, err := strconv.Atoi(args[3])
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	var user = Person{Name: args[0], Id: args[1], Gender: args[2], Balance: balance, TimeIn: args[4], TimeOut: args[5]}

	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllUsers(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "USER1"
	endKey := "USER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllUsers:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) In(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	user := Person{}

	json.Unmarshal(userAsBytes, &user)
	user.TimeIn = args[1]

	userAsBytes, _ = json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) Out(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	user := Person{}

	json.Unmarshal(userAsBytes, &user)
	user.TimeOut = args[1]
	timeout, err1 := strconv.Atoi(args[1])
	if err1 != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	timein, err2 := strconv.Atoi(user.TimeIn)
	if err2 != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	user.Balance = user.Balance + timeout - timein
	userAsBytes, _ = json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
