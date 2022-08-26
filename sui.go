package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type SUIClient struct {
	ip     string
	client http.Client
}

// Constructor for SUI Client
func (sc *SUIClient) Init(ip string) {
	sc.ip = ip
	sc.client = http.Client{}
}

// Get total transactions count
func (sc *SUIClient) GetTotalTransactionNumber() uint64 {

	body := []byte(`{"jsonrpc":"2.0", "id":1, "method": "sui_getTotalTransactionNumber", "params": []}`)

	// Typical json response
	type Response struct {
		Result float64 `json:"result"`
	}

	// Placeholder for storing decoded json response
	var x Response

	// Creates new POST request with body
	req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err := sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err := io.ReadAll(res.Body)
	check(err)

	// Decodes bytearray
	err = json.Unmarshal(arr, &x)
	check(err)

	// Casting float64 result to uint32
	return uint64(x.Result)
}

// Get specific transaction
func (sc *SUIClient) GetTransaction(id string) (TX, error) {

	// Part 1: Get transaction and parse
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getTransaction", "params": ["%s"]}`, id))

	// Returned struct with proper formatting
	var x TX

	// Placeholder for temporary unmarshalling
	z := make(map[string]interface{})

	// Creates new POST request with body
	req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err := sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err := io.ReadAll(res.Body)
	check(err)

	/*
		Decodes entire transaction into placeholder
		This has to be done as golang cannot automatically decode unknown json field keys and structures
		We decode first then set fields
	*/
	err = json.Unmarshal(arr, &z)
	check(err)

	// Convert map to struct
	err = mapstructure.Decode(z, &x)
	check(err)

	if reflect.ValueOf(x.Result).IsZero() {
		return TX{}, errors.New("not valid transaction")
	}

	// Part 2: Get nicely formatted events

	// Placeholder to store our nicely formatted events
	s := []map[string]interface{}{}

	// Loop through all events
	for i := range x.Result.Effects.Events {
		tmp := map[string]interface{}{}

		// Set event transaction to current transaction
		tmp["tx"] = id
		for j, k := range x.Result.Effects.Events[i] {
			switch j {
			case "newObject":
				tmp["type"] = "mint"
				tmp["sender"] = k.(map[string]interface{})["sender"].(string)
				if reflect.TypeOf(k.(map[string]interface{})["recipient"]) == reflect.TypeOf(map[string]interface{}{}) {
					tmp["recipient"] = k.(map[string]interface{})["recipient"].(map[string]interface{})["AddressOwner"].(string)
				} else if reflect.TypeOf(k.(map[string]interface{})["recipient"]) == reflect.TypeOf("") {
					tmp["recipient"] = k.(map[string]interface{})["recipient"].(string)
				} else {
					tmp["recipient"] = nil
				}
				tmp["objectId"] = k.(map[string]interface{})["objectId"].(string)
				tmp["version"] = uint32(0)
			case "deleteObject":
				tmp["type"] = "burn"
				tmp["sender"] = k.(map[string]interface{})["sender"].(string)
				tmp["recipient"] = nil
				tmp["objectId"] = k.(map[string]interface{})["objectId"].(string)
				tmp["version"] = nil
			case "transferObject":
				tmp["type"] = "transfer"
				tmp["sender"] = k.(map[string]interface{})["sender"].(string)
				if reflect.TypeOf(k.(map[string]interface{})["recipient"]) == reflect.TypeOf(map[string]interface{}{}) {
					tmp["recipient"] = k.(map[string]interface{})["recipient"].(map[string]interface{})["AddressOwner"].(string)
				} else if reflect.TypeOf(k.(map[string]interface{})["recipient"]) == reflect.TypeOf("") {
					tmp["recipient"] = k.(map[string]interface{})["recipient"].(string)
				} else {
					tmp["recipient"] = nil
				}
				tmp["objectId"] = k.(map[string]interface{})["objectId"].(string)
				tmp["version"] = uint32(k.(map[string]interface{})["version"].(float64))
			}
		}
		s = append(s, tmp)
	}

	x.Events = &s

	// Part 3: Get func args
	/***********************************
	Only if transaction is a Call
	***********************************/

	if x.GetType() == "Call" {

		pkg, err := x.GetContractPackage()
		check(err)
		mod, err := x.GetContractModule()
		check(err)
		fn, err := x.GetContractFunction()
		check(err)

		body = []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getNormalizedMoveFunction", "params": ["%s", "%s", "%s"]}`, pkg, mod, fn))

		// Placeholder for temporary unmarshalling
		z := make(map[string]interface{})

		// Creates new POST request with body
		req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
		check(err)
		req.Header.Set("Content-Type", "application/json")

		// Dispatches request
		res, err := sc.client.Do(req)
		check(err)
		defer res.Body.Close()

		// Converting Response body to byte array
		arr, err := io.ReadAll(res.Body)
		check(err)

		/*
			Decodes entire transaction into placeholder
			This has to be done as golang cannot automatically decode unknown json field keys and structures
			We decode first then set fields
		*/
		err = json.Unmarshal(arr, &z)
		check(err)

		tmp := []map[string]interface{}{}

		// Get raw arguments data from transaction for later indexing use
		raw, err := x.GetRawContractArguments()
		check(err)

		// Loop through all the parameters that the function call takes
		for i, v := range z["result"].(map[string]interface{})["parameters"].([]interface{}) {

			// Skip last item in parameters list (gas: tx_context)
			if i == len(raw.([]interface{})) {
				break
			}
			tmp = append(tmp, map[string]interface{}{})

			// Create placeholder structure for arguments
			tmp[i]["id"] = id

			// If the parameter is an object, recursively get the value
			if reflect.TypeOf(v) == reflect.TypeOf(map[string]interface{}{}) {

				// If parameter object is a reference
				if recurseKey(v.(map[string]interface{}), "name") != nil && recurseKey(v.(map[string]interface{}), "address") != nil && recurseKey(v.(map[string]interface{}), "module") != nil {
					tmp[i]["name"] = recurseKey(v.(map[string]interface{}), "name").(string)
					tmp[i]["type"] = recurseKey(v.(map[string]interface{}), "address").(string) +
						"::" +
						recurseKey(v.(map[string]interface{}), "module").(string) +
						"::" +
						recurseKey(v.(map[string]interface{}), "name").(string)

					// If parameter object is something else (vector)
				} else {
					for k, j := range v.(map[string]interface{}) {
						tmp[i]["name"] = nil
						tmp[i]["type"] = k + "<" + j.(string) + ">"
					}
				}

				// Just a string with type inside
			} else {
				tmp[i]["name"] = nil
				tmp[i]["type"] = v.(string)
			}
			tmp[i]["data"] = raw.([]interface{})[i]
		}

		x.Arguments = &tmp
	}

	return x, nil
}

// Get transactions in specific range
func (sc *SUIClient) GetTransactionsInRange(start uint64, end uint64) ([]string, error) {

	if start > end {
		return []string{}, errors.New("start must not exceed end")
	} else if end-start > 4096 {
		return []string{}, errors.New("maximum 4096 transactions allowed")
	}

	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getTransactionsInRange", "params": [%d, %d]}`, start, end))

	type Response struct {
		Result []interface{} `json:"result"`
	}

	// Placeholder for storing decoded json response
	var x Response

	// Placeholder for returned array
	var list []string

	// Creates new POST request with body
	req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err := sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err := io.ReadAll(res.Body)
	check(err)

	// Decodes bytearray
	err = json.Unmarshal(arr, &x)
	check(err)

	// Look for the string value in the index, id array and appends it to the return array
	for _, i := range x.Result {
		for _, value := range i.([]interface{}) {
			switch value := value.(type) {
			case string:
				list = append(list, value)
			}
		}
	}

	return list, nil
}

// Get object
func (sc *SUIClient) GetObject(id string) (Obj, error) {
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getObject", "params": ["%s"]}`, id))

	// Returned struct with proper formatting
	var x Obj

	// Placeholder for temporary unmarshalling
	z := make(map[string]interface{})

	// Creates new POST request with body
	req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err := sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err := io.ReadAll(res.Body)
	check(err)

	/*
		Decodes entire object into placeholder
		This has to be done as golang cannot automatically decode unknown json field keys and structures
		We decode first then set fields
	*/
	err = json.Unmarshal(arr, &z)
	check(err)

	// Convert map to struct
	err = mapstructure.Decode(z, &x)
	check(err)

	if reflect.ValueOf(x.Result).IsZero() {
		return Obj{}, errors.New("not valid object")
	}

	return x, nil
}

// Get transactions related to object
func (sc *SUIClient) GetTransactionsByObject(id string) ([]string, error) {
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getTransactionsByMutatedObject", "params": ["%s"]}`, id))

	type Response struct {
		Result []interface{} `json:"result"`
	}

	// Placeholder for storing decoded json response
	var x Response

	// Placeholder for returned array
	var list []string

	// Creates new POST request with body
	req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err := sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err := io.ReadAll(res.Body)
	check(err)

	// Decodes bytearray
	err = json.Unmarshal(arr, &x)
	check(err)

	if reflect.ValueOf(x.Result).IsZero() {
		return []string{}, errors.New("not valid object")
	}

	// Look for the string value in the index, id array and appends it to the return array
	for _, i := range x.Result {
		for _, value := range i.([]interface{}) {
			switch value := value.(type) {
			case string:
				list = append(list, value)
			}
		}
	}

	return list, nil
}

// Get an account
func (sc *SUIClient) GetAccount(id string) (Acc, error) {

	// Account object to be returned at the end
	var x Acc

	// Transactions format that we receive
	type Response struct {
		Result []interface{} `json:"result"`
	}

	// To house all the transactions
	var list []string

	// Unmarshalling json request in part 1
	var y AccResponse

	// Unmarshalling in part 2
	var y2 Response

	// Ummarshalling in part 3
	var y3 Response

	/***********************************************
	Part 1: Get objects
	************************************************/
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getObjectsOwnedByAddress", "params": ["%s"]}`, id))

	// Creates new POST request with body
	req, err := http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err := sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err := io.ReadAll(res.Body)
	check(err)

	/*
		Decodes entire object into placeholder
		This has to be done as golang cannot automatically decode unknown json field keys and structures
		We decode first then set fields
	*/
	err = json.Unmarshal(arr, &y)
	check(err)

	if reflect.ValueOf(y.Result).IsZero() {
		return Acc{}, errors.New("not valid account")
	}

	// Set objects of returned account to be the response from request
	x.Objects = []struct {
		ObjectId string
		Type     string
	}(y.Result)

	// Set balance of account
	for _, v := range y.Result {

		// Check if object is of SUI Type
		if v.Type == "0x2::coin::Coin<0x2::sui::SUI>" {

			// Grab a copy of the object
			tmp, _ := sc.GetObject(v.ObjectId)

			// Grab its metadata
			mtdt := tmp.GetObjectMetadata()
			bal := mtdt["balance"]

			// Check if object balance is of float64 type
			if reflect.TypeOf(bal) == reflect.TypeOf(1.0) {
				x.Balance += uint64(bal.(float64))
			}
		}
	}

	// Set account ID
	x.ID = id

	/***********************************************
	Part 2: Transactions to account
	************************************************/
	body = []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getTransactionsToAddress", "params": ["%s"]}`, id))

	// Creates new POST request with body
	req, err = http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err = sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err = io.ReadAll(res.Body)
	check(err)

	// Decodes bytearray
	err = json.Unmarshal(arr, &y2)
	check(err)

	// Look for the string value in the index, id array and appends it to the return array
	for _, i := range y2.Result {
		for _, value := range i.([]interface{}) {
			switch value := value.(type) {
			case string:
				list = append(list, value)
			}
		}
	}

	/***********************************************
	Part 3: Transactions from account
	************************************************/
	body = []byte(fmt.Sprintf(`{"jsonrpc":"2.0", "id":1, "method": "sui_getTransactionsFromAddress", "params": ["%s"]}`, id))

	// Creates new POST request with body
	req, err = http.NewRequest(http.MethodPost, sc.ip, bytes.NewBuffer(body))
	check(err)
	req.Header.Set("Content-Type", "application/json")

	// Dispatches request
	res, err = sc.client.Do(req)
	check(err)
	defer res.Body.Close()

	// Converting Response body to byte array
	arr, err = io.ReadAll(res.Body)
	check(err)

	// Decodes bytearray
	err = json.Unmarshal(arr, &y3)
	check(err)

	// Look for the string value in the index, id array and appends it to the return array
	for _, i := range y3.Result {
		for _, value := range i.([]interface{}) {
			switch value := value.(type) {
			case string:
				list = append(list, value)
			}
		}
	}

	// Set the transactions in account to be the populated list
	x.Transactions = list
	return x, nil
}
