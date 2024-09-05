package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		ModelList:    []*Model{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 大模型上链，传入用户ID、大模型上链信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过溯源码获取大模型的上链信息
	ModelAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将大模型的信息转换为结构体
	var model Model
	if ModelAsBytes != nil {
		err = json.Unmarshal(ModelAsBytes, &model)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal model: %v", err)
		}
	}

	//获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	time := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 给大模型信息中加上溯源码
	model.Traceability_code = traceability_code
	// 不同用户类型的上链的参数不一致
	switch userType {
	// 研发者
	case "研发者":
		// 将传入的大模型上链信息转换为Developer_input结构体
		model.Developer_input.De_TraceCode = arg1
		model.Developer_input.De_AIModelName = arg2
		model.Developer_input.De_DevelopmentBatch = arg3
		model.Developer_input.De_TrainingParams = arg4
		model.Developer_input.De_ResearchOrg = arg5
		model.Developer_input.De_Txid = txid
		model.Developer_input.De_Timestamp = time
	// 发布者
	case "发布者":
		// 将传入的大模型上链信息转换为Publisher_input结构体
		model.Publisher_input.Pu_OrganizationName = arg1
		model.Publisher_input.Pu_Platform = arg2
		model.Publisher_input.Pu_PublishInfo = arg3
		model.Publisher_input.Pu_OtherModelsHistory = arg4
		model.Publisher_input.Pu_ContactInfo = arg5
		model.Publisher_input.Pu_Txid = txid
		model.Publisher_input.Pu_Timestamp = time
	// 共享者
	case "共享者":
		// 将传入的大模型上链信息转换为Sharer_input结构体
		model.Sharer_input.Sh_SharingTime = arg1
		model.Sharer_input.Sh_UsageTime = arg2
		model.Sharer_input.Sh_PhoneNumber = arg3
		model.Sharer_input.Sh_UsingOrganization = arg4
		model.Sharer_input.Sh_ContactInfo = arg5
		model.Sharer_input.Sh_Txid = txid
		model.Sharer_input.Sh_Timestamp = time
	// 使用者
	case "使用者":
		// 将传入的大模型上链信息转换为Shop_input结构体
		model.User_input.U_FeedbackTime = arg1
		model.User_input.U_SalesTime = arg2
		model.User_input.U_UserName = arg3
		model.User_input.U_UserLocation = arg4
		model.User_input.U_UserContactInfo = arg5
		model.User_input.U_Txid = txid
		model.User_input.U_Timestamp = time

	}

	//将大模型的信息转换为json格式
	modelAsBytes, err := json.Marshal(model)
	if err != nil {
		return "", fmt.Errorf("failed to marshal model: %v", err)
	}
	//将大模型的信息存入区块链
	err = ctx.GetStub().PutState(traceability_code, modelAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put model: %v", err)
	}
	//将大模型存入用户的大模型列表
	err = s.AddModel(ctx, userID, &model)
	if err != nil {
		return "", fmt.Errorf("failed to add model to user: %v", err)

	}

	return txid, nil
}

// 添加大模型到用户的大模型列表
func (s *SmartContract) AddModel(ctx contractapi.TransactionContextInterface, userID string, model *Model) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.ModelList = append(user.ModelList, model)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取大模型的上链信息
func (s *SmartContract) GetModelInfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*Model, error) {
	ModelAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return &Model{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	// 将返回的结果转换为Model结构体
	var model Model
	if ModelAsBytes != nil {
		err = json.Unmarshal(ModelAsBytes, &model)
		if err != nil {
			return &Model{}, fmt.Errorf("failed to unmarshal model: %v", err)
		}
		if model.Traceability_code != "" {
			return &model, nil
		}
	}
	return &Model{}, fmt.Errorf("the model %s does not exist", traceability_code)
}

// 获取用户的大模型ID列表
func (s *SmartContract) GetModelList(ctx contractapi.TransactionContextInterface, userID string) ([]*Model, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.ModelList, nil
}

// 获取所有的大模型信息
func (s *SmartContract) GetAllModelInfo(ctx contractapi.TransactionContextInterface) ([]Model, error) {
	modelListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer modelListAsBytes.Close()
	var models []Model
	for modelListAsBytes.HasNext() {
		queryResponse, err := modelListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var model Model
		err = json.Unmarshal(queryResponse.Value, &model)
		if err != nil {
			return nil, err
		}
		// 过滤非大模型的信息
		if model.Traceability_code != "" {
			models = append(models, model)
		}
	}
	return models, nil
}

// 获取大模型上链历史
func (s *SmartContract) GetModelHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", traceability_code)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
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
			model = Model{
				Traceability_code: traceability_code,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
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
