package sui

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Transaction Structure
type TX struct {
	Result struct {
		Certificate struct {
			TransactionDigest string `json:"transactionDigest"`

			Data struct {
				Sender string `json:"sender"`
				// Need keys to iterate
				Transactions []map[string]interface{} `json:"transactions"`
				GasPayment   struct {
					Digest   string `json:"digest"`
					ObjectId string `json:"objectId"`
				} `json:"gasPayment"`
			} `json:"data"`

			TxSignature string `json:"txSignature"`

			AuthSignInfo struct {
				PeerSignature []string `json:"signature"`
			} `json:"authSignInfo"`
		} `json:"Certificate"`

		Effects struct {
			Status struct {
				Status string `json:"status"`
			} `json:"status"`

			GasObject struct {
				Owner     interface{} `json:"owner"`
				Reference struct {
					ObjectId string `json:"objectId"`
				} `json:"reference"`
			} `json:"gasObject"`

			GasUsed struct {
				ComputationCost float64 `json:"computationCost"`
				StorageCost     float64 `json:"storageCost"`
				StorageRebate   float64 `json:"storageRebate"`
			} `json:"gasUsed"`

			Mutated []struct {
				Owner     interface{} `json:"owner"`
				Reference struct {
					ObjectId string `json:"objectId"`
				} `json:"reference"`
			} `json:"mutated"`

			Created []struct {
				Owner     interface{} `json:"owner"`
				Reference struct {
					ObjectId string `json:"objectId"`
				} `json:"reference"`
			} `json:"created"`

			Deleted []struct {
				ObjectId string `json:"objectId"`
			} `json:"deleted"`

			SharedObjects []struct {
				ObjectId string `json:"objectId"`
			} `json:"sharedObjects"`

			Wrapped []struct {
				ObjectId string `json:"objectId"`
			} `json:"wrapped"`

			Unwrapped []struct {
				Owner     interface{} `json:"owner"`
				Reference struct {
					ObjectId string `json:"objectId"`
				} `json:"reference"`
			} `json:"unwrapped"`
			// Need keys to iterate
			Events []map[string]interface{} `json:"events"`
		} `json:"effects"`
		Timestamp_ms float64 `json:"timestamp_ms"`
	} `json:"result"`
	Arguments *[]Arg
	Events    *[]Event
}

// Object Structure
type Obj struct {
	Result struct {
		Status  string `json:"status"`
		Details struct {
			Data struct {
				DataType            string `json:"dataType"`
				Type                string `json:"type"`
				Has_public_transfer bool   `json:"has_public_transfer"`
				// Need keys to iterate
				Fields map[string]interface{} `json:"fields"`
			} `json:"data"`
			// Can be pure string, therefore must use interface{} then cast to map[string]interface{} if type is not string
			Owner         interface{} `json:"owner"`
			StorageRebate float64     `json:"storageRebate"`
			Reference     struct {
				ObjectId string `json:"objectId"`
			} `json:"reference"`
		} `json:"details"`
	} `json:"result"`
}

// Response structure from calling sui_getObjectsOwnedByAddress
type AccResponse struct {
	Result []struct {
		ObjectId string                 `json:"objectId"`
		Type     string                 `json:"type"`
		Metadata map[string]interface{} `json:"-"`
	} `json:"result"`
}

// Account Structure
type Acc struct {
	ID           string
	Balance      uint64
	Objects      []AccObject
	Transactions []string
}

type Arg struct {
	Name string
	Type string
	ID   string
	Data string
}

type Event struct {
	Type      string
	Sender    string
	Recipient *string
	TX        string
	ObjectId  string
	Version   uint32
}

type AccObject struct {
	ObjectId string
	Type     string
	Metadata map[string]interface{}
}

type Package struct {
	DeployTX string
	ID       string
	Bytecode map[string]interface{}
}

// Get the type of a transaction (e.g. Call, TransferSui, TransferObject)
func (tx *TX) GetType() string {
	var result string

	/*
		The actual transaction data is within the first index of the transactions array within response body (the length, however, is fixed at 1)
		Since we do not know the key of this data, we need to iterate through (only one iteration needed) to retrieve the key field
	*/
	for k := range tx.Result.Certificate.Data.Transactions[0] {
		result = k
	}
	return result
}

// Get the time in ms epoch of a transaction
func (tx *TX) GetTime() time.Time {
	// integ, decim := math.Modf(tx.Result.Timestamp_ms / 1000)
	// return time.Unix(int64(integ), int64(decim*(1e9))).UTC()
	fmt.Println(time.Now())
	return time.Now()
}

// Get the digest of a transaction
func (tx *TX) GetID() string {
	return tx.Result.Certificate.TransactionDigest
}

// Get the status of a transaction (success: true, fail: false)
func (tx *TX) GetStatus() bool {
	switch tx.Result.Effects.Status.Status {
	case "success":
		return true
	default:
		return false
	}
}

// Get the sender of a transaction
func (tx *TX) GetSender() string {
	return tx.Result.Certificate.Data.Sender
}

// Get the recipient of a transaction, returns error if there is no recipient
func (tx *TX) GetRecipient() *string {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()
	if rec := tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["recipient"]; rec != nil {
		temp := rec.(string)
		return &temp
	} else {
		return nil
	}
}

// Get the amount transferred in a transaction, strictly referring to SUI
func (tx *TX) GetTransferAmount() *float64 {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()

	if amt := tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["amount"]; amt != nil {
		temp := amt.(float64)
		return &temp
	} else {
		return nil
	}
}

// Get gas cost
func (tx *TX) GetGas() uint32 {
	return uint32(tx.Result.Effects.GasUsed.StorageCost) + uint32(tx.Result.Effects.GasUsed.ComputationCost)

}

// Get the package of a Call transaction
func (tx *TX) GetContractPackage() *string {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()
	if tp != "Call" {
		return nil
	}
	if id := tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["package"].(map[string]interface{})["objectId"]; id != nil {
		temp := id.(string)
		return &temp
	} else {
		return nil
	}
}

// Get the module of a Call transaction
func (tx *TX) GetContractModule() *string {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()
	if tp != "Call" {
		return nil
	}
	if mod := tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["module"]; mod != nil {
		temp := mod.(string)
		return &temp
	} else {
		return nil
	}
}

// Get the function of a Call transaction
func (tx *TX) GetContractFunction() *string {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()
	if tp != "Call" {
		return nil
	}
	if fn := tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["function"]; fn != nil {
		temp := fn.(string)
		return &temp
	} else {
		return nil
	}
}

// Get data on a package deploy
func (tx *TX) GetContractDeploy() (Package, error) {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()

	// If transaction is not publish throw error
	if tp != "Publish" {
		return Package{}, errors.New("no package publish")
	}

	// Create returned map string interface object
	result := Package{}

	// Set the deployment contract
	result.DeployTX = tx.GetID()

	// Find the deployed package id
	for _, v := range tx.Result.Effects.Events {

		// Loop through each key of each event (there's only one key)
		for k, n := range v {
			if k == "publish" {
				result.ID = n.(map[string]interface{})["packageId"].(string)
			}
		}
	}

	// Set the bytecode to be the bytecode from transaction
	result.Bytecode = tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["disassembled"].(map[string]interface{})

	// If any of the fields isn't filled throw error
	if result.ID == "" || result.Bytecode == nil {
		return Package{}, errors.New("no package publish")
	}

	return result, nil
}

// Get the arguments of a Call transaction
func (tx *TX) GetRawContractArguments() (interface{}, error) {

	// First get the transaction type so we can go straight into transaction data
	tp := tx.GetType()
	if tp != "Call" {
		return nil, errors.New("no contract call")
	}
	if fn := tx.Result.Certificate.Data.Transactions[0][tp].(map[string]interface{})["arguments"]; fn != nil {
		return fn, nil
	} else {
		return nil, nil
	}
}

// Get the datatype of an object (only moveObject so far...?)
func (obj *Obj) GetObjectDataType() string {
	return obj.Result.Details.Data.DataType
}

// Get the type of an object
func (obj *Obj) GetObjectType() string {
	return obj.Result.Details.Data.Type
}

// Get the object id of an object
func (obj *Obj) GetObjectID() string {
	return obj.Result.Details.Reference.ObjectId
}

// Get the object status of an object
func (obj *Obj) GetObjectStatus() string {
	return obj.Result.Status
}

// Get the package that the object belongs to
func (obj *Obj) GetObjectPackage() string {

	// Split the objec type and get the first element
	return strings.Split(obj.Result.Details.Data.Type, "::")[0]
}

// Get the module that the object belongs to
func (obj *Obj) GetObjectModule() string {

	// Split the objec type and get the first and second elements
	return strings.Split(obj.Result.Details.Data.Type, "::")[0] + "::" + strings.Split(obj.Result.Details.Data.Type, "::")[1]
}

// Get whether the object can be publicly transferred (not sure what this means)
func (obj *Obj) HasPublicTransfer() bool {
	return obj.Result.Details.Data.Has_public_transfer
}

// Get the owner of an object
func (obj *Obj) GetOwner() string {

	// First check if owner field contains an object or string (string usually occurs when the object is shared)
	if reflect.TypeOf(obj.Result.Details.Owner) == reflect.TypeOf(map[string]interface{}{}) {

		// If it's an object, iterate through its key (only one key)
		for k, v := range obj.Result.Details.Owner.(map[string]interface{}) {

			// Just in case
			if k == "AddressOwner" {
				return v.(string)
			}
		}

		// Else if owner field is a string
	} else if reflect.TypeOf(obj.Result.Details.Owner) == reflect.TypeOf("") {
		return obj.Result.Details.Owner.(string)
	}
	return ""
}

// Get the metadatas of an object
func (obj *Obj) GetObjectMetadata() map[string]interface{} {
	return obj.Result.Details.Data.Fields
}

// Get NFTs owned by account
func (acc *Acc) GetAccountNFTs() []AccObject {

	// Placeholder empty map string interface slice
	result := []AccObject{}

	// Iterate through all account objects
	for _, v := range acc.Objects {
		if v.Type != "0x2::coin::Coin<0x2::sui::SUI>" {
			obj := AccObject{
				ObjectId: v.ObjectId,
				Type:     v.Type,
				Metadata: v.Metadata,
			}
			result = append(result, obj)
		}
	}
	if len(result) == 0 {
		return nil
	} else {
		return result
	}
}
