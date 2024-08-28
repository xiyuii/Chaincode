package chaincode

import (
    "bytes"
    "encoding/json"
    "fmt"
	"log"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, wallet string) error {
	user := User{
		UserID:     userID,
		UserType:   userType,
		Wallet:     wallet,
		Models:     []*Model{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(userID, userAsBytes)
	
}

func (s *SmartContract) RegisterModel(ctx contractapi.TransactionContextInterface, modelID string, modelName string, api string, version string, owner string, submissionTime string, callID string, description string) error {
	model := Model{
		ModelName:     modelName,  
		API:           api,
		ModelID:       modelID,
		Version:       version,
		Owner:         owner, 
		SubmissionTime: submissionTime,
		Description:   description,
	}
	modelAsBytes, err := json.Marshal(model)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(modelID, modelAsBytes)
}


// 调用模型
func (s *SmartContract) InvokeModel(ctx contractapi.TransactionContextInterface, modelID string, userID string, inputData string, outputData string) (string, error) {
    // 获取模型信息
    model, err := s.GetModel(ctx, modelID)
    if err != nil {
        return "", err
    }

    // 获取用户信息
    user, err := s.GetUser(ctx, userID)
    if err != nil {
        return "", err
    }


    // 记录调用结果
    err = s.RecordInvocation(ctx, modelID, userID, inputData, outputData)
    if err != nil {
        return "", err
    }

    // 返回调用结果
    return outputData, nil
}

// 记录调用结果
func (s *SmartContract) RecordInvocation(ctx contractapi.TransactionContextInterface, modelID string, userID string, inputData string, outputData string) error {
    // 获取时间戳
    txtime, err := ctx.GetStub().GetTxTimestamp()
    if err != nil {
        return err
    }
    timeLocation, _ := time.LoadLocation("Asia/Shanghai")
    timestamp := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

    // 创建调用记录
    invocation := map[string]string{
        "modelID":    modelID,
        "userID":     userID,
        "inputData":  inputData,
        "outputData": outputData,
        "timestamp":  timestamp,
    }

    // 序列化并存储调用记录
    invocationAsBytes, err := json.Marshal(invocation)
    if err != nil {
        return err
    }

    txID := ctx.GetStub().GetTxID()
    return ctx.GetStub().PutState(txID, invocationAsBytes)
}


// 获取用户信息
func (s *SmartContract) GetUser(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 获取模型信息
func (s *SmartContract) GetModel(ctx contractapi.TransactionContextInterface, modelID string) (*Model, error) {
	modelBytes, err := ctx.GetStub().GetState(modelID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if modelBytes == nil {
		return nil, fmt.Errorf("the model %s does not exist", modelID)
	}
	var model Model
	err = json.Unmarshal(modelBytes, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// 获取模型历史记录
func (s *SmartContract) GetModelHistory(ctx contractapi.TransactionContextInterface, modelID string) ([]HistoryQueryResult, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(modelID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var model Model
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &model)
			if err != nil {
				return nil, err
			}
		} else {
			model = Model{ModelID: modelID}
		}

		timestamp := time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos))

		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}
		
		timestamp = timestamp.In(targetLocation)
		formattedTime := timestamp.Format("2006-01-02 15:04:05")

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &model,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}
	return records, nil
}

